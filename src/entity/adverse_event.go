package entity

import (
	"database/sql"
	"time"
)

const (
	Male    = "M"
	Female  = "F"
	UnKnown = "N"
)

type AdverseEvent struct {
	Id                  int64        `json:"id"`
	Uid                 int64        `json:"uid"`
	Code                string       `json:"code"`
	Name                string       `json:"name"`
	Sex                 string       `json:"sex"`
	Birth               sql.NullTime `json:"birth"`
	Phone               string       `json:"phone"`
	Address             string       `json:"address"`
	OnsetDate           sql.NullTime `json:"onsetDate"`
	CreateDate          time.Time    `json:"createDate"`
	Description         string       `json:"description"`
	TreatmentDepartment string       `json:"treatmentDepartment"`
	Rapporteur          string       `json:"rapporteur"`
	RapporteurPhone     string       `json:"rapporteurPhone"`
	RapporteurAddress   string       `json:"rapporteurAddress"`
}

func (a *AdverseEvent) TableName() string {
	return "adverse_event"
}
