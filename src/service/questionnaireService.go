package service

import "time"

type QuestionnaireService interface {
	GetQuestionnaireList(int, int) ([]*Questionnaire, int64, error)
	GetQuestionnaireByUid(uid int64) ([]*Questionnaire, error)
	CreateQuestionnaire(questionnaire *Questionnaire) error
	DeleteQuestionnaireByID(id int64) error
	GetQuestionnaireById(id int64) (*Questionnaire, error)
	GetQuestionnaireOwnerIdByID(id int64) (uid int64, err error)
}

type Questionnaire struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	OwnerId     int64     `json:"ownerId"`
	OwnerName   string    `json:"ownerName"`
	CreateTime  time.Time `json:"createTime"`
}
