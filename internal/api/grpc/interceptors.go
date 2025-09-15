package grpcapi

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	h, err := handler(ctx, req)

	slog.Info(
		"request",
		"method", info.FullMethod,
		"duration", time.Since(start),
		"error", err,
	)

	return h, err
}

func ErrorInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	h, err := handler(ctx, req)

	if err != nil {
		slog.Error("error", "method", info.FullMethod, "error", err)
		return nil, err
	}

	return h, nil
}

func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			slog.Error("panic", "method", info.FullMethod, "panic", r)
		}
	}()

	return handler(ctx, req)
}
