package entity

type AdverseSymptom struct {
	Id      int64  `json:"id"`
	EventId int64  `json:"eventId"`
	Symptom string `json:"symptom"`
	OaeId   *int64 `json:"oaeId"`
}

func (a *AdverseSymptom) TableName() string {
	return "adverse_symptom"
}
