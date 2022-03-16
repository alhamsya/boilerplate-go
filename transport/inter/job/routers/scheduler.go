package jobRouters

import "context"

func (job *JobServer) ChunkCountingData(ctx context.Context) (interface{}, error) {
	return nil, job.JobInteractor.SchedulerInterface.DoChunkCountingData()
}
