package definition

import "context"

type ConsumerUsecase interface {
	DoPayment(ctx context.Context) error
}
