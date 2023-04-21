package repo

import (
	"about-vaccine/src/base/dao"
	"about-vaccine/src/entity"
)

type VaersRepo struct {
	DB *dao.DB
}

func NewVaersRepo(db *dao.DB) *VaersRepo {
	return &VaersRepo{
		DB: db,
	}
}

func (repo *VaersRepo) GetById(id int64) (*entity.Vaers, bool, error) {
	v := &entity.Vaers{}
	exist, err := repo.DB.ID(id).Get(v)
	if err != nil {
		return nil, false, err
	}
	return v, exist, nil
}

func (repo *VaersRepo) GetList(page, pageSize int) ([]*entity.Vaers, error) {
	var v []*entity.Vaers
	err := repo.DB.Limit(pageSize, (page-1)*pageSize).Find(v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
