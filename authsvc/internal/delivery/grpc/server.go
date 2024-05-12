package grpc

import (
	"context"
	"fmt"
	"github.com/nofendian17/gopkg/logger"
	"github.com/nofendian17/openota/authsvc/internal/container"
	"github.com/nofendian17/openota/authsvc/internal/delivery/grpc/handler"
	"github.com/nofendian17/openota/authsvc/internal/delivery/grpc/middleware"
	"github.com/nofendian17/openota/authsvc/internal/usecase"
	"github.com/nofendian17/openota/authsvc/proto"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

// Server defines the methods for a gRPC server.
type Server interface {
	Start(port int) error
	Stop(ctx context.Context) error
}

// server represents the gRPC server.
type server struct {
	grpcServer *grpc.Server
	logger     logger.Logger
	useCase    *usecase.UseCase
}

// Start starts the gRPC server.
func (s *server) Start(port int) error {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer(
		// grpc.ChainUnaryInterceptor(middleware.LoggingInterceptor, middleware.AuthorizationInterceptor),
		grpc.ChainUnaryInterceptor(middleware.LoggingInterceptor),
	)

	// Register your gRPC service implementation with the gRPC server
	// e.g., pb.RegisterYourServiceServer(s.grpcServer, s.service)
	healthCheckServiceServer := handler.NewHealthCheckServiceServer(s.useCase.Health)
	proto.RegisterHealthCheckServiceServer(s.grpcServer, healthCheckServiceServer)

	s.logger.Info(context.Background(), fmt.Sprintf("gRPC server started on port %d", port), nil)

	if err = s.grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

// Stop gracefully shuts down the gRPC server.
func (s *server) Stop(ctx context.Context) error {
	if s.grpcServer == nil {
		return nil
	}

	s.logger.Info(ctx, "gRPC server stopped", nil)
	// Gracefully stop the gRPC server
	s.grpcServer.GracefulStop()

	return nil
}

// New creates a new instance of the gRPC server.
func New(c *container.Container) Server {
	srv := &server{
		logger:  c.Logger,
		useCase: c.UseCase,
	}

	return srv
}
