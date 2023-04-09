package service

import "time"

type AdverseEventService interface {
	CreateAdverseEvent(AdverseEvent) error
	GetAdverseEvent(int64) (*AdverseEvent, error)
	DeleteAdverseEvent(int64) error
	GetAdverseEventList(int64, int64) ([]*AdverseEvent, int64, error)
}

type AdverseEvent struct {
	Id                  int64             `json:"id"`
	Uid                 int64             `json:"uid"`
	Code                string            `json:"code"`
	Name                string            `json:"name"`
	Sex                 string            `json:"sex"`
	Birth               *time.Time        `json:"birth"`
	Phone               string            `json:"phone"`
	Address             string            `json:"address"`
	OnsetDate           *time.Time        `json:"onsetDate"`
	CreateDate          time.Time         `json:"createDate"`
	Description         string            `json:"description"`
	TreatmentDepartment string            `json:"treatmentDepartment"`
	Rapporteur          string            `json:"rapporteur"`
	RapporteurPhone     string            `json:"rapporteurPhone"`
	RapporteurAddress   string            `json:"rapporteurAddress"`
	VaccineList         []*AdverseVaccine `json:"vaccineList"`
}

type AdverseVaccine struct {
	Id            int64      `json:"id"`
	Type          string     `json:"type"`
	Manufacturer  string     `json:"manufacturer"`
	Name          string     `json:"name"`
	VaccinateDate *time.Time `json:"vaccinateDate"`
	Dose          string     `json:"dose"`
	Route         string     `json:"route"`
	Site          string     `json:"site"`
}
