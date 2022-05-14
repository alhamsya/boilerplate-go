package repositorys

import (
	"context"
)

type FirestoreRepo interface {
	GetAccount(ctx context.Context) (oa interface{}, err error)
}
