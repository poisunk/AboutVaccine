package schama

import (
	"about-vaccine/internal/entity"
	"about-vaccine/internal/utile"
	"time"
)

type AdverseEvent struct {
	Id                  int64             `json:"id"`
	Uid                 *int64            `json:"uid"`
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
	SymptomList         []*AdverseSymptom `json:"symptomList"`
}

type AdverseVaccine struct {
	Id            int64      `json:"id"`
	VaccineId     int64      `json:"vaccineId"`
	Type          string     `json:"type"`
	Manufacturer  string     `json:"manufacturer"`
	Name          string     `json:"name"`
	VaccinateDate *time.Time `json:"vaccinateDate"`
	Dose          string     `json:"dose"`
	Route         string     `json:"route"`
	Site          string     `json:"site"`
}

type AdverseSymptom struct {
	Id      int64  `json:"id"`
	EventId int64  `json:"eventId"`
	Symptom string `json:"symptom"`
	OaeId   *int64 `json:"oaeId"`
}

func (a *AdverseEvent) ToEntity() *entity.AdverseEvent {
	event := &entity.AdverseEvent{}
	_ = utile.StructConv(a, event)
	return event
}

func (a *AdverseVaccine) ToEntity() *entity.AdverseVaccine {
	vaccine := &entity.AdverseVaccine{}
	_ = utile.StructConv(a, vaccine)
	return vaccine
}

func (a *AdverseSymptom) ToEntity() *entity.AdverseSymptom {
	symptom := &entity.AdverseSymptom{}
	_ = utile.StructConv(a, symptom)
	return symptom
}

func (a *AdverseEvent) GetFromEntity(e *entity.AdverseEvent) {
	_ = utile.StructConv(e, a)
}

func (a *AdverseVaccine) GetFromEntity(e *entity.AdverseVaccine, v *Vaccine) {
	_ = utile.StructConv(e, a)
	a.Type = v.Type
	a.Name = v.ProductName
	a.Manufacturer = v.ProductionCompany
}

func (a *AdverseSymptom) GetFromEntity(e *entity.AdverseSymptom) {
	_ = utile.StructConv(e, a)
}
