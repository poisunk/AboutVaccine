package schama

import (
	"about-vaccine/src/entity"
	"about-vaccine/src/utile"
	"database/sql"
	"time"
)

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
	SymptomList         []*AdverseSymptom `json:"symptomList"`
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

type AdverseSymptom struct {
	Id      int64  `json:"id"`
	EventId int64  `json:"eventId"`
	Symptom string `json:"symptom"`
	OaeId   int64  `json:"oaeId"`
}

func (a *AdverseEvent) ToEntity() *entity.AdverseEvent {
	event := &entity.AdverseEvent{}
	_ = utile.StructConv(a, event)
	if a.Uid != 0 {
		event.Uid = sql.NullInt64{
			Int64: a.Uid,
			Valid: true,
		}
	}
	if a.Birth != nil {
		event.Birth = sql.NullTime{
			Time:  *a.Birth,
			Valid: true,
		}
	}
	if a.OnsetDate != nil {
		event.OnsetDate = sql.NullTime{
			Time:  *a.OnsetDate,
			Valid: true,
		}
	}
	return event
}

func (a *AdverseVaccine) ToEntity() *entity.AdverseVaccine {
	vaccine := &entity.AdverseVaccine{}
	_ = utile.StructConv(a, vaccine)
	if a.VaccinateDate != nil {
		vaccine.VaccinateDate = sql.NullTime{
			Time:  *a.VaccinateDate,
			Valid: true,
		}
	}
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

func (a *AdverseVaccine) GetFromVaccine(e *entity.AdverseVaccine, v *Vaccine) {
	_ = utile.StructConv(e, a)
	a.Type = v.Type
	a.Name = v.ProductName
	a.Manufacturer = v.ProductionCompany
}

func (a *AdverseSymptom) GetFromSymptom(e *entity.AdverseSymptom) {
	_ = utile.StructConv(e, a)
}
