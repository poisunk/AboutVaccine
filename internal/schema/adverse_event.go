package schema

import (
	"time"
)

type AdverseEventInfo struct {
	Id                  int64                 `json:"id"`
	UserName            string                `json:"userName"`
	Code                string                `json:"code"`
	Name                string                `json:"name"`
	Sex                 string                `json:"sex"`
	Birth               *time.Time            `json:"birth"`
	Phone               string                `json:"phone"`
	Address             string                `json:"address"`
	OnsetDate           *time.Time            `json:"onsetDate"`
	CreateDate          time.Time             `json:"createDate"`
	Description         string                `json:"description"`
	TreatmentDepartment string                `json:"treatmentDepartment"`
	Rapporteur          string                `json:"rapporteur"`
	RapporteurPhone     string                `json:"rapporteurPhone"`
	RapporteurAddress   string                `json:"rapporteurAddress"`
	VaccineList         []*AdverseVaccineInfo `json:"vaccineList"`
	SymptomList         []*AdverseSymptomInfo `json:"symptomList"`
}

type AdverseEventBriefInfo struct {
	Id          int64                 `json:"id"`
	UserName    string                `json:"userName"`
	CreateDate  time.Time             `json:"createDate"`
	Description string                `json:"description"`
	SymptomList []*AdverseSymptomInfo `json:"symptomList"`
}

type AdverseEventAdd struct {
	Code                string               `json:"code"`
	Name                string               `json:"name"`
	Sex                 string               `json:"sex"`
	Birth               *time.Time           `json:"birth"`
	Phone               string               `json:"phone"`
	Address             string               `json:"address"`
	OnsetDate           *time.Time           `json:"onsetDate"`
	Description         string               `json:"description"`
	TreatmentDepartment string               `json:"treatmentDepartment"`
	Rapporteur          string               `json:"rapporteur"`
	RapporteurPhone     string               `json:"rapporteurPhone"`
	RapporteurAddress   string               `json:"rapporteurAddress"`
	VaccineList         []*AdverseVaccineAdd `json:"vaccineList"`
	SymptomList         []*AdverseSymptomAdd `json:"symptomList"`
}
