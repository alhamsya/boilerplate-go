package database

import (
	"fmt"
	"time"

	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

//monitoringDB monitoring database with pooling mechanism
func (d *DB) monitoringDB(driver string) (*sqlx.DB, error) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	if d.TimeoutExceededInMinutes < 1 {
		d.TimeoutExceededInMinutes = 5
	}

	timeoutExceeded := time.After(time.Duration(d.TimeoutExceededInMinutes) * time.Minute)
	for {
		select {
		case <-timeoutExceeded:
			return nil, fmt.Errorf("db connection failed after %s timeout", d.ConnectionMaxLifetime)
		case <-ticker.C:
			db, err := d.connectDB(driver)
			if err == nil {
				return db, nil
			}
			log.Println("[Error]: DB reconnect error", err.Error())
		}
	}
}

//connectDB connect database with configuration
func (d *DB) connectDB(driver string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, d.DSN)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	if d.MaxIdleConn > 0 {
		db.SetMaxIdleConns(d.MaxIdleConn)
	}else {
		db.SetMaxIdleConns(0)
	}

	if d.MaxOpenConn > 0 {
		db.SetMaxOpenConns(d.MaxOpenConn)
	}

	if d.ConnectionMaxLifetime > 0 {
		db.SetConnMaxLifetime(d.ConnectionMaxLifetime)
	}

	return db, nil
}
