package service

import (
	"github.com/google/wire"
	"vax/internal/service/adverse_report"
	"vax/internal/service/oae"
	"vax/internal/service/user"
	"vax/internal/service/vaccine"
	"vax/internal/service/vaers"
)

var ProviderSetService = wire.NewSet(
	NewUserService,
	user.NewUserCommon,
	NewVaccineService,
	vaccine.NewVaccineCommon,
	vaccine.NewVaccineTypeCommon,
	NewAdverseReportService,
	adverse_report.NewAdverseEventCommon,
	adverse_report.NewAdverseResultCommon,
	NewOaeTermService,
	oae.NewOAETermCommon,
	NewVaersService,
	vaers.NewVaersResultCommon,
	vaers.NewVaersVaxCommon,
	vaers.NewVaersSymptomCommon,
)
