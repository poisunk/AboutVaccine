package repo

import (
	"github.com/google/wire"
)

var ProviderSetRepo = wire.NewSet(
	NewAdverseEventRepo,
	NewAdverseSymptomRepo,
	NewAdverseVaccineRepo,
	NewOAETermRepo,
	NewUserRepo,
	NewVaccineRepo,
	NewVaccineTypeRepo,
	NewVaersRepo,
	NewVaersResultRepo,
	NewVaersSymptomRepo,
	NewVaersSymptomTermRepo,
	NewVaersVaxRepo,
	NewVaersVaxTermRepo,
	NewIssueRepo,
)
