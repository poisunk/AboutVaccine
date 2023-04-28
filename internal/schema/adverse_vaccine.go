package schema

import "time"

type AdverseVaccineInfo struct {
	Type          string     `json:"type"`
	Manufacturer  string     `json:"manufacturer"`
	Name          string     `json:"name"`
	VaccinateDate *time.Time `json:"vaccinateDate"`
	Dose          string     `json:"dose"`
	Route         string     `json:"route"`
	Site          string     `json:"site"`
}

type AdverseVaccineAdd struct {
	VaccineId     int64      `json:"vaccineId"`
	VaccinateDate *time.Time `json:"vaccinateDate"`
	Dose          string     `json:"dose"`
	Route         string     `json:"route"`
	Site          string     `json:"site"`
}
