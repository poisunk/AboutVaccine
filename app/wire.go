//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"vax/internal/base/dao"
	"vax/internal/config"
	"vax/internal/controller"
	"vax/internal/repo"
	"vax/internal/router"
	"vax/internal/server"
	"vax/internal/service"
)

func InitApplication(serverConf *config.Server, databaseConf *config.Database) (*server.AppServer, error) {
	wire.Build(
		server.ProviderSetServer,
		router.ProviderSetRouter,
		controller.ProviderSetController,
		service.ProviderSetService,
		repo.ProviderSetRepo,
		dao.ProviderSetDao,
	)
	return &server.AppServer{}, nil
}
