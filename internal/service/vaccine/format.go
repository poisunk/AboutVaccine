package vaccine

import (
	"vax/internal/entity"
	"vax/internal/schema"
)

func (vc *VaccineCommon) FormatVaccineInfo(entity *entity.Vaccine) *schema.VaccineInfo {
	vaccine := schema.VaccineInfo{
		Id:                entity.Id,
		Type:              entity.Type,
		RegisterNumber:    entity.RegisterNumber,
		ProductName:       entity.ProductName,
		EnglishName:       entity.EnglishName,
		TradeName:         entity.TradeName,
		Dosage:            entity.Dosage,
		Specification:     entity.Specification,
		Owner:             entity.Owner,
		OwnerAddress:      entity.OwnerAddress,
		ProductionCompany: entity.ProductionCompany,
		ApprovalDate:      entity.ApprovalDate,
		ProductionAddress: entity.ProductionAddress,
		ProductionClass:   entity.ProductionClass,
		OriginalNumber:    entity.OriginalNumber,
		DrugCode:          entity.DrugCode,
		DrugCodeNote:      entity.DrugCodeNote,
	}
	return &vaccine
}

func (vc *VaccineCommon) FormatVaccineSimpleInfo(entity *entity.Vaccine) *schema.VaccineSimpleInfo {
	vaccine := schema.VaccineSimpleInfo{
		Id:   entity.Id,
		Type: entity.Type,
		Name: entity.ProductName,
	}
	return &vaccine
}

func (vc *VaccineTypeCommon) FormatVaccineTypeInfo(entity *entity.VaccineType) *schema.VaccineTypeInfo {
	vaccine := schema.VaccineTypeInfo{
		Id:   entity.Id,
		Type: entity.Type,
	}
	return &vaccine
}

func (vc *VaccineTypeCommon) FormatVaccineTypeDetailInfo(entity *entity.VaccineType) *schema.VaccineTypeDetailInfo {
	vaccine := schema.VaccineTypeDetailInfo{
		Id:                  entity.Id,
		Type:                entity.Type,
		DiseaseIntroduction: entity.DiseaseIntroduction,
		PreventiveMeasures:  entity.PreventiveMeasures,
		Target:              entity.Target,
		VaccinationBan:      entity.VaccinationBan,
		AdverseEvent:        entity.AdverseEvent,
	}
	return &vaccine
}

func (vc *VaccineCommon) FormatVaccineBriefInfo(entity *entity.Vaccine) *schema.VaccineBriefInfo {
	vaccine := schema.VaccineBriefInfo{
		Id:                entity.Id,
		Type:              entity.Type,
		ProductName:       entity.ProductName,
		ProductionCompany: entity.ProductionCompany,
	}
	return &vaccine
}
