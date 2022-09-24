package pinger

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/chenjiandongx/pinger"
)

type Core struct {
	DB

	FetchTimeout time.Duration

	stopPingRoutine chan struct{}
	pingRoutineDone sync.WaitGroup
}

type Host struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	URL   string `json:"url" db:"url"`
	Mean  int    `json:"mean" db:"mean"`
	Last  int    `json:"last" db:"last"`
	Best  int    `json:"best" db:"best"`
	Worst int    `json:"worst" db:"worst"`
}

func ValidateHost(host *Host) error {
	if host.Name == "" {
		return fmt.Errorf("missing host name")
	}
	if host.URL == "" {
		return fmt.Errorf("missing host url")
	}

	_, err := url.Parse(host.URL)
	if err != nil {
		return fmt.Errorf("invalid host url: %s", err)
	}

	return nil
}

type DB interface {
	UpsertHost(ctx context.Context, host *Host) error
	GetHost(ctx context.Context, hostID int) (*Host, error)
	GetHosts(ctx context.Context, limit, skip int) ([]*Host, error)
	RemoveHost(ctx context.Context, hostID int) error
}

func NewCore(db DB, timeout time.Duration) *Core {
	core := Core{
		DB:              db,
		FetchTimeout:    timeout,
		pingRoutineDone: sync.WaitGroup{},
		stopPingRoutine: make(chan struct{}),
	}

	// Defaults
	if timeout == 0 {
		core.FetchTimeout = 5 * time.Second
	}

	return &core
}

func (c *Core) Run() *Core {
	c.pingRoutineDone.Add(1)

	go func() {
		for {
			select {
			case <-c.stopPingRoutine:
				c.pingRoutineDone.Done()
				return
			default:
				if err := c.pingAll(); err != nil {
					log.Println(err)
					continue
				}
			}
		}
	}()

	return c
}

func (c *Core) Stop() {
	c.stopPingRoutine <- struct{}{}
	c.pingRoutineDone.Wait()
}

func (c *Core) pingAll() error {
	ctx, cancel := context.WithTimeout(context.Background(), c.FetchTimeout)
	defer cancel()
	hosts, err := c.GetHosts(ctx, 10, 0)
	if err != nil {
		return err
	}

	return c.Ping(hosts...)
}

func (c *Core) Ping(hosts ...*Host) error {
	urls := make([]string, len(hosts))
	for idx, host := range hosts {
		urls[idx] = host.URL
	}

	stats, err := pinger.HTTPPing(nil, urls...)
	if err != nil {
		return err
	}

	for _, stat := range stats {
		var currentHost *Host
		for _, host := range hosts {
			if host.URL == stat.Host {
				currentHost = host
			}
		}

		if currentHost == nil {
			continue
		}

		// Make sure the host still exists
		_, err := c.GetHost(context.Background(), currentHost.ID)
		if err != nil {
			log.Println(err)
			continue
		}

		// Update host stats
		currentHost.Mean = int(stat.Mean.Milliseconds())
		currentHost.Best = int(stat.Best.Milliseconds())
		currentHost.Last = int(stat.Last.Milliseconds())
		currentHost.Worst = int(stat.Worst.Milliseconds())

		err = c.UpsertHost(context.Background(), currentHost)
		if err != nil {
			return err
		}
	}

	return nil
}
