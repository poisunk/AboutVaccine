package entity

type AdverseSymptom struct {
	Id      int64  `xorm:"notnull pk autoincr INT(11) id" json:"id"`
	EventId int64  `xorm:"notnull INT(11) event_id" json:"eventId"`
	Symptom string `xorm:"null VARCHAR(255) symptom" json:"symptom"`
	OAEId   *int64 `xorm:"null INT(11) oae_id" json:"oaeId"`
	OAETerm string `xorm:"null VARCHAR(255) oae_term" json:"oaeTerm"`
}

func (a *AdverseSymptom) TableName() string {
	return "adverse_symptom"
}
