package server

import (
	"github.com/nextmicro/logger"
	helloworldV1 "next.data.layout/api/helloworld/v1"
	shortUrlV1 "next.data.layout/api/shorturl/v1"
	"next.data.layout/internal/service"

	"github.com/nextmicro/next/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	greeter *service.GreeterService,
	shortUrl *service.ShortUrlService,
	logger logger.Logger) *http.Server {
	srv := http.NewServer()

	shortUrlV1.RegisterShortUrlHTTPServer(srv, shortUrl)
	helloworldV1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
