package repo

import (
	"about-vaccine/internal/base/dao"
	"about-vaccine/internal/entity"
)

type IssueRepo struct {
	DB *dao.DB
}

func NewIssueRepo(db *dao.DB) *IssueRepo {
	return &IssueRepo{DB: db}
}

func (r *IssueRepo) CreateOne(issue *entity.Issue) error {
	_, err := r.DB.Insert(issue)
	return err
}

func (r *IssueRepo) UpdateOne(issue *entity.Issue) error {
	_, err := r.DB.ID(issue.Id).Update(issue)
	return err
}

func (r *IssueRepo) DeleteOne(id int64) error {
	_, err := r.DB.ID(id).Delete(entity.Issue{})
	return err
}

func (r *IssueRepo) GetOne(id int64) (*entity.Issue, bool, error) {
	issue := &entity.Issue{}
	has, err := r.DB.ID(id).Get(issue)
	if err != nil {
		return nil, false, err
	}
	return issue, has, nil
}

func (r *IssueRepo) GetList(page, pageSize int) ([]*entity.Issue, error) {
	var issues []*entity.Issue
	err := r.DB.Limit(pageSize, (page-1)*pageSize).Find(&issues)
	return issues, err
}

func (r *IssueRepo) GetByUid(uid int64, page, pageSize int) ([]*entity.Issue, error) {
	var issues []*entity.Issue
	err := r.DB.Where("uid = ?", uid).Limit(pageSize, (page-1)*pageSize).Find(&issues)
	return issues, err
}

func (r *IssueRepo) GetBySimilarTitle(keyword string, page, pageSize int) ([]*entity.Issue, error) {
	var issues []*entity.Issue
	err := r.DB.Where("title like ?", "%"+keyword+"%").Limit(pageSize, (page-1)*pageSize).Find(&issues)
	return issues, err
}
