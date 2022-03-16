package jobRouters

func New(this *JobServer) *JobServer {
	return &JobServer{
		Cfg:           this.Cfg,
		JobInteractor: this.JobInteractor,
	}
}

func (job *JobServer) Register() map[string][]func() error {
	return map[string][]func() error{
		"tes": {
			job.ChunkCountingData,
		},
		//"tes1": {
		//	h.Interactor.SchedulerInterface.DoCreateDummyData,
		//},
	}
}
