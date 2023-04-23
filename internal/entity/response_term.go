package entity

type ResponseTerm struct {
	Id         int64  `json:"id"`
	QuestionId int64  `json:"questionId"`
	Answer     string `json:"answer"`
	ResponseId int64  `json:"responseId"`
}

func (r *ResponseTerm) TableName() string {
	return "response_term"
}
