package grpcHandler

import (
	"github.com/alhamsya/boilerplate-go/infrastructure/config"
	"github.com/alhamsya/boilerplate-go/service/inter/grpc/routers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthChecker struct {
	grpc_health_v1.HealthServer
}

type Handler struct {
	Cfg        *config.ServiceConfig
	GrpcServer *grpc.Server
	Interactor *grpcRouters.GrpcInteractor
}
