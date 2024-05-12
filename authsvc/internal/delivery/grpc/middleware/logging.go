package middleware

import (
	"context"
	"fmt"
	"google.golang.org/grpc/peer"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log/slog"
)

// LoggingInterceptor is a gRPC interceptor for logging incoming requests and outgoing responses.
func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	// Convert request to string for logging
	requestString := fmt.Sprintf("%v", req)

	// Get the peer information from the context
	p, ok := peer.FromContext(ctx)
	var source string
	if ok {
		if addr, ok := p.Addr.(*net.TCPAddr); ok {
			// addr is the remote address of the client
			source = fmt.Sprintf("%s:%d", addr.IP, addr.Port)
		}
	}

	// Proceed with the request
	resp, err := handler(ctx, req)

	// Capture response status
	statusCode := status.Code(err)

	// Convert response to string for logging
	var responseString string
	if resp != nil {
		responseString = fmt.Sprintf("%v", resp)
	}

	// Calculate elapsed time
	elapsed := time.Since(start)

	// Log the request and response details
	slog.With(slog.String("type", "tdr")).
		// With("request_id", r.Context().Value(DefaultRequestIDKey)).
		With(slog.Group("request",
			slog.String("source", source),
			slog.String("method", "gRPC"),
			slog.String("path", info.FullMethod),
			// slog.Any("headers", r.Header),
			slog.Any("body", requestString),
		)).With(slog.Group("response",
		slog.Int("code", int(statusCode)),
		slog.String("body", responseString),
		slog.String("elapsed", fmt.Sprintf("%v", elapsed)),
	)).Info("gRPC request complete")

	return resp, err
}
