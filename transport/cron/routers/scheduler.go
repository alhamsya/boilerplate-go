package cronRouters

import (
	"context"
	"time"
)

func (cron *CronServer) ChunkCountingData(ctx context.Context) (interface{}, error) {
	time.Sleep(5 * time.Minute)
	return nil, nil
}

func (cron *CronServer) CreateDummyData(ctx context.Context) (interface{}, error) {
	time.Sleep(5 * time.Second)
	return nil, nil
}
