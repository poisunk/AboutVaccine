package schema

type AdverseResultInfo struct {
	VaccineId   int64  `json:"vaccineId"`
	VaccineName string `json:"vaccineName"`
	OAEId       int64  `json:"oaeId"`
	OAETerm     string `json:"oaeTerm"`
	Total       int64  `json:"total"`
	Prr         string `json:"prr"`
	Chi         string `json:"chi"`
}
