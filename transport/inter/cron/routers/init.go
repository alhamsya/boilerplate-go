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
		"tes": {
			cron.ChunkCountingData,
		},
		//"tes1": {
		//	h.Interactor.SchedulerInterface.DoCreateDummyData,
		//},
	}
}
