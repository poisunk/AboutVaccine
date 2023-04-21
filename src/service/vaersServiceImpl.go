package service

import (
	"about-vaccine/src/models"
	"errors"
	"log"
)

type VaersServiceImpl struct {
	VaersSymptomService
	VaersVaxService
}

func InitVaersService() VaersService {
	return &VaersServiceImpl{
		VaersSymptomService: InitVaersSymptomService(),
		VaersVaxService:     InitVaersVaxService(),
	}
}

func (vs *VaersServiceImpl) GetVaersByVaersId(vid int64) (vaers *Vaers, err error) {
	// 1. 查询Vaers数据
	v, err := models.GetVaersByVaersId(vid)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("查询Vaers失败")
	}
	// 2. 查询VaersVax数据
	vList, err := vs.GetVaersVaxListByVaersId(vid)
	if err != nil {
		return nil, err
	}
	// 3. 查询VaersSymptom数据
	sList, err := vs.GetVaersSymptomListByVaersId(vid)
	if err != nil {
		return nil, err
	}
	// 4. 组装数据
	vaers = &Vaers{
		Id:               v.Id,
		VaersId:          v.VaersId,
		CreateDate:       v.CreateDate,
		Sex:              v.Sex,
		SymptomText:      v.SymptomText,
		Age:              v.Age,
		VaccinatedDate:   v.VaccinatedDate,
		OnsetDate:        v.OnsetDate,
		VaersVaxList:     vList,
		VaersSymptomList: sList,
	}
	return vaers, nil
}

func (vs *VaersServiceImpl) GetVaersResults(vaccineId, symptomId int64) (result *VaersResult, err error) {
	// 总共需要四个数据a, b, c, d
	// 			 	目标疫苗		其他疫苗
	// 目标不良反应	a			b
	// 其他不良反应	c			d

	r, err := models.GetVaersResult(vaccineId, symptomId)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("查询VaersResult失败")
	}
	total_ac, _ := models.SumVaersResultTotalEqVaccineId(vaccineId)
	total_b, _ := models.SumVaersResultTotalEqSymptomIdAndNotEqVaccineId(symptomId, vaccineId)
	total_bd, _ := models.SumVaersResultTotalNotEqVaccineId(vaccineId)
	var prr float64
	var chi float64
	if total_ac == 0 || total_b == 0 || total_bd == 0 {
		prr = 0
		chi = 0
	} else {
		prr = (float64(r.Total) / float64(total_ac)) / (float64(total_b) / float64(total_bd))
		chi = calculateChiSquare(float64(r.Total), float64(total_ac), float64(total_b), float64(total_bd))
	}
	result = &VaersResult{
		Symptom: r.Symptom,
		Vaccine: r.Name,
		Total:   r.Total,
		Prr:     prr,
		Chi:     chi,
	}
	return result, nil
}

func (vs *VaersServiceImpl) GetVaersResultsByVaccineId(vid int64, page int, pageSize int) (list []*VaersResult, total int64, err error) {
	// 总共需要四个数据a, b, c, d
	// 			 	目标疫苗		其他疫苗
	// 目标不良反应	a			b
	// 其他不良反应	c			d

	rLisr, err := models.GetVaersResultListByVaccineId(vid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, total, errors.New("查询VaersResult失败")
	}
	list = make([]*VaersResult, len(rLisr))
	for i, r := range rLisr {
		total_ac, _ := models.SumVaersResultTotalEqVaccineId(vid)
		total_b, _ := models.SumVaersResultTotalEqSymptomIdAndNotEqVaccineId(r.SymptomId, vid)
		total_bd, _ := models.SumVaersResultTotalNotEqVaccineId(vid)
		var prr float64
		var chi float64
		if total_ac == 0 || total_b == 0 || total_bd == 0 {
			prr = 0
			chi = 0
		} else {
			prr = (float64(r.Total) / float64(total_ac)) / (float64(total_b) / float64(total_bd))
			chi = calculateChiSquare(float64(r.Total), float64(total_ac), float64(total_b), float64(total_bd))
		}
		list[i] = &VaersResult{
			Symptom: r.Symptom,
			Vaccine: r.Name,
			Total:   r.Total,
			Prr:     prr,
			Chi:     chi,
		}
	}
	total, err = models.CountVaersResultByVaccineId(vid)
	return list, total, nil
}

func (vs *VaersServiceImpl) GetVaersResultsBySymptomId(sid int64, page int, pageSize int) (list []*VaersResult, total int64, err error) {
	// 总共需要四个数据a, b, c, d
	// 			 	目标疫苗		其他疫苗
	// 目标不良反应	a			b
	// 其他不良反应	c			d

	rLisr, err := models.GetVaersResultListBySymptomId(sid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, total, errors.New("查询VaersResult失败")
	}
	list = make([]*VaersResult, len(rLisr))
	for i, r := range rLisr {
		total_ac, _ := models.SumVaersResultTotalEqVaccineId(r.VaccineId)
		total_b, _ := models.SumVaersResultTotalEqSymptomIdAndNotEqVaccineId(sid, r.VaccineId)
		total_bd, _ := models.SumVaersResultTotalNotEqVaccineId(r.VaccineId)
		var prr float64
		var chi float64
		if total_ac == 0 || total_b == 0 || total_bd == 0 {
			prr = 0
			chi = 0
		} else {
			prr = (float64(r.Total) / float64(total_ac)) / (float64(total_b) / float64(total_bd))
			chi = calculateChiSquare(float64(r.Total), float64(total_ac), float64(total_b), float64(total_bd))
		}
		list[i] = &VaersResult{
			Symptom: r.Symptom,
			Vaccine: r.Name,
			Total:   r.Total,
			Prr:     prr,
			Chi:     chi,
		}
	}
	total, err = models.CountVaersResultBySymptomId(sid)
	return list, total, nil
}

func calculateChiSquare(a, total_ac, b, total_bd float64) float64 {
	total_abcd := total_ac + total_bd
	d := total_bd - b
	c := total_ac - a
	result := (total_abcd * (a*d - b*c) * (a*d - b*c)) / ((total_bd * total_ac) * (a + b) * (d + c))
	return result
}
