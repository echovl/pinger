package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/echovl/pinger"
	"github.com/echovl/pinger/db/mysql"
	pingerhttp "github.com/echovl/pinger/http"
)

type Config struct {
	// Server port
	Port int `env:"PORT" envDefault:"8080"`

	// Server timeouts in seconds
	ServerReadTimeout       int64 `env:"SERVER_READ_TIMEOUT"`
	ServerIdleTimeout       int64 `env:"SERVER_IDLE_TIMEOUT"`
	ServerReadHeaderTimeout int64 `env:"SERVER_READ_HEADER_TIMEOUT"`

	// Databases
	MySQLDSN string `env:"MYSQL_DSN"`
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("config: %+v", cfg)

	db, err := mysql.NewDB(cfg.MySQLDSN)
	if err != nil {
		log.Fatal(err)
	}

	core := pinger.NewCore(db, 10*time.Second)
	defer core.Run().Stop()

	// Server defaults
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           pingerhttp.Handler(core),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		IdleTimeout:       5 * time.Minute,
	}

	if cfg.ServerReadHeaderTimeout > 0 {
		server.ReadHeaderTimeout = time.Duration(cfg.ServerReadHeaderTimeout) * time.Second
	}
	if cfg.ServerIdleTimeout > 0 {
		server.IdleTimeout = time.Duration(cfg.ServerIdleTimeout) * time.Second
	}
	if cfg.ServerReadTimeout > 0 {
		server.ReadTimeout = time.Duration(cfg.ServerReadTimeout) * time.Second
	}

	log.Fatal(server.ListenAndServe())
}

func loadConfig() (cfg Config, err error) {
	err = env.Parse(&cfg)
	return
}
