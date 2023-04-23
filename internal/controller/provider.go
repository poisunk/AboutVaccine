package controller

import "github.com/google/wire"

var ProviderSetController = wire.NewSet(
	NewUserController,
	NewVaccineController,
	NewVaersController,
	NewAdverseEventController,
	NewOAETermController,
)
