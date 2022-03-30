package firestore

import (
	"context"
	"errors"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/alhamsya/boilerplate-go/domain/models/opening_account"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_error"
	"google.golang.org/api/iterator"
)

func (fs *ServiceFirestore) GetOpeningAccountBySecuritiesID(ctx context.Context, securitiesID string) (oa *modelOpeningAccount.OpeningAccount, err error) {
	coll := fs.clients["sekuritas_opening_account"].Collection("sekuritas-opening-account")
	//defer fs.Clients["sekuritas_opening_account"].Close()

	snap, err := coll.Where("securities_id", "==", securitiesID).Documents(ctx).Next()
	if errors.Is(err, iterator.Done) {
		return nil, err
	}

	if err != nil {
		return nil, customError.WrapFlag(err, "firestore", "Collection")
	}

	err = fs.UtilsRepo.ToStruct(snap.Data(), &oa)
	if err != nil {
		return nil, customError.WrapFlag(err, "utils", "ToStruct")
	}

	oa.ID = snap.Ref.ID

	return oa, nil
}

func (fs *ServiceFirestore) GetOpeningAccountByStatusKYCAndDateRange(ctx context.Context, statusKYC int64, startDate, endDate *time.Time) ([]*firestore.DocumentSnapshot, error) {
	coll := fs.clients["sekuritas_opening_account"].Collection("test")

	snap, err := coll.Where("status_kyc", "==", statusKYC).
		Where("created_at", ">=", startDate).
		Where("created_at", "<", endDate).
		Select("id").
		Documents(ctx).GetAll()

	if err != nil {
		return nil, customError.WrapFlag(err, "firestore", "Documents.GetAll")
	}

	return snap, nil
}

func (fs *ServiceFirestore) GetSnapOpeningAccountByStatusKYC(ctx context.Context, statusKYC int64) (*firestore.QuerySnapshot, error) {
	coll := fs.clients["sekuritas_opening_account"].Collection("sekuritas-opening-account")

	iter := coll.Where("status_kyc", "==", statusKYC).Snapshots(ctx)
	snap, err := iter.Next()
	if err != nil {
		return nil, err
	}

	return snap, nil
}

func (fs *ServiceFirestore) CreateOpeningAccount(ctx context.Context, oa *modelOpeningAccount.OpeningAccount) (err error) {
	coll := fs.clients["sekuritas_opening_account"].Collection("test")
	//defer fs.Clients["sekuritas_opening_account"].Close()

	_, err = coll.Doc(fs.UtilsRepo.RandomString(10)).Set(ctx, fs.UtilsRepo.ToMap(&oa))
	if err != nil {
		return customError.WrapFlag(err, "firestore", "Document")
	}

	return nil
}
