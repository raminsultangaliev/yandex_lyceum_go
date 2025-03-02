package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"order_service/internal/service"
	"order_service/pkg/logger"
	test "order_service/pkg/api/test/api"
)

func main() {
	ctx := context.Background()
	ctx, _ = logger.New(ctx)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := service.New()
	server := grpc.NewServer()
	test.RegisterOrderServiceServer(server, srv)

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		logger.GetLoggerFromCtx(ctx).Info(ctx, "failed to serve:", zap.Error(err))
	}
}
