package handler

import (
	"context"
	"github.com/nofendian17/openota/authsvc/internal/usecase/healthcheck"
	"github.com/nofendian17/openota/authsvc/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HealthCheckServiceServer struct {
	healthCheck healthcheck.UseCase
	proto.UnimplementedHealthCheckServiceServer
}

func NewHealthCheckServiceServer(healthCheck healthcheck.UseCase) *HealthCheckServiceServer {
	return &HealthCheckServiceServer{
		healthCheck: healthCheck,
	}
}

func (h *HealthCheckServiceServer) GetHealthCheck(ctx context.Context, request *emptypb.Empty) (*proto.GetHealthCheckResponse, error) {
	health, err := h.healthCheck.Health(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.GetHealthCheckResponse{
		Health: &proto.Health{
			Version: health.Version,
			Uptime:  health.Uptime,
			Cpu:     health.CPU,
			Memory:  health.Memory,
		}}, nil
}

func (h *HealthCheckServiceServer) GetReadiness(ctx context.Context, request *emptypb.Empty) (*proto.GetReadinessResponse, error) {
	readiness, err := h.healthCheck.Readiness(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	checks := make([]*proto.Check, len(readiness.Checks))
	for i, check := range readiness.Checks {
		checks[i] = &proto.Check{
			Name:   check.Name,
			Status: check.Status,
			Error:  check.Error,
		}
	}

	return &proto.GetReadinessResponse{
		Readiness: &proto.Readiness{
			Status: readiness.Status,
			Checks: checks,
		},
	}, nil
}

func (h *HealthCheckServiceServer) Ping(context.Context, *emptypb.Empty) (*proto.PingResponse, error) {
	return &proto.PingResponse{
		Reply: "pong",
	}, nil
}
