package repo

import (
	"about-vaccine/internal/base/dao"
	"about-vaccine/internal/entity"
)

type VaersVaxTermRepo struct {
	DB *dao.DB
}

func NewVaersVaxTermRepo(db *dao.DB) *VaersVaxTermRepo {
	return &VaersVaxTermRepo{DB: db}
}

func (repo *VaersVaxTermRepo) GetByName(keyword string, page, pageSize int) ([]*entity.VaersVaxTerm, int64, error) {
	list := make([]*entity.VaersVaxTerm, 0)
	total, err := repo.DB.Where("name LIKE ?", "%"+keyword+"%").Limit(pageSize, (page-1)*pageSize).FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
