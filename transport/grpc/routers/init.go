package grpcRouters

import (
	pb "github.com/alhamsya/boilerplate-go/protos"
)

func New(this *GrpcServer) *GrpcServer {
	return &GrpcServer{
		Cfg:            this.Cfg,
		App:            this.App,
		GrpcInteractor: this.GrpcInteractor,
	}
}

func (grpc *GrpcServer) Register() {
	pb.RegisterServicesServer(grpc.App, grpc)
}
