package database

import (
	"fmt"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
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
	connMaster, err := dbMaster.monitoringDB(dbDriver)
	if err != nil {
		return nil, err
	}

	dbSlave := &DB{
		DSN:                   cfg.SlaveDSN,
		MaxOpenConn:           cfg.MaxOpenConn,
		MaxIdleConn:           cfg.MaxIdleConn,
		ConnectionMaxLifetime: cfg.ConnectionMaxLifetime,
	}
	connSlave, err := dbSlave.monitoringDB(dbDriver)
	if err != nil {
		return nil, err
	}

	return &Store{
		Master: connMaster,
		Slave:  connSlave,
	}, nil
}
