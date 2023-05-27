package adverse_report

import (
	"about-vaccine/internal/entity"
	"about-vaccine/internal/schema"
	"about-vaccine/internal/service/oae"
	"about-vaccine/internal/service/vaccine"
	"sync"
)

type AdverseResultRepo interface {
	Count(vid, oid int64) (int64, error)
	CountByVaccineId(vid int64) (int64, error)
	CountByOAEId(oid int64) (int64, error)
}

type AdverseResultCommon struct {
	adverseEventRepo   AdverseEventRepo
	adverseSymptomRepo AdverseSymptomRepo
	adverseVaccineRepo AdverseVaccineRepo
	adverseResultRepo  AdverseResultRepo
	vaccineCommon      *vaccine.VaccineCommon
	oaeCommon          *oae.OAETermCommon
}

func NewAdverseResultCommon(
	adverseEventRepo AdverseEventRepo,
	adverseSymptomRepo AdverseSymptomRepo,
	adverseVaccineRepo AdverseVaccineRepo,
	adverseResultRepo AdverseResultRepo,
	vaccineCommon *vaccine.VaccineCommon,
	oaeCommon *oae.OAETermCommon,
) *AdverseResultCommon {
	return &AdverseResultCommon{
		adverseEventRepo:   adverseEventRepo,
		adverseSymptomRepo: adverseSymptomRepo,
		adverseVaccineRepo: adverseVaccineRepo,
		adverseResultRepo:  adverseResultRepo,
		vaccineCommon:      vaccineCommon,
		oaeCommon:          oaeCommon,
	}
}

func (common *AdverseResultCommon) GetByVaccineId(vid int64, page, pageSize int) ([]*schema.AdverseResultInfo, error) {
	vaccineName, _, err := common.vaccineCommon.GetName(vid)
	if err != nil {
		return nil, err
	}
	total, err := common.adverseEventRepo.Count()
	if err != nil {
		return nil, err
	}
	symptomList, err := common.adverseSymptomRepo.GetListByVaccineId(vid, page, pageSize)

	res := make([]*schema.AdverseResultInfo, 0, len(symptomList))
	var wg sync.WaitGroup
	for _, symptom := range symptomList {
		wg.Add(1)
		r := &schema.AdverseResultInfo{
			VaccineId:   vid,
			VaccineName: vaccineName,
			OAEId:       symptom.OAEId,
			OAETerm:     symptom.OAETerm,
		}
		go common.calculateCoreData(r, total, &wg)
		res = append(res, r)
	}
	wg.Wait()
	return res, nil
}

func (common *AdverseResultCommon) GetByOAEId(oid int64, page, pageSize int) ([]*schema.AdverseResultInfo, error) {
	oaeLabel, _, err := common.oaeCommon.GetName(oid)
	if err != nil {
		return nil, err
	}
	total, err := common.adverseEventRepo.Count()
	if err != nil {
		return nil, err
	}
	vaccineList, err := common.adverseVaccineRepo.GetListByOAEId(oid, page, pageSize)
	res := make([]*schema.AdverseResultInfo, 0, len(vaccineList))
	var wg sync.WaitGroup
	for _, v := range vaccineList {
		r := &schema.AdverseResultInfo{
			VaccineId: v.VaccineId,
			OAEId:     oid,
			OAETerm:   oaeLabel,
		}
		wg.Add(1)
		go func(v *entity.AdverseVaccine) {
			r.VaccineName, _, err = common.vaccineCommon.GetName(v.VaccineId)
			common.calculateCoreData(r, total, &wg)
		}(v)
		res = append(res, r)
	}
	wg.Wait()
	return res, nil
}

func (common *AdverseResultCommon) calculateCoreData(result *schema.AdverseResultInfo, total int64, wg *sync.WaitGroup) {
	defer wg.Done()
	result.Total, _ = common.adverseResultRepo.Count(result.VaccineId, result.OAEId)
	totalAB, _ := common.adverseResultRepo.CountByOAEId(result.OAEId)
	totalAC, _ := common.adverseResultRepo.CountByVaccineId(result.VaccineId)
	a := float64(result.Total)
	b, c := float64(totalAB)-a, float64(totalAC)-a
	d := float64(result.Total) - a - b - c
	result.Prr = common.calculatePrr(a, b, c, d)
	result.Chi = common.calculateChi(a, b, c, d)
}

func (common *AdverseResultCommon) calculatePrr(a, b, c, d float64) float64 {
	if a <= 0 || b <= 0 || c <= 0 || d <= 0 {
		return 0
	}
	result := (a / (a + c)) / (b / (b + d))
	return result
}

func (common *AdverseResultCommon) calculateChi(a, b, c, d float64) float64 {
	if a <= 0 || b <= 0 || c <= 0 || d <= 0 {
		return 0
	}
	total := a + b + c + d
	result := (total * (a*d - b*c) * (a*d - b*c)) / ((b + d) * (a + c) * (a + b) * (d + c))
	return result
}
