//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"vax/internal/base/dao"
	"vax/internal/controller"
	"vax/internal/repo"
	"vax/internal/router"
	"vax/internal/service"
)

func InitApplication() (*router.APIRouter, error) {
	wire.Build(
		router.ProviderSetRouter,
		controller.ProviderSetController,
		service.ProviderSetService,
		repo.ProviderSetRepo,
		dao.ProviderSetDao,
	)
	return &router.APIRouter{}, nil
}
