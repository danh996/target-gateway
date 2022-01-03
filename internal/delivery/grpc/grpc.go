package grpc

import (
	"context"

	"gateway/internal/domain"

	"google.golang.org/grpc"
)

func NewGRPCServer(ctx context.Context, authDomain domain.AuthDomain) *grpc.Server {
	return nil
}
