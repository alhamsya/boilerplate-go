package definition

import "context"

type CronUsecase interface {
	DoCreateDummyData(ctx context.Context) error
	DoChunkCountingData(ctx context.Context) error
}
