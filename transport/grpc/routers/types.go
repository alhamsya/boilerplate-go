package grpcRouters

import (
	"github.com/alhamsya/boilerplate-go/domain/repository"
	"github.com/alhamsya/boilerplate-go/lib/helpers/config"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	Cfg            *config.ServiceConfig
	GrpcInteractor *GrpcInteractor
	App            *grpc.Server
}

type GrpcInteractor struct {
	GrpcInterface repository.GrpcUsecase
}
