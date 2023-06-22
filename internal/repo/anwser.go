package repo

import (
	"vax/internal/base/dao"
	"vax/internal/entity"
)

type AnswerRepo struct {
	DB *dao.DB
}

func NewAnswerRepo(db *dao.DB) *AnswerRepo {
	return &AnswerRepo{
		DB: db,
	}
}

func (a *AnswerRepo) CreateOne(answer *entity.Answer) error {
	_, err := a.DB.InsertOne(answer)
	return err
}

func (a *AnswerRepo) UpdateOne(answer *entity.Answer) error {
	_, err := a.DB.ID(answer.Id).Update(answer)
	return err
}

func (a *AnswerRepo) DeleteOne(id int64) error {
	_, err := a.DB.ID(id).Delete(&entity.Answer{})
	return err
}

func (a *AnswerRepo) GetOne(id int64) (*entity.Answer, error) {
	answer := &entity.Answer{}
	_, err := a.DB.ID(id).Get(answer)
	return answer, err
}

func (a *AnswerRepo) GetByIssueId(issueId int64) ([]*entity.Answer, error) {
	var answers []*entity.Answer
	err := a.DB.Where("issue_id = ?", issueId).Find(&answers)
	return answers, err
}
