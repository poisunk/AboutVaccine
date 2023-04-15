package models

import "MyWeb/dao"

type AdverseSymptom struct {
	Id      int64  `json:"id"`
	EventId int64  `json:"eventId"`
	Symptom string `json:"symptom"`
	OaeId   int64  `json:"oaeId"`
}

func (a *AdverseSymptom) TableName() string {
	return "adverse_symptom"
}

func CreateAdverseSymptom(a *AdverseSymptom) error {
	return dao.DB.Create(a).Error
}

func CreateAdverseSymptomList(a []*AdverseSymptom) error {
	db := dao.DB.Begin()
	for _, v := range a {
		if err := db.Create(v).Error; err != nil {
			return err
		}
	}
	db.Commit()
	return nil
}

func GetAdverseSymptom(id int64) (*AdverseSymptom, error) {
	var a AdverseSymptom
	err := dao.DB.Where("id = ?", id).First(&a).Error
	return &a, err
}

func GetAdverseSymptomByEventId(eventId int64) ([]*AdverseSymptom, error) {
	var a []*AdverseSymptom
	err := dao.DB.Where("event_id = ?", eventId).Find(&a).Error
	return a, err
}
