package routers

import (
	"github.com/alhamsya/boilerplate-go/domain/definition"
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	Cfg            *config.MainConfig
	GrpcInteractor *GrpcInteractor
	App            *grpc.Server
}

type GrpcInteractor struct {
	GrpcInterface definition.GrpcInterface
}
