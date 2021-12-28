package config

import "time"

type ServiceConfig struct {
	RestServer        RestServer
	GrpcServer        GrpcServer
	Databases         map[string]*DBConfig
	Toggle            map[string]*Toggle
	External          map[string]*External
	CallWrapper map[string]*CallWrapper
	CORS              CORS
	Redis             Redis
}

type DBConfig struct {
	SlaveDSN              string        `json:"slave_dsn"`
	MasterDSN             string        `json:"master_dsn"`
	MaxOpenConn           int           `json:"max_open_conn"`
	MaxIdleConn           int           `json:"max_idle_conn"`
	ConnectionMaxLifetime time.Duration `json:"connection_max_lifetime"`

	TimeoutExceededInMinutes int  `json:"timeout_exceeded_in_minutes"`
	NoPingOnOpen             bool `json:"no_ping_on_open"`
}

type RestServer struct {
	Port          string
	Timeout       time.Duration
	AccessLogFile string
	ErrorLogFile  string
}

type CORS struct {
	AllowOrigins []string
}

type GrpcServer struct {
	Port                string
	AuthTimeoutInHour   int
	ContextTimeoutInSec int
	AccessLogFile       string
	ErrorLogFile        string
}

type Toggle struct {
	IsActive bool
}

type External struct {
	Endpoint string
	Key      string
}

type CallWrapper struct {
	APIType     string
	CallType    string
	TimeoutInMS int64
}

type Redis struct {
	Address          string
	Username         string
	Password         string
	DB               int
	MinIdleConn      int
	MaxConnAgeInSec  int
	IdleTimeoutInSec int
}
