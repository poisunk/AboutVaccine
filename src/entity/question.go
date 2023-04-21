package entity

type Question struct {
	Id              int64  `json:"id"`
	QuestionnaireId int64  `json:"questionnaireId"`
	Type            string `json:"type"`
	Content         string `json:"content"`
	IsRequired      bool   `json:"isRequired"`
	Options         string `json:"options"`
	Order           int64  `json:"order"`
}

func (q *Question) TableName() string {
	return "question"
}
