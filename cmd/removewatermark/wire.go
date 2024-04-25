//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	removewatermarkBiz "remove-watermark/internal/biz"
	"remove-watermark/internal/conf"
	removewatermarkData "remove-watermark/internal/data"
	"remove-watermark/internal/server"
	removewatermarkService "remove-watermark/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		newApp,
		server.ProviderSet,
		removewatermarkData.ProviderSet,
		removewatermarkBiz.ProviderSet,
		removewatermarkService.ProviderSet,
	))
}
