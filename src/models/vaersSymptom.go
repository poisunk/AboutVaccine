package models

import "MyWeb/dao"

type VaersSymptom struct {
	Id        int64  `json:"id"`
	VaersId   int64  `json:"vaersId"`
	Symptom   string `json:"symptom"`
	SymptomId int64  `json:"symptomId"`
}

func (v *VaersSymptom) TableName() string {
	return "vaers_symptom"
}

func GetVaersSymptomById(id int64) (v *VaersSymptom, err error) {
	v = &VaersSymptom{}
	if err = dao.DB.Where("id = ?", id).First(&v).Error; err != nil {
		return nil, err
	}
	return v, nil
}

func GetVaersSymptomListByVaersId(id int64) (v []*VaersSymptom, err error) {
	if err = dao.DB.Where("vaers_id = ?", id).Find(&v).Error; err != nil {
		return nil, err
	}
	return v, nil
}

func GetVaersSymptomListBySymptomId(id int64, page, pageSize int) (v []*VaersSymptom, err error) {
	if err = dao.DB.Where("symptom_id = ?", id).Offset((page - 1) * pageSize).Limit(pageSize).Find(&v).Error; err != nil {
		return nil, err
	}
	return v, nil
}

func CountVaersSymptomByVaersId(id int64) (count int64, err error) {
	if err = dao.DB.Where("vaers_id = ?", id).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func CountVaersSymptom() (count int64, err error) {
	if err = dao.DB.Model(VaersSymptom{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetVaersIdListBySymptomId(id int64) (list []*VaersSymptom, err error) {
	if err = dao.DB.Select("vaers_id").Where("symptom_id = ?", id).Limit(10).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
