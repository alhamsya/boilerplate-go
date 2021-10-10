package database

import (
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	DriverMySQL    = "mysql"
	DriverPostgres = "postgres"
)

type Store struct {
	Master *sqlx.DB
	Slave  *sqlx.DB
}

type DB struct {
	DSN                   string
	MaxOpenConn           int
	MaxIdleConn           int
	ConnectionMaxLifetime time.Duration

	TimeoutExceededInMinutes int
}
