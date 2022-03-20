package cronRouters

import "context"

func (cron *CronServer) ChunkCountingData(ctx context.Context) (interface{}, error) {
	return nil, cron.CronInteractor.CronInterface.DoChunkCountingData(ctx)
}
