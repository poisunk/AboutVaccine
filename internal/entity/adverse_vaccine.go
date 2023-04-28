package entity

import (
	"time"
)

type AdverseVaccine struct {
	Id            int64      `xorm:"notnull pk autoincr INT(11) id" json:"id"`
	EventId       int64      `xorm:"notnull INT(11) event_id" json:"adverseEventId"`
	VaccineId     int64      `xorm:"notnull INT(11) vaccine_id" json:"vaccineId"`
	VaccinateDate *time.Time `xorm:"null DATETIME vaccinate_date" json:"vaccinateDate"`
	Dose          string     `xorm:"null VARCHAR(255) dose" json:"dose"`
	Route         string     `xorm:"null VARCHAR(255) route" json:"route"`
	Site          string     `xorm:"null VARCHAR(255) site" json:"site"`
}

func (a *AdverseVaccine) TableName() string {
	return "adverse_vaccine"
}
