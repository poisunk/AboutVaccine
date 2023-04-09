package models

import (
	"MyWeb/dao"
	"time"
)

type Vaers struct {
	Id             int64      `json:"id"`
	VaersId        int64      `json:"vaersId"`
	CreateDate     time.Time  `json:"createDate"`
	Sex            string     `json:"sex"`
	SymptomText    string     `json:"symptomText"`
	Age            int64      `json:"age"`
	VaccinatedDate *time.Time `json:"vaccinatedDate"`
	OnsetDate      *time.Time `json:"onsetDate"`
}

func (v *Vaers) TableName() string {
	return "vaers"
}

func GetVaersByVaersId(vid int64) (v *Vaers, err error) {
	v = &Vaers{}
	if err = dao.DB.Where("vaers_id = ?", vid).First(&v).Error; err != nil {
		return nil, err
	}
	return v, nil
}

func GetVaersById(id int64) (v *Vaers, err error) {
	v = &Vaers{}
	if err = dao.DB.Where("id = ?", id).First(&v).Error; err != nil {
		return nil, err
	}
	return v, nil
}
