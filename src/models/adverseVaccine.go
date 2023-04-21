package models

import (
	"about-vaccine/dao"
	"time"
)

type AdverseVaccine struct {
	Id             int64      `json:"id"`
	AdverseEventId int64      `json:"adverseEventId"`
	VaccineId      int64      `json:"vaccineId"`
	VaccinateDate  *time.Time `json:"vaccinateDate"`
	Dose           string     `json:"dose"`
	Route          string     `json:"route"`
	Site           string     `json:"site"`
}

func (a AdverseVaccine) TableName() string {
	return "adverse_vaccine"
}

func CreateAdverseVaccine(vaccine *AdverseVaccine) error {
	if err := dao.DB.Create(vaccine).Error; err != nil {
		return err
	}
	return nil
}

func CreateAdverseVaccineList(list []*AdverseVaccine) error {
	db := dao.DB.Begin()
	for _, v := range list {
		if err := db.Create(v).Error; err != nil {
			return err
		}
	}
	db.Commit()
	return nil
}

func GetAdverseVaccineListByVid(vid int64) (v []*AdverseVaccine, err error) {
	if err = dao.DB.Find(&v, "adverse_event_id = ?", vid).Error; err != nil {
		return nil, err
	}
	return v, nil
}

func GetAdverseVaccineById(id int64) (v *AdverseVaccine, err error) {
	v = new(AdverseVaccine)
	if err = dao.DB.Where("id = ?", id).First(v).Error; err != nil {
		return nil, err
	}
	return v, nil
}

func DeleteAdverseVaccineById(id int64) error {
	if err := dao.DB.Delete(AdverseVaccine{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
