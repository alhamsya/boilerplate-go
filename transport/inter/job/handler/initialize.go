package jobHandler

import "github.com/alhamsya/boilerplate-go/transport/inter/job/handler/scheduler"

func initScheduler(c *schedulerHandler.Handler) error {
	return schedulerHandler.New(c).Register()
}
