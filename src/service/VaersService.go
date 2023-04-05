package service

import "time"

type VaersService interface {
	GetVaersByVaersId(vid int64) (list []*Vaers, err error)
}

type Vaers struct {
	ID             int64       `json:"id"`
	VaersId        int64       `json:"vaersId"`
	CreateDate     time.Time   `json:"createDate"`
	Sex            string      `json:"sex"`
	SymptomText    string      `json:"symptomText"`
	Age            int64       `json:"age"`
	VaccinatedDate *time.Time  `json:"vaccinatedDate"`
	OnsetDate      *time.Time  `json:"onsetDate"`
	VaersVaxList   []*VaersVax `json:"vaersVaxList"`
}
