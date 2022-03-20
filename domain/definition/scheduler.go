package definition

type SchedulerUsecase interface {
	DoCreateDummyData() error
	DoChunkCountingData() error
}
