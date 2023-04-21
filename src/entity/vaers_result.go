package entity

type VaersResult struct {
	Id        int64  `json:"id"`
	VaccineId int64  `json:"vaccineId"`
	Name      string `json:"name"`
	SymptomId int64  `json:"symptomId"`
	Symptom   string `json:"symptom"`
	Total     int64  `json:"total"`
}

func (v *VaersResult) TableName() string {
	return "vaers_result"
}
