package cronRouters

import (
	"github.com/alhamsya/boilerplate-go/infrastructure/middlewares/cron"
)

func New(this *CronServer) *CronServer {
	return &CronServer{
		Cfg:            this.Cfg,
		CronInteractor: this.CronInteractor,
	}
}

func (cron *CronServer) Register() map[string][]cronMiddleware.FuncOrigin {
	return map[string][]cronMiddleware.FuncOrigin{
		"chunk_counting": {
			cron.ChunkCountingData,
		},
		"create_dummy": {
			cron.CreateDummyData,
		},
	}
}
