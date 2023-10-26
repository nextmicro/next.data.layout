package server

import (
	"github.com/nextmicro/logger"
	helloworldV1 "next.data.layout/api/helloworld/v1"
	shortUrlV1 "next.data.layout/api/shorturl/v1"
	"next.data.layout/internal/service"

	"github.com/nextmicro/next/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	greeter *service.GreeterService,
	shortUrl *service.ShortUrlService,
	logger logger.Logger) *grpc.Server {
	srv := grpc.NewServer()
	helloworldV1.RegisterGreeterServer(srv, greeter)
	shortUrlV1.RegisterShortUrlServer(srv, shortUrl)
	return srv
}
