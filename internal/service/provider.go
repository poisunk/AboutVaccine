package service

import (
	"about-vaccine/internal/service/adverse_report"
	"about-vaccine/internal/service/oae"
	"about-vaccine/internal/service/user"
	"about-vaccine/internal/service/vaccine"
	"about-vaccine/internal/service/vaers"
	"github.com/google/wire"
)

var ProviderSetService = wire.NewSet(
	NewUserService,
	user.NewUserCommon,
	NewVaccineService,
	vaccine.NewVaccineCommon,
	vaccine.NewVaccineTypeCommon,
	NewAdverseReportService,
	adverse_report.NewAdverseEventCommon,
	NewOaeTermService,
	oae.NewOAETermCommon,
	NewVaersService,
	vaers.NewVaersResultCommon,
	vaers.NewVaersVaxCommon,
	vaers.NewVaersSymptomCommon,
)
