package models

import (
	"MyWeb/dao"
	"MyWeb/utile"
)

type VaersSymptomTerm struct {
	Id      int64  `json:"id"`
	Symptom string `json:"symptom"`
}

func (v *VaersSymptomTerm) TableName() string {
	return "vaers_symptom_term"
}

func GetVaersSymptomTermById(id int64) (v *VaersSymptomTerm, err error) {
	v = &VaersSymptomTerm{}
	if err = dao.DB.Where("id = ?", id).First(&v).Error; err != nil {
		return nil, err
	}
	return v, nil
}

func GetVaersSymptomTermList(keyword string, page, pageSize int) (v []*VaersSymptomTerm, count int64, err error) {
	db := dao.DB.Model(VaersSymptomTerm{}).Where("`symptom` LIKE ?", utile.HandleSearchWord(keyword)).Count(&count)
	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&v).Error
	return v, count, err
}
