package utils

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func WithRequestID(ctx context.Context, requestID string) context.Context {
	println("Adding request ID to context:", requestID)
	if requestID == "" {
		return ctx
	}
	md := metadata.Pairs("x-request-id", requestID)
	return metadata.NewOutgoingContext(ctx, md)
}
