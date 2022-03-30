package cronRouters

import "context"

func New(this *CronServer) *CronServer {
	return &CronServer{
		Cfg:            this.Cfg,
		CronInteractor: this.CronInteractor,
	}
}

func (cron *CronServer) Register() map[string][]func(context.Context) (interface{}, error) {
	return map[string][]func(context.Context) (interface{}, error){
		"chunk_counting": {
			cron.ChunkCountingData,
		},
		"create_dummy": {
			cron.CreateDummyData,
		},
	}
}
