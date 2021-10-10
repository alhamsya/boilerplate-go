package repository

import "github.com/alhamsya/boilerplate-go/infrastructure/services/exter/omdb"

type OMDBRepo interface {
	GetListMovie(page int, search string) (*omdb.OMDBList, error)
}
