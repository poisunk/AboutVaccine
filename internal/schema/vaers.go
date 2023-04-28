package schema

import (
	"time"
)

type VaersInfo struct {
	Id               int64               `json:"id"`
	VaersId          int64               `json:"vaersId"`
	CreateDate       time.Time           `json:"createDate"`
	Sex              string              `json:"sex"`
	SymptomText      string              `json:"symptomText"`
	Age              int64               `json:"age"`
	VaccinatedDate   *time.Time          `json:"vaccinatedDate"`
	OnsetDate        *time.Time          `json:"onsetDate"`
	VaersVaxList     []*VaersVaxInfo     `json:"vaersVaxList"`
	VaersSymptomList []*VaersSymptomInfo `json:"vaersSymptomList"`
}
