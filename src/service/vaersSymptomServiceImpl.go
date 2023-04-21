package service

import (
	"about-vaccine/src/models"
	"errors"
	"log"
)

type VaersSymptomServiceImpl struct {
}

func InitVaersSymptomService() VaersSymptomService {
	return &VaersSymptomServiceImpl{}
}

func (vs *VaersSymptomServiceImpl) CountSymptomByVaersId(id int64) (count int64, err error) {
	count, err = models.CountVaersSymptomByVaersId(id)
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("查询symptom失败")
	}
	return count, nil
}

func (vs *VaersSymptomServiceImpl) GetVaersSymptomListByVaersId(vid int64) (list []*VaersSymptom, err error) {
	l, err := models.GetVaersSymptomListByVaersId(vid)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("查询symptom失败")
	}

	// 组装数据
	for _, v := range l {
		list = append(list, &VaersSymptom{
			Id:        v.Id,
			VaersId:   v.VaersId,
			Symptom:   v.Symptom,
			SymptomId: v.SymptomId,
		})
	}
	return list, nil
}

func (vs *VaersSymptomServiceImpl) GetVaersSymptomTermList(keyword string, page, pageSize int) (
	list []*VaersSymptomTerm, count int64, err error) {
	l, c, err := models.GetVaersSymptomTermList(keyword, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("查询symptom失败")
	}
	// 组装数据
	for _, v := range l {
		list = append(list, &VaersSymptomTerm{
			Id:      v.Id,
			Symptom: v.Symptom,
		})
	}
	return list, c, nil
}

func (vs *VaersSymptomServiceImpl) CountSymptom() (count int64, err error) {
	count, err = models.CountVaersSymptom()
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("查询symptom失败")
	}
	return count, nil
}

func (vs *VaersSymptomServiceImpl) GetVaersSymptomTerm(id int64) (term *VaersSymptomTerm, err error) {
	t, err := models.GetVaersSymptomTermById(id)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("查询symptom失败")
	}
	term = &VaersSymptomTerm{
		Id:      t.Id,
		Symptom: t.Symptom,
	}
	return term, nil
}

func (vs *VaersSymptomServiceImpl) GetVaersIdListBySymptomId(id int64) (list []int64, err error) {
	l, err := models.GetVaersIdListBySymptomId(id)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("查询symptom失败")
	}
	// 组装数据
	for _, v := range l {
		list = append(list, v.VaersId)
	}
	return list, nil
}
