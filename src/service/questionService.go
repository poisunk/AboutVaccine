package service

type QuestionService interface {
	GetQuestionList(qid int64) ([]*Question, error)
	GetQuestionTypeList() []string
	CreateQuestion(q []*Question) error
	DeleteQuestionById(id int64) error
}

var QuestionType = map[string]bool{
	"单选题": true,
	"多选题": true,
	"判断题": true,
	"填空题": true,
	"问答题": true,
	"简答题": true,
}

type Question struct {
	Id              int64    `json:"id"`
	QuestionnaireId int64    `json:"questionnaireId"`
	Type            string   `json:"type"`
	Content         string   `json:"content"`
	IsRequired      bool     `json:"isRequired"`
	Options         []string `json:"options"`
	Order           int64    `json:"order"`
}
