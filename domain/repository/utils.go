package repository

type UtilsRepo interface {
	CurrentTimeF(format string) (string, error)
}
