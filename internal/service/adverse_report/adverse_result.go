package adverse_report

import (
	"about-vaccine/internal/schema"
	"about-vaccine/internal/service/oae"
	"about-vaccine/internal/service/vaccine"
)

type AdverseResultCommon struct {
	adverseEventRepo   AdverseEventRepo
	adverseSymptomRepo AdverseSymptomRepo
	adverseVaccineRepo AdverseVaccineRepo
	vaccineCommon      *vaccine.VaccineCommon
	oaeCommon          *oae.OAETermCommon
}

func NewAdverseResultCommon(
	adverseEventRepo AdverseEventRepo,
	adverseSymptomRepo AdverseSymptomRepo,
	adverseVaccineRepo AdverseVaccineRepo,
	vaccineCommon *vaccine.VaccineCommon,
	oaeCommon *oae.OAETermCommon,
) *AdverseResultCommon {
	return &AdverseResultCommon{
		adverseEventRepo:   adverseEventRepo,
		adverseSymptomRepo: adverseSymptomRepo,
		adverseVaccineRepo: adverseVaccineRepo,
		vaccineCommon:      vaccineCommon,
		oaeCommon:          oaeCommon,
	}
}

func (common *AdverseResultCommon) GetByVaccineId(vid int64, page, pageSize int) ([]*schema.AdverseResultInfo, error) {
	vaccineName, _, err := common.vaccineCommon.GetName(vid)
	if err != nil {
		return nil, err
	}
	total, err := common.adverseEventRepo.CountByVaccineId(vid)
	if err != nil {
		return nil, err
	}
	symptomList, err := common.adverseSymptomRepo.GetListByVaccineId(vid, page, pageSize)

	res := make([]*schema.AdverseResultInfo, 0, len(symptomList))
	for _, symptom := range symptomList {
		r := &schema.AdverseResultInfo{
			VaccineId:   vid,
			VaccineName: vaccineName,
			OAEId:       symptom.OAEId,
			OAETerm:     symptom.OAETerm,
			Total:       total,
		}
		res = append(res, r)
	}
	return res, nil
}

func (common *AdverseResultCommon) GetByOAEId(oid int64, page, pageSize int) ([]*schema.AdverseResultInfo, error) {
	oaeLabel, _, err := common.oaeCommon.GetName(oid)
	if err != nil {
		return nil, err
	}
	total, err := common.adverseEventRepo.CountByOAEId(oid)
	if err != nil {
		return nil, err
	}
	vaccineList, err := common.adverseVaccineRepo.GetListByOAEId(oid, page, pageSize)
	res := make([]*schema.AdverseResultInfo, 0, len(vaccineList))
	for _, v := range vaccineList {
		r := &schema.AdverseResultInfo{
			VaccineId: v.VaccineId,
			OAEId:     oid,
			OAETerm:   oaeLabel,
			Total:     total,
		}
		res = append(res, r)
	}
	return res, nil
}

func (common *AdverseResultCommon) calculatePrr(a, b, c, d float64) float64 {
	result := (a / (a + c)) / (b / (b + d))
	return result
}

func (common *AdverseResultCommon) calculateChi(a, b, c, d float64) float64 {
	total := a + b + c + d
	result := (total * (a*d - b*c) * (a*d - b*c)) / ((b + d) * (a + c) * (a + b) * (d + c))
	return result
}
