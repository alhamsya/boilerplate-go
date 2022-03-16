package schedulerUC

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/domain/constants/opening_account"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_error"
	"github.com/alhamsya/boilerplate-go/lib/helpers/custom_log"
	"golang.org/x/sync/errgroup"
)

func (uc *UCInteractor) DoCreateDummyData() error {

	customLog.InfoF("Start Test")
	time.Sleep(1 * time.Hour)

	const SecuritiesID = "bc1de9e521e61a7b227c9b3ff7bbec3d80edaef33c85c11e22e3d90f11a9fc3eo5YiLdi9sV"

	ctx := context.Background()

	oa, err := uc.Firestore.GetOpeningAccountBySecuritiesID(ctx, SecuritiesID)
	if err != nil {
		return customError.WrapFlag(err, "firestore", "GetOpeningAccountBySecuritiesID")
	}

	loc, _ := time.LoadLocation(constCommon.TimeLocalJakarta)
	dummyCreatedAt := time.Date(2022, 03, 1, 01, 01, 01, 0, loc)

	var (
		eg errgroup.Group
		m  sync.Mutex
	)

	ctx, cancel := context.WithTimeout(ctx, 30*time.Hour)
	defer cancel()

	for i := 1; i <= 2; i++ {
		m.Lock()
		oa.Partner.Data.Name = fmt.Sprintf("alhamsya%d", i)
		oa.StatusKYC = -1
		oa.CreatedAt = &dummyCreatedAt
		m.Unlock()

		eg.Go(func() error {
			err = uc.Firestore.CreateOpeningAccount(ctx, oa)
			if err != nil {
				return customError.WrapFlag(err, "firestore", "CreateOpeningAccount")
			}
			customLog.WarnF("data-%d", i)

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		customLog.ErrorLn(err)
	}

	customLog.InfoF("End Test")

	return nil
}

func (uc *UCInteractor) DoChunkCountingData() error {
	result, err := uc.UtilsRepo.GenerateRangeDate(2018, 1)
	if err != nil {
		return customError.WrapFlag(err, "utils", "GenerateRangeDate")
	}

	var (
		wg    sync.WaitGroup
		m     sync.Mutex
		total int
	)

	wg.Add(len(result))

	now := time.Now()
	num := 1
	for _, dataRange := range result {
		go func(dataRange map[string]time.Time) {
			defer wg.Done()
			startDate := dataRange["start_date"]
			endDate := dataRange["end_date"]
			data, err := uc.Firestore.GetOpeningAccountByStatusKYCAndDateRange(context.Background(), constOpeningAccount.IncompleteStatusKYC, &startDate, &endDate)
			if err != nil {
				customLog.ErrorF(err.Error())
				return
			}

			if len(data) == 0 {
				return
			}

			m.Lock()
			total = total + len(data)
			fmt.Printf("%d - %v: %d \n", num, startDate, total)
			num++
			m.Unlock()
		}(dataRange)
	}

	wg.Wait()

	fmt.Println("data:", total, "| elapsed time: ", time.Since(now).Seconds(), " secs")

	if err := uc.CacheRepo.SetKYCByStatus(context.Background(), constOpeningAccount.IncompleteStatusKYC, total); err != nil {
		return customError.WrapFlag(err, "redis", "SetKYCByStatus")
	}

	return nil
}
