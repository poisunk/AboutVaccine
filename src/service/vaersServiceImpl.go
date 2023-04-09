package service

import (
	"MyWeb/models"
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
	result = &VaersResult{
		Symptom: r.Symptom,
		Vaccine: r.Name,
		Total:   r.Total,
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
		list[i] = &VaersResult{
			Symptom: r.Symptom,
			Vaccine: r.Name,
			Total:   r.Total,
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
		list[i] = &VaersResult{
			Symptom: r.Symptom,
			Vaccine: r.Name,
			Total:   r.Total,
		}
	}
	total, err = models.CountVaersResultBySymptomId(sid)
	return list, total, nil
}
