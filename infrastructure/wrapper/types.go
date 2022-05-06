package wrapper

import (
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
	"github.com/sony/gobreaker"
)

type Wrapper struct {
	Cfg     *config.ServiceConfig
	Wrapper *gobreaker.CircuitBreaker
	Err     error
	CW      map[string]*gobreaker.CircuitBreaker
}
