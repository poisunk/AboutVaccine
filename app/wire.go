//go:build wireinject
// +build wireinject

package main

import (
	"about-vaccine/src/base/dao"
	"about-vaccine/src/controller"
	"about-vaccine/src/repo"
	"about-vaccine/src/router"
	"about-vaccine/src/service"
	"github.com/google/wire"
)

func InitApplication(dsn string) (*router.APIRouter, error) {
	wire.Build(
		router.ProviderSetRouter,
		controller.ProviderSetController,
		service.ProviderSetService,
		repo.ProviderSetRepo,
		dao.ProviderSetDao,
	)
	return &router.APIRouter{}, nil
}
