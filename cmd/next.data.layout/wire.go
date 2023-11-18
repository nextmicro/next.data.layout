//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/nextmicro/logger"
	"next.data.layout/internal/biz"
	"next.data.layout/internal/conf"
	"next.data.layout/internal/data"
	"next.data.layout/internal/server"
	"next.data.layout/internal/service"
	"next.data.layout/internal/svc"

	"github.com/google/wire"
)

// wireApp init next application.
func wireApp(*conf.Data, logger.Logger) (*Injector, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, svc.ProviderSet, InjectorSet, newApp))

	return new(Injector), nil, nil
}
