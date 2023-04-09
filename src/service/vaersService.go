package service

import "time"

type VaersService interface {
	GetVaersByVaersId(vid int64) (vaers *Vaers, err error)
	GetVaersResults(vaccineId, symptomId int64) (list *VaersResult, err error)
	GetVaersResultsByVaccineId(vid int64, page int, pageSize int) (list []*VaersResult, total int64, err error)
	GetVaersResultsBySymptomId(sid int64, page int, pageSize int) (list []*VaersResult, total int64, err error)
}

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
	Symptom string  `json:"symptom"`
	Vaccine string  `json:"vaccine"`
	Total   int64   `json:"total"`
	Prr     float64 `json:"prr"`
	Chi     float64 `json:"chi"`
}
