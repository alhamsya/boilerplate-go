package jobRouters

import (
	"github.com/alhamsya/boilerplate-go/domain/definition"
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
)

type JobServer struct {
	Cfg           *config.ServiceConfig
	JobInteractor *JobInteractor
}

type JobInteractor struct {
	SchedulerInterface definition.SchedulerInterface
}
