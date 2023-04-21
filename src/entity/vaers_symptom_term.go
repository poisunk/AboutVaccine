package entity

type VaersSymptomTerm struct {
	Id      int64  `json:"id"`
	Symptom string `json:"symptom"`
}

func (v *VaersSymptomTerm) TableName() string {
	return "vaers_symptom_term"
}
