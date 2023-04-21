package entity

type VaersSymptom struct {
	Id        int64  `json:"id"`
	VaersId   int64  `json:"vaersId"`
	Symptom   string `json:"symptom"`
	SymptomId int64  `json:"symptomId"`
}

func (v *VaersSymptom) TableName() string {
	return "vaers_symptom"
}
