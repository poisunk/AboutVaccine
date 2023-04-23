package entity

import (
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

func (a *AdverseVaccine) TableName() string {
	return "adverse_vaccine"
}
