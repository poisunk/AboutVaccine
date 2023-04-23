package entity

import (
	"database/sql"
	"time"
)

type Vaers struct {
	Id             int64        `json:"id"`
	VaersId        int64        `json:"vaersId"`
	CreateDate     time.Time    `json:"createDate"`
	Sex            string       `json:"sex"`
	SymptomText    string       `json:"symptomText"`
	Age            int64        `json:"age"`
	VaccinatedDate sql.NullTime `json:"vaccinatedDate"`
	OnsetDate      sql.NullTime `json:"onsetDate"`
}

func (v *Vaers) TableName() string {
	return "vaers"
}
