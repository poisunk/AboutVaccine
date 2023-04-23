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
	Uid                 *int64     `json:"uid"`
	Code                string     `json:"code"`
	Name                string     `json:"name"`
	Sex                 string     `json:"sex"`
	Birth               *time.Time `json:"birth"`
	Phone               string     `json:"phone"`
	Address             string     `json:"address"`
	OnsetDate           *time.Time `json:"onsetDate"`
	CreateDate          time.Time  `xorm:"created" json:"createDate"`
	Description         string     `json:"description"`
	TreatmentDepartment string     `json:"treatmentDepartment"`
	Rapporteur          string     `json:"rapporteur"`
	RapporteurPhone     string     `json:"rapporteurPhone"`
	RapporteurAddress   string     `json:"rapporteurAddress"`
}

func (a *AdverseEvent) TableName() string {
	return "adverse_event"
}
