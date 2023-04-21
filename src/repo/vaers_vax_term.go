package repo

import (
	"about-vaccine/src/base/dao"
	"about-vaccine/src/entity"
)

type VaersVaxTermRepo struct {
	DB *dao.DB
}

func NewVaersVaxTermRepo(db *dao.DB) *VaersVaxTermRepo {
	return &VaersVaxTermRepo{DB: db}
}

func (repo *VaersVaxTermRepo) GerByVaccine(keyword string, page, pageSize int) ([]*entity.VaersVaxTerm, error) {
	list := make([]*entity.VaersVaxTerm, 0)
	err := repo.DB.Where("vaccine LIKE ?", "%"+keyword+"%").Limit(pageSize, (page-1)*pageSize).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
