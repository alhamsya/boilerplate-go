package consumerRouters

import (
	"context"
)

func New(this *ConsumerServer) *ConsumerServer {
	return &ConsumerServer{
		Cfg:                this.Cfg,
		ConsumerInteractor: this.ConsumerInteractor,
	}
}

func (consume *ConsumerServer) Register() map[string][]func(context.Context) (interface{}, error) {
	return map[string][]func(context.Context) (interface{}, error){
		"tes": {
			consume.Payment,
		},
	}
}
