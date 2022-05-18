package cronUC

import (
	"context"
)

func (uc *UCInteractor) DoCreateDummyData(ctx context.Context) error {
	uc.HelpersRepo.Tes(ctx)
	return nil
}
