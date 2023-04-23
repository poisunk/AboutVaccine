package schama

import (
	"about-vaccine/internal/entity"
	"about-vaccine/internal/utile"
	"time"
)

type Vaers struct {
	Id               int64           `json:"id"`
	VaersId          int64           `json:"vaersId"`
	CreateDate       time.Time       `json:"createDate"`
	Sex              string          `json:"sex"`
	SymptomText      string          `json:"symptomText"`
	Age              int64           `json:"age"`
	VaccinatedDate   *time.Time      `json:"vaccinatedDate"`
	OnsetDate        *time.Time      `json:"onsetDate"`
	VaersVaxList     []*VaersVax     `json:"vaersVaxList"`
	VaersSymptomList []*VaersSymptom `json:"vaersSymptomList"`
}

type VaersResult struct {
	SymptomId int64   `json:"symptomId"`
	Symptom   string  `json:"symptom"`
	VaccineId int64   `json:"vaccineId"`
	Vaccine   string  `json:"vaccine"`
	Total     int64   `json:"total"`
	Prr       float64 `json:"prr"`
	Chi       float64 `json:"chi"`
}

type VaersSymptom struct {
	Id        int64  `json:"id"`
	VaersId   int64  `json:"vaersId"`
	Symptom   string `json:"symptom"`
	SymptomId int64  `json:"symptomId"`
}

type VaersSymptomTerm struct {
	Id      int64  `json:"id"`
	Symptom string `json:"symptom"`
}

type VaersVax struct {
	Id           int64  `json:"id"`
	VaersId      int64  `json:"vaersId"`
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
	Dose         string `json:"dose"`
	Route        string `json:"route"`
	Site         string `json:"site"`
	VaxId        int64  `json:"vaxId"`
}

type VaersVaxTerm struct {
	Id           int64  `json:"id"`
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
}

func (v *VaersResult) GetFormEntity(e *entity.VaersResult) {
	v.SymptomId = e.SymptomId
	v.Symptom = e.Symptom
	v.VaccineId = e.VaccineId
	v.Vaccine = e.Name
	v.Total = e.Total
}

func (v *VaersSymptomTerm) GetFormEntity(e *entity.VaersSymptomTerm) {
	_ = utile.StructConv(e, v)
}

func (v *VaersVaxTerm) GetFormEntity(e *entity.VaersVaxTerm) {
	_ = utile.StructConv(e, v)
}
