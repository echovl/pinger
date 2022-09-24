package mysql

import (
	"context"
	"time"

	"github.com/echovl/pinger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type mysqlDB struct {
	x *sqlx.DB
}

var _ pinger.DB = (*mysqlDB)(nil)

const (
	upsertHostQuery = `INSERT INTO hosts (
        id,
        name, 
        url,
        mean,
        last,
        best,
        worst
    ) VALUES (?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE
        name=VALUES(name),
        url=VALUES(url),
        mean=VALUES(mean),
        last=VALUES(last),
        best=VALUES(best),
        worst=VALUES(worst)
    `
	getHostQuery    = `SELECT * FROM hosts WHERE id = ?`
	getHostsQuery   = `SELECT * FROM hosts LIMIT ? OFFSET ?`
	removeHostQuery = `DELETE FROM hosts WHERE id = ?`
)

func NewDB(dsn string) (pinger.DB, error) {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Defaults
	db.SetConnMaxIdleTime(15 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &mysqlDB{db}, nil
}

func (db *mysqlDB) UpsertHost(ctx context.Context, host *pinger.Host) error {
	res, err := db.x.ExecContext(ctx, upsertHostQuery,
		host.ID,
		host.Name,
		host.URL,
		host.Mean,
		host.Last,
		host.Best,
		host.Worst)
	if err != nil {
		return err
	}

	if host.ID == 0 {
		lastInsertID, err := res.LastInsertId()
		if err != nil {
			return err
		}
		host.ID = int(lastInsertID)
	}

	return nil
}

func (db *mysqlDB) GetHost(ctx context.Context, hostID int) (*pinger.Host, error) {
	var host pinger.Host
	err := db.x.GetContext(ctx, &host, getHostQuery, hostID)
	if err != nil {
		return nil, err
	}
	return &host, nil
}

func (db *mysqlDB) GetHosts(ctx context.Context, limit, skip int) ([]*pinger.Host, error) {
	var hosts []*pinger.Host
	err := db.x.SelectContext(ctx, &hosts, getHostsQuery, limit, skip)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func (db *mysqlDB) RemoveHost(ctx context.Context, hostID int) error {
	_, err := db.x.ExecContext(ctx, removeHostQuery, hostID)
	return err
}
