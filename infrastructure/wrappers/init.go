package wrappers

import (
	"log"

	"github.com/alhamsya/boilerplate-go/lib/helpers/circuit_breaker"
	"github.com/sony/gobreaker"
)

func New(this *Wrapper) *Wrapper {
	var err error
	this.CW = make(map[string]*gobreaker.CircuitBreaker)
	for usecase, cwCfg := range this.Cfg.CallWrapper {
		this.CW[usecase], err = circuitBreaker.NewCircuitBreaker(cwCfg, usecase)
		if err != nil {
			log.Fatalf("[INIT] [CALL WRAPPER] fail NewCircuitBreaker: %v", err)
		}
	}

	log.Printf("[IGNORE] [CALL WRAPPER] initialize")
	return this
}
