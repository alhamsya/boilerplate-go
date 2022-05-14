package wrappers

import (
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
	"github.com/sony/gobreaker"
)

type Wrapper struct {
	Cfg     *config.ServiceConfig
	Wrapper *gobreaker.CircuitBreaker
	Err     error
	CW      map[string]*gobreaker.CircuitBreaker
}
