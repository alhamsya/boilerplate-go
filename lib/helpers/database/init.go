package database

import (
	"fmt"

	"github.com/alhamsya/boilerplate-go/infrastructure/config"
)

func New(cfg *config.DBConfig, dbDriver string) (*Store, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is nil")
	}

	dbMaster := &DB{
		DSN:                   cfg.MasterDSN,
		MaxOpenConn:           cfg.MaxOpenConn,
		MaxIdleConn:           cfg.MaxIdleConn,
		ConnectionMaxLifetime: cfg.ConnectionMaxLifetime,
	}
	connMaster, err := dbMaster.connectAndMonitoring(dbDriver, cfg.NoPingOnOpen)
	if err != nil {
		return nil, err
	}

	dbSlave := &DB{
		DSN:                   cfg.SlaveDSN,
		MaxOpenConn:           cfg.MaxOpenConn,
		MaxIdleConn:           cfg.MaxIdleConn,
		ConnectionMaxLifetime: cfg.ConnectionMaxLifetime,
	}
	connSlave, err := dbSlave.connectAndMonitoring(dbDriver, cfg.NoPingOnOpen)
	if err != nil {
		return nil, err
	}

	return &Store{
		Master: connMaster,
		Slave:  connSlave,
	}, nil
}