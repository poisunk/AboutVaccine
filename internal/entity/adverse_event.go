package entity

import (
	"time"
)

const (
	Male    = "M"
	Female  = "F"
	UnKnown = "N"
)

type AdverseEvent struct {
	Id                  int64      `xorm:"notnull pk autoincr INT(11) id" json:"id"`
	Uid                 *int64     `xorm:"null INT(11) uid" json:"uid"`
	Code                string     `xorm:"null VARCHAR(255) code" json:"code"`
	Name                string     `xorm:"null VARCHAR(255) name" json:"name"`
	Sex                 string     `xorm:"null VARCHAR(255) sex" json:"sex"`
	Birth               *time.Time `xorm:"null DATETIME birth" json:"birth"`
	Phone               string     `xorm:"null VARCHAR(255) phone" json:"phone"`
	Address             string     `xorm:"null VARCHAR(255) address" json:"address"`
	OnsetDate           *time.Time `xorm:"null DATETIME onset_date" json:"onsetDate"`
	CreateDate          time.Time  `xorm:"created DATETIME create_date" json:"createDate"`
	Description         string     `xorm:"not null VARCHAR(255) description" json:"description"`
	TreatmentDepartment string     `xorm:"null VARCHAR(255) treatment_department" json:"treatmentDepartment"`
	Rapporteur          string     `xorm:"null VARCHAR(255) rapporteur" json:"rapporteur"`
	RapporteurPhone     string     `xorm:"null VARCHAR(255) rapporteur_phone" json:"rapporteurPhone"`
	RapporteurAddress   string     `xorm:"null VARCHAR(255) rapporteur_address" json:"rapporteurAddress"`
}

func (a *AdverseEvent) TableName() string {
	return "adverse_event"
}
