package entity

import "time"

type Response struct {
	Id              int64     `json:"id"`
	QuestionnaireId int64     `json:"questionnaireId"`
	UserId          int64     `json:"userId"`
	ResponseTime    time.Time `json:"responseTime"`
}

func (r *Response) TableName() string {
	return "response"
}
