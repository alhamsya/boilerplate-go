package jobRouters

import "context"

func New(this *JobServer) *JobServer {
	return &JobServer{
		Cfg:           this.Cfg,
		JobInteractor: this.JobInteractor,
	}
}

func (job *JobServer) Register() map[string][]func(context.Context) (interface{}, error) {
	return map[string][]func(context.Context) (interface{}, error){
		"tes": {
			job.ChunkCountingData,
		},
		//"tes1": {
		//	h.Interactor.SchedulerInterface.DoCreateDummyData,
		//},
	}
}
