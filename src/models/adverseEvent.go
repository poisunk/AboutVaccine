package models

import (
	"MyWeb/dao"
	"time"
)

type AdverseEvent struct {
	ID                  int64      `json:"id"`
	Uid                 int64      `json:"uid"`
	Code                string     `json:"code"`
	Name                string     `json:"name"`
	Sex                 string     `json:"sex"`
	Birth               *time.Time `json:"birth"`
	Phone               string     `json:"phone"`
	Address             string     `json:"address"`
	OnsetDate           *time.Time `json:"onsetDate"`
	CreateDate          time.Time  `json:"createDate"`
	Description         string     `json:"description"`
	TreatmentDepartment string     `json:"treatmentDepartment"`
	Rapporteur          string     `json:"rapporteur"`
	RapporteurPhone     string     `json:"rapporteurPhone"`
	RapporteurAddress   string     `json:"rapporteurAddress"`
}

func (a AdverseEvent) TableName() string {
	return "adverse_event"
}

func CreateAdverseEvent(event *AdverseEvent) (err error) {
	if err = dao.DB.Create(event).Error; err != nil {
		return err
	}
	return nil
}

func GetAdverseEventById(id int64) (event *AdverseEvent, err error) {
	event = new(AdverseEvent)
	if err = dao.DB.Where("id = ?", id).First(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func GetAdverseEventList(page int64, pageSize int64) (list []*AdverseEvent, err error) {
	if err = dao.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func DeleteAdverseEventById(id int64) (err error) {
	if err = dao.DB.Delete(AdverseEvent{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func CountAdverseEvent() (count int64, err error) {
	if err = dao.DB.Model(&AdverseEvent{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
