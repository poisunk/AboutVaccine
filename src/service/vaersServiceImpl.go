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

var total_abcd int64 = 0

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
		ID:               v.ID,
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

	// 1. 查询目标疫苗信息
	vax, err := vs.GetVaersVaxTerm(vaccineId)
	if err != nil {
		return nil, err
	}
	// 2. 查询所有接种了该疫苗的VaersId数据
	vaersId, err := vs.GetVaersIdListByVaxId(vaccineId)
	if err != nil {
		return nil, err
	}
	// 3. 查询所有接种该疫苗后出现的不良反应，得到a, c
	var symptom *VaersSymptom
	symptomMap := make(map[*VaersSymptom]int64)
	var total_ac int64 = 0
	for _, v := range vaersId {
		sList, err := vs.GetVaersSymptomListByVaersId(v)
		if err != nil {
			return nil, err
		}
		for _, s := range sList {
			if s.SymptomId == symptomId && symptom == nil {
				symptom = s
			}
			symptomMap[s] = symptomMap[s] + 1
			total_ac++
		}
	}
	// 4. 查询所有接种该疫苗出现的不良反应的总出现次数，得到b
	var total_ab int64 = 0
	vaersId, err = vs.GetVaersIdListBySymptomId(symptomId)
	if err != nil {
		return nil, err
	}
	for _, v := range vaersId {
		total, err := vs.CountVaersVaxByVaersId(v)
		if err != nil {
			return nil, err
		}
		total_ab += total
	}
	// 5. 查询所有不良反应的数量，得到d
	total_abcd, err = vs.GetTotalABCD()
	if err != nil {
		return nil, err
	}
	// 6. 组装数据
	v := symptomMap[symptom]
	a, b, c, d := v, total_ab-v, total_ac-v, total_abcd-total_ab-total_ac+v
	prr := (float64(a) / float64(b+c)) / (float64(b) / float64(b+d))
	chi := float64((a*b-b*c)*(a*b-b*c)*(total_abcd)) / float64(total_ab*total_ac*(c+d)*(b+d))
	result = &VaersResult{
		Vaccine: vax.Name,
		Symptom: symptom.Symptom,
		Total:   v,
		Prr:     prr,
		Chi:     chi,
	}
	return nil, nil
}

func (vs *VaersServiceImpl) GetVaersResultsByVaccineId(vid int64) (list []*VaersResult, err error) {
	// 总共需要四个数据a, b, c, d
	// 			 	目标疫苗		其他疫苗
	// 目标不良反应	a			b
	// 其他不良反应	c			d

	// 1. 查询目标疫苗信息
	vax, err := vs.GetVaersVaxTerm(vid)
	if err != nil {
		return nil, err
	}
	// 2. 查询所有接种了该疫苗的VaersId数据
	vaersId, err := vs.GetVaersIdListByVaxId(vid)
	if err != nil {
		return nil, err
	}
	// 3. 查询所有接种该疫苗后出现的不良反应，得到a, c
	symptomCountMap := make(map[*VaersSymptom]int64)
	var total_ac int64 = 0
	for _, v := range vaersId {
		sList, err := vs.GetVaersSymptomListByVaersId(v)
		if err != nil {
			return nil, err
		}
		for _, s := range sList {
			symptomCountMap[s] = symptomCountMap[s] + 1
			total_ac++
		}
	}
	// 4. 查询每种接种该疫苗出现的不良反应的总出现次数，得到b
	totalABMap := make(map[*VaersSymptom]int64)
	for k, _ := range symptomCountMap {
		vaersId, err = vs.GetVaersIdListBySymptomId(k.SymptomId)
		if err != nil {
			return nil, err
		}
		var total int64
		for _, v := range vaersId {
			t, err := vs.CountVaersVaxByVaersId(v)
			if err != nil {
				return nil, err
			}
			total += t
		}
		totalABMap[k] = total
	}
	// 5. 查询所有不良反应的数量，得到d
	total_abcd, err = vs.GetTotalABCD()
	if err != nil {
		return nil, err
	}
	// 6. 组装数据
	list = make([]*VaersResult, 0, len(symptomCountMap))
	for k, v := range symptomCountMap {
		a, b, c, d := v, totalABMap[k]-v, total_ac-v, total_abcd-totalABMap[k]-total_ac+v
		prr := (float64(a) / float64(b+c)) / (float64(b) / float64(b+d))
		chi := float64((a*b-b*c)*(a*b-b*c)*(total_abcd)) / float64(totalABMap[k]*total_ac*(c+d)*(b+d))
		list = append(list, &VaersResult{
			Vaccine: vax.Name,
			Symptom: k.Symptom,
			Total:   v,
			Prr:     prr,
			Chi:     chi,
		})
	}
	return list, nil
}

func (vs *VaersServiceImpl) GetVaersResultsBySId(sid int64) (list []*VaersResult, err error) {
	// 总共需要四个数据a, b, c, d
	// 			 	目标疫苗		其他疫苗
	// 目标不良反应	a			b
	// 其他不良反应	c			d

	// 1. 查询目标不良反应
	symptom, err := vs.GetVaersSymptomTerm(sid)
	if err != nil {
		return nil, err
	}
	// 2. 查询所有含有该不良反应的VaersId数据
	vaersId, err := vs.GetVaersIdListBySymptomId(sid)
	if err != nil {
		return nil, err
	}
	// 3. 查询这些VaersId接种的所有疫苗数据，得到a, b
	vaccineMap := make(map[*VaersVax]int64)
	var total_ab int64 = 0
	for _, v := range vaersId {
		vList, err := vs.GetVaersVaxListByVaersId(v)
		if err != nil {
			return nil, err
		}
		for _, v := range vList {
			vaccineMap[v] = vaccineMap[v] + 1
			total_ab++
		}
	}
	// 4. 查询所有含有该不良反应的Vaers接种疫苗的所有数据
	totalACMap := make(map[*VaersVax]int64)
	for k, _ := range vaccineMap {
		vaersId, err = vs.GetVaersIdListByVaxId(k.VaxId)
		if err != nil {
			return nil, err
		}
		var total int64
		for _, v := range vaersId {
			t, err := vs.CountSymptomByVaersId(v)
			if err != nil {
				return nil, err
			}
			total += t
		}
		totalACMap[k] = total
	}
	// 5. 查询所有不良反应的数量
	total_abcd, err = vs.GetTotalABCD()
	if err != nil {
		return nil, err
	}
	// 6. 组装数据
	list = make([]*VaersResult, 0, len(vaccineMap))
	for k, v := range vaccineMap {
		a, b, c, d := v, total_ab-v, totalACMap[k]-v, total_abcd-total_ab-totalACMap[k]+v
		prr := (float64(a) / float64(b+c)) / (float64(b) / float64(b+d))
		chi := float64((a*b-b*c)*(a*b-b*c)*(total_abcd)) / float64(total_ab*totalACMap[k]*(c+d)*(b+d))
		list = append(list, &VaersResult{
			Vaccine: k.Name,
			Symptom: symptom.Symptom,
			Total:   v,
			Prr:     prr,
			Chi:     chi,
		})
	}
	return list, nil
}

func (vs *VaersServiceImpl) GetTotalABCD() (total int64, err error) {
	if total_abcd == 0 {
		total_abcd, err = vs.CountSymptom()
		if err != nil {
			return 0, err
		}
	}
	return total_abcd, nil
}
