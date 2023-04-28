package schema

type VaersResultInfo struct {
	SymptomId int64   `json:"symptomId"`
	Symptom   string  `json:"symptom"`
	VaccineId int64   `json:"vaccineId"`
	Vaccine   string  `json:"vaccine"`
	Total     int64   `json:"total"`
	Prr       float64 `json:"prr"`
	Chi       float64 `json:"chi"`
}
