package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nofendian17/openota/authsvc/internal/container"
	"github.com/nofendian17/openota/authsvc/internal/delivery/grpc"
)

// Server is responsible for starting and stopping the GRPC server.
type Server interface {
	Run(ctx context.Context) error
}

type server struct {
	cntr *container.Container
	grpc grpc.Server
}

// New creates a new instance of the Server.
func New(cntr *container.Container, grpc grpc.Server) Server {
	return &server{
		cntr: cntr,
		grpc: grpc,
	}
}

// Run starts the GRPC server and handles graceful shutdown.
func (s *server) Run(ctx context.Context) error {
	// Channel to communicate server errors
	serverErrCh := make(chan error, 1)

	// Start the gRPC server in a separate goroutine
	go func() {
		serverErrCh <- s.startGRPCServer(ctx)
	}()

	// Listen for signals to initiate graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	select {
	case sig := <-sigCh:
		s.cntr.Logger.Info(ctx, fmt.Sprintf("Received signal: %v. Starting graceful shutdown...", sig), nil)
		// Initiate graceful shutdown
		return s.stopGRPCServer(ctx)

	case err := <-serverErrCh:
		if err != nil {
			// Log or handle the server error as needed
			s.cntr.Logger.Error(ctx, "Failed to start gRPC server", err)
			return err
		}
	}

	return nil
}

// startGRPCServer starts the GRPC server.
func (s *server) startGRPCServer(ctx context.Context) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
			s.cntr.Logger.Warn(ctx, fmt.Sprintf("Recovered from panic: %v", r), nil)
		}
	}()

	err := s.grpc.Start(s.cntr.Config.Application.Port)
	if err != nil {
		s.cntr.Logger.Error(ctx, "Failed to start server", err)
		return err
	}
	return nil
}

// stopGRPCServer stops the GRPC server gracefully.
func (s *server) stopGRPCServer(ctx context.Context) error {
	err := s.grpc.Stop(ctx)
	if err != nil {
		s.cntr.Logger.Error(ctx, fmt.Sprintf("Error during graceful shutdown: %v", err), nil)
		return err
	}
	return err
}
