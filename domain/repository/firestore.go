package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/alhamsya/boilerplate-go/domain/models/opening_account"
)

type FirestoreRepo interface {
	GetOpeningAccountBySecuritiesID(ctx context.Context, securitiesID string) (oa *modelOpeningAccount.OpeningAccount, err error)
	GetOpeningAccountByStatusKYCAndDateRange(ctx context.Context, statusKYC int64, startDate, endDate *time.Time) ([]*firestore.DocumentSnapshot, error)
	GetSnapOpeningAccountByStatusKYC(ctx context.Context, statusKYC int64) (*firestore.QuerySnapshot, error)
	CreateOpeningAccount(ctx context.Context, oa *modelOpeningAccount.OpeningAccount) (err error)
}
