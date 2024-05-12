package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthorizationInterceptor is a gRPC interceptor for authorization.
func AuthorizationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Extract metadata from the context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata not found in request")
	}

	// Extract authorization token from metadata
	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, status.Error(codes.Unauthenticated, "authorization token not found")
	}
	authToken := authHeader[0]

	// Perform authorization logic here, e.g., validate the token

	// If authorization fails, return an error
	if !isValidAuthToken(authToken) {
		return nil, status.Error(codes.PermissionDenied, "invalid authorization token")
	}

	// Proceed with the request if authorization succeeds
	return handler(ctx, req)
}

// isValidAuthToken is a placeholder function to validate the authorization token.
func isValidAuthToken(token string) bool {
	// Implement your authorization logic here, e.g., validate the token against a database or external service
	return token == "valid_token"
}
