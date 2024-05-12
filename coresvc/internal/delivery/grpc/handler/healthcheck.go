package handler

import (
	"context"
	pb "github.com/nofendian17/openota/coresvc/gen/go/proto/healthcheck"
	"github.com/nofendian17/openota/coresvc/internal/usecase/healthcheck"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HealthCheckServiceServer struct {
	healthCheck healthcheck.UseCase
	pb.UnimplementedHealthCheckServiceServer
}

func NewHealthCheckServiceServer(healthCheck healthcheck.UseCase) *HealthCheckServiceServer {
	return &HealthCheckServiceServer{
		healthCheck: healthCheck,
	}
}

func (h *HealthCheckServiceServer) HealthCheck(ctx context.Context, request *emptypb.Empty) (*pb.GetHealthCheckResponse, error) {
	health, err := h.healthCheck.Health(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetHealthCheckResponse{
		Health: &pb.Health{
			Version: health.Version,
			Uptime:  health.Uptime,
			Cpu:     health.CPU,
			Memory:  health.Memory,
		}}, nil
}

func (h *HealthCheckServiceServer) Readiness(ctx context.Context, request *emptypb.Empty) (*pb.GetReadinessResponse, error) {
	readiness, err := h.healthCheck.Readiness(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	checks := make([]*pb.Check, len(readiness.Checks))
	for i, check := range readiness.Checks {
		checks[i] = &pb.Check{
			Name:   check.Name,
			Status: check.Status,
			Error:  check.Error,
		}
	}

	return &pb.GetReadinessResponse{
		Readiness: &pb.Readiness{
			Status: readiness.Status,
			Checks: checks,
		},
	}, nil
}

func (h *HealthCheckServiceServer) Ping(context.Context, *emptypb.Empty) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Reply: "pong",
	}, nil
}
