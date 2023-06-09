package repo

import (
	"vax/internal/base/dao"
	"vax/internal/entity"
	"vax/internal/service/vaers"
)

type VaersVaxTermRepo struct {
	DB *dao.DB
}

func NewVaersVaxTermRepo(db *dao.DB) vaers.VaersVaxTermRepo {
	return &VaersVaxTermRepo{DB: db}
}

func (repo *VaersVaxTermRepo) GetListBySimilarName(keyword string, page, pageSize int) ([]*entity.VaersVaxTerm, int64, error) {
	list := make([]*entity.VaersVaxTerm, 0)
	total, err := repo.DB.Where("name LIKE ?", "%"+keyword+"%").Limit(pageSize, (page-1)*pageSize).FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
