package entity

import "time"

type Questionnaire struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerId     int64     `json:"ownerId"`
	CreateTime  time.Time `json:"createTime"`
}

func (q *Questionnaire) TableName() (name string) {
	return "questionnaire"
}
