package consumerRouters

import "context"

func (consume *ConsumerServer) Payment(ctx context.Context) (interface{}, error) {
	err := consume.ConsumerInteractor.ConsumerInterface.DoPayment(ctx)
	return nil, err
}
