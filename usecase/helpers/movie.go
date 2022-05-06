package helpersUC

import (
	"context"
	"fmt"

	"github.com/alhamsya/boilerplate-go/domain/constants"
)

func (u *UCInteractor) Tes(ctx context.Context) bool {
	u.CacheRepo.GetDetailMovie(ctx, "")
	tes1, _ := u.UtilsRepo.CurrentTimeF(constCommon.DateTime)
	fmt.Println(tes1)
	return false
}
