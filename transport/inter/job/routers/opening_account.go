package jobRouters

func (job *JobServer) ChunkCountingData() error {
	return job.JobInteractor.SchedulerInterface.DoChunkCountingData()
}
