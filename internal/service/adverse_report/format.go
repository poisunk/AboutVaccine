package adverse_report

import (
	"about-vaccine/internal/entity"
	"about-vaccine/internal/schema"
)

func (a *AdverseReportCommon) FormatInfo(
	event *entity.AdverseEvent,
	vaccines []*entity.AdverseVaccine,
	symptoms []*entity.AdverseSymptom,
) *schema.AdverseEventInfo {
	eventInfo := a.FormatEventInfo(event)
	vaccineInfos := make([]*schema.AdverseVaccineInfo, 0)
	for _, v := range vaccines {
		vaccineInfo := a.FormatVaccineInfo(v, nil)
		vaccineInfos = append(vaccineInfos, vaccineInfo)
	}
	symptomInfos := make([]*schema.AdverseSymptomInfo, 0)
	for _, s := range symptoms {
		symptomInfo := a.FormatSymptomInfo(s, nil)
		symptomInfos = append(symptomInfos, symptomInfo)
	}
	eventInfo.VaccineList = vaccineInfos
	eventInfo.SymptomList = symptomInfos
	return eventInfo
}

func (a *AdverseReportCommon) FormatEventInfo(entity *entity.AdverseEvent) *schema.AdverseEventInfo {
	eventInfo := &schema.AdverseEventInfo{
		Id:                  entity.Id,
		Code:                entity.Code,
		Name:                entity.Name,
		Sex:                 entity.Sex,
		Birth:               entity.Birth,
		Phone:               entity.Phone,
		Address:             entity.Address,
		OnsetDate:           entity.OnsetDate,
		CreateDate:          entity.CreateDate,
		Description:         entity.Description,
		TreatmentDepartment: entity.TreatmentDepartment,
		Rapporteur:          entity.Rapporteur,
		RapporteurPhone:     entity.RapporteurPhone,
	}
	return eventInfo
}

func (a *AdverseReportCommon) FormatVaccineInfo(
	entity *entity.AdverseVaccine,
	handler func(info *schema.AdverseVaccineInfo),
) *schema.AdverseVaccineInfo {
	vaccineInfo := &schema.AdverseVaccineInfo{
		Id:            entity.VaccineId,
		VaccinateDate: entity.VaccinateDate,
		Dose:          entity.Dose,
		Route:         entity.Route,
		Site:          entity.Site,
	}
	if handler != nil {
		handler(vaccineInfo)
	}
	return vaccineInfo
}

func (a *AdverseReportCommon) FormatSymptomInfo(
	entity *entity.AdverseSymptom,
	handler func(info *schema.AdverseSymptomInfo),
) *schema.AdverseSymptomInfo {
	symptomInfo := &schema.AdverseSymptomInfo{
		Symptom: entity.Symptom,
		OAETerm: entity.OAETerm,
		OAEId:   entity.OAEId,
	}
	if handler != nil {
		handler(symptomInfo)
	}
	return symptomInfo
}

func (a *AdverseReportCommon) FormatEntity(schema *schema.AdverseEventAdd, uid *int64) (
	*entity.AdverseEvent, []*entity.AdverseVaccine, []*entity.AdverseSymptom) {
	event := &entity.AdverseEvent{
		Uid:                 uid,
		Code:                schema.Code,
		Name:                schema.Name,
		Sex:                 schema.Sex,
		Birth:               schema.Birth,
		Phone:               schema.Phone,
		Address:             schema.Address,
		OnsetDate:           schema.OnsetDate,
		Description:         schema.Description,
		TreatmentDepartment: schema.TreatmentDepartment,
		Rapporteur:          schema.Rapporteur,
		RapporteurPhone:     schema.RapporteurPhone,
		RapporteurAddress:   schema.RapporteurAddress,
	}
	vaccineList := make([]*entity.AdverseVaccine, 0, len(schema.VaccineList))
	for _, v := range schema.VaccineList {
		vaccine := &entity.AdverseVaccine{
			VaccineId:     v.VaccineId,
			VaccinateDate: v.VaccinateDate,
			Dose:          v.Dose,
			Route:         v.Route,
			Site:          v.Site,
		}
		vaccineList = append(vaccineList, vaccine)
	}
	symptomList := make([]*entity.AdverseSymptom, 0, len(schema.SymptomList))
	for _, s := range schema.SymptomList {
		symptom := &entity.AdverseSymptom{
			Symptom: s.Symptom,
			OAEId:   *(s.OAEId),
			OAETerm: s.OAETerm,
		}
		symptomList = append(symptomList, symptom)
	}
	return event, vaccineList, symptomList
}

func (a *AdverseReportCommon) FormatEventBriefInfo(entity *entity.AdverseEvent) *schema.AdverseEventBriefInfo {
	eventBriefInfo := &schema.AdverseEventBriefInfo{
		Id:          entity.Id,
		CreateDate:  entity.CreateDate,
		Description: entity.Description,
	}
	return eventBriefInfo
}
