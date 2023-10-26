// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/nextmicro/logger"
	"github.com/nextmicro/next"
	"next.data.layout/internal/biz"
	"next.data.layout/internal/conf"
	"next.data.layout/internal/data"
	"next.data.layout/internal/server"
	"next.data.layout/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init next application.
func wireApp(confData *conf.Data, loggerLogger logger.Logger) (*next.Next, func(), error) {
	dataData, cleanup, err := data.NewData(confData, loggerLogger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, loggerLogger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, loggerLogger)
	greeterService := service.NewGreeterService(greeterUsecase)
	shortUrlService := service.NewShortUrlService()
	grpcServer := server.NewGRPCServer(greeterService, shortUrlService, loggerLogger)
	httpServer := server.NewHTTPServer(greeterService, shortUrlService, loggerLogger)
	nextNext, err := newApp(loggerLogger, grpcServer, httpServer)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	return nextNext, func() {
		cleanup()
	}, nil
}
