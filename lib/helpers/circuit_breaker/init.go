package circuitBreaker

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/sony/gobreaker"
)

// shouldBeSwitchedToOpen checks if the circuit breaker should
// switch to the Open state
func shouldBeSwitchedToOpen(counts gobreaker.Counts) bool {
	failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
	return counts.Requests >= 3 && failureRatio >= 0.6
}

func NewCircuitBreaker(cfg *config.CallWrapper, usecase string) (*gobreaker.CircuitBreaker, error) {
	var reqTimeout time.Duration

	//validate usecase
	if strings.TrimSpace(usecase) == "" {
		return nil, fmt.Errorf("the usecase shoul be setup '%s'", usecase)
	}

	///validate handle type in the configuration
	if ok := EnumAPIType[cfg.APIType]; !ok {
		return nil, fmt.Errorf("handler type '%s' must be setup in the %s", cfg.APIType, usecase)
	}

	//validate call type for config
	if ok := EnumCallType[cfg.CallType]; !ok {
		return nil, fmt.Errorf("unknown call type '%s' must be setup in the usecase '%s'", cfg.CallType, usecase)
	}

	//set timeout from config
	if cfg.TimeoutInMS != 0 {
		reqTimeout = time.Duration(cfg.TimeoutInMS) * time.Millisecond
	}

	return gobreaker.NewCircuitBreaker(gobreaker.Settings{
		//is the name of the CircuitBreaker
		Name: fmt.Sprintf("[%s]: [%s]: %s", cfg.APIType, cfg.CallType, usecase),

		//is the maximum number of requests allowed to pass through when the CircuitBreaker is half-open.
		//If MaxRequests is 0, CircuitBreaker allows only 1 request.
		MaxRequests: 0,

		/*When to flush counters int the Closed state*/
		//is the cyclic period of the closed state for CircuitBreaker to clear
		//the internal Counts, described later in this section.
		//If Interval is 0, CircuitBreaker doesn't clear the internal Counts during the closed state.
		Interval: 5 * time.Second,

		/*Time to switch from Open to Half-open*/
		//is the period of the open state, after which the state of CircuitBreaker becomes half-open.
		//If Timeout is 0, the timeout value of CircuitBreaker is set to 60 seconds.
		Timeout: reqTimeout,

		/*Function with check when to switch from Closed to Open*/
		//is called with a copy of Counts whenever a request fails in the closed state.
		//If ReadyToTrip returns true, CircuitBreaker will be placed into the open state.
		//If ReadyToTrip is nil, default ReadyToTrip is used. Default ReadyToTrip returns true
		//when the number of consecutive failures is more than 5.
		ReadyToTrip: shouldBeSwitchedToOpen,

		//is called whenever the state of CircuitBreaker changes
		OnStateChange: func(_ string, from gobreaker.State, to gobreaker.State) {
			// Handler for every state change. We'll use for debugging purpose
			log.Println("state changed from", from.String(), "to", to.String())
		},
	}), nil
}
