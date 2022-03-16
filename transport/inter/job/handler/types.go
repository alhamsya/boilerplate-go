package jobHandler

import (
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/alhamsya/boilerplate-go/transport/inter/job/routers"
)

type Handler struct {
	Cfg        *config.ServiceConfig
	Interactor *jobRouters.JobInteractor
}
