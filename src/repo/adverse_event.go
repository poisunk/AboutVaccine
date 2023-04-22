package repo

import (
	"about-vaccine/src/base/dao"
	"about-vaccine/src/entity"
)

type AdverseEventRepo struct {
	DB *dao.DB
}

func NewAdverseEventRepo(DB *dao.DB) *AdverseEventRepo {
	return &AdverseEventRepo{
		DB: DB,
	}
}

func (repo *AdverseEventRepo) GetById(id int64) (*entity.AdverseEvent, bool, error) {
	event := &entity.AdverseEvent{}
	exist, err := repo.DB.ID(id).Get(event)
	if err != nil {
		return nil, false, err
	}
	return event, exist, nil
}

func (repo *AdverseEventRepo) GetList(page, pageSize int) ([]*entity.AdverseEvent, int64, error) {
	var list []*entity.AdverseEvent
	total, err := repo.DB.Limit(pageSize, (page-1)*pageSize).FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (repo *AdverseEventRepo) GetListByUid(uid int64, page, pageSize int) ([]*entity.AdverseEvent, int64, error) {
	var list []*entity.AdverseEvent
	total, err := repo.DB.Where("uid = ?", uid).Limit(pageSize, (page-1)*pageSize).FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (repo *AdverseEventRepo) Count() (int64, error) {
	total, err := repo.DB.Count(&entity.AdverseEvent{})
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (repo *AdverseEventRepo) Create(event ...*entity.AdverseEvent) error {
	_, err := repo.DB.Insert(event)
	return err
}

func (repo *AdverseEventRepo) Update(event *entity.AdverseEvent) error {
	_, err := repo.DB.ID(event.Id).Update(event)
	return err
}

func (repo *AdverseEventRepo) Delete(id int64) error {
	_, err := repo.DB.ID(id).Delete(&entity.AdverseEvent{})
	return err
}
