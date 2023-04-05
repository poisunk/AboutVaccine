package models

import "MyWeb/dao"

type VaersSymptom struct {
	ID        int64  `json:"id"`
	VaersId   int64  `json:"vaersId"`
	Symptom   string `json:"symptom"`
	SymptomId int64  `json:"symptomId"`
}

func (v *VaersSymptom) TableName() string {
	return "vaers_symptom"
}

func GetVaersSymptomById(id int64) (v *VaersSymptom, err error) {
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
