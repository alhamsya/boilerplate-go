package grpcHandler

import (
	"fmt"
	"github.com/alhamsya/boilerplate-go/infrastructure/middleware"
	"github.com/alhamsya/boilerplate-go/service/inter/grpc/routers"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/lib/helpers/grace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcOpenTracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

func New(this *Handler) *Handler {
	srv := grpc.NewServer(
		// MaxConnectionIdle is a duration for the amount of time after which an
		// idle connection would be closed by sending a GoAway. Idleness duration is
		// defined since the most recent time the number of outstanding RPCs became
		// zero or the connection establishment
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: constCommon.DefaultGRCPMaxConnectionIdle,
		}),
		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				grpcOpenTracing.UnaryServerInterceptor(),
				grpcValidator.UnaryServerInterceptor(),
				grpcPrometheus.UnaryServerInterceptor,
				middleware.GrpcLoggingInterceptor,
			),
		),
	)

	//enabling histogram to prometheus
	grpcPrometheus.EnableHandlingTimeHistogram()

	//healthy check gRPC grpc-health-probe
	grpc_health_v1.RegisterHealthServer(srv, NewHealthChecker())

	//gRPC Server Reflection provides information about publicly-accessible gRPC service on a server,
	//and assists clients at runtime to construct RPC requests and responses without precompiled service information.
	reflection.Register(srv)

	routeHandler := &grpcRouters.GrpcServer{
		Cfg:            this.Cfg,
		GrpcInteractor: this.Interactor,
		App:            srv,
	}

	grpcRouters.New(routeHandler).Register()

	return &Handler{
		Cfg:        this.Cfg,
		Interactor: this.Interactor,
		GrpcServer: srv,
	}
}

func (h *Handler) Run() error {
	return grace.ServeGRPC(h.GrpcServer,fmt.Sprintf(":%s", h.Cfg.GrpcServer.Port), constCommon.DefaultServerGRPCGraceTimeout)
}
