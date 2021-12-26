package wrapper

import (
	"log"

	"github.com/alhamsya/boilerplate-go/lib/helpers/circuit_breaker"
	"github.com/sony/gobreaker"
)

func New(this *Wrapper) *Wrapper {
	var err error
	this.CW = make(map[string]*gobreaker.CircuitBreaker)
	for usecase, cwCfg := range this.Cfg.CallWrapperConfig {
		this.CW[usecase], err = circuitBreaker.NewCircuitBreaker(cwCfg, usecase)
		if err != nil {
			log.Fatalf("fail initialize call wrapper")
		}
	}
	return this
}
