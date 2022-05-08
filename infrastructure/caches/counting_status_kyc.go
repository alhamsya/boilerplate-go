package caches

import (
	"context"
	"fmt"

	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_error"
)

const (
	KeyCountingStatusKYC = "counting_status_kyc:%d"
)

func (cache *ServiceCache) IncrKYCByStatus(ctx context.Context, statusKYC int64) (err error) {
	err = cache.Redis.Incr(ctx, fmt.Sprintf(KeyCountingStatusKYC, statusKYC)).Err()
	if err != nil {
		return customError.WrapFlag(err, "redis", "Incr")
	}

	return nil
}

func (cache *ServiceCache) DecrKYCByStatus(ctx context.Context, statusKYC int64) (err error) {
	err = cache.Redis.Decr(ctx, fmt.Sprintf(KeyCountingStatusKYC, statusKYC)).Err()
	if err != nil {
		return customError.WrapFlag(err, "redis", "Incr")
	}

	return nil
}

func (cache *ServiceCache) SetKYCByStatus(ctx context.Context, statusKYC int64, value int) (err error) {
	err = cache.Redis.Set(ctx, fmt.Sprintf(KeyCountingStatusKYC, statusKYC), value, 0).Err()
	if err != nil {
		return customError.WrapFlag(err, "redis", "Incr")
	}

	return nil
}

func (cache *ServiceCache) GetKYCByStatus(ctx context.Context, statusKYC int64) (total int, err error) {
	total, err = cache.Redis.Get(ctx, fmt.Sprintf(KeyCountingStatusKYC, statusKYC)).Int()
	if err != nil {
		return -1, customError.WrapFlag(err, "redis", "Incr")
	}

	return total, nil
}
