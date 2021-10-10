package grpcHandler

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"
)

//Check check grpc
func (s *HealthChecker) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

//NewHealthChecker new check grpc health
func NewHealthChecker() *HealthChecker {
	return &HealthChecker{}
}
