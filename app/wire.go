//go:build wireinject
// +build wireinject

package main

import (
	"about-vaccine/internal/base/dao"
	"about-vaccine/internal/controller"
	"about-vaccine/internal/repo"
	"about-vaccine/internal/router"
	"about-vaccine/internal/service"
	"github.com/google/wire"
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
