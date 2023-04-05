package models

import "MyWeb/dao"

type VaersSymptomTerm struct {
	Id      int    `json:"id"`
	Symptom string `json:"symptom"`
}

func (v *VaersSymptomTerm) TableName() string {
	return "vaers_symptom_terms"
}

func GetVaersSymptomTermById(id int) (v *VaersSymptomTerm, err error) {
	if err = dao.DB.Where("id = ?", id).First(&v).Error; err != nil {
		return nil, err
	}
	return v, nil
}

func GetVaersSymptomTermBySymptom(symptom string) (v []*VaersSymptomTerm, err error) {
	if err = dao.DB.Where("symptom LIKE ?", symptom).Find(&v).Error; err != nil {
		return nil, err
	}
	return v, nil
}
