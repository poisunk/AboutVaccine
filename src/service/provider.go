package service

import "github.com/google/wire"

var ProviderSetService = wire.NewSet(
	NewUserService,
	NewVaccineService,
	NewAdverseEventService,
	NewOaeTermService,
	NewVaersService,
)
