package grpcRouters

import (
	"github.com/alhamsya/boilerplate-go/domain/repositorys"
	"github.com/alhamsya/boilerplate-go/lib/managers/config"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	Cfg            *config.ServiceConfig
	GrpcInteractor *GrpcInteractor
	App            *grpc.Server
}

type GrpcInteractor struct {
	GrpcInterface repositorys.GrpcUsecase
}
