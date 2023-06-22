package repo

import (
	"vax/internal/base/dao"
	"vax/internal/entity"
)

type VaersVaxRepo struct {
	DB *dao.DB
}

func NewVaersVaxRepo(db *dao.DB) *VaersVaxRepo {
	return &VaersVaxRepo{
		DB: db,
	}
}

func (repo *VaersVaxRepo) GetByVaersId(vaersId int) ([]*entity.VaersResult, error) {
	list := make([]*entity.VaersResult, 0)
	err := repo.DB.Where("vaers_id = ?", vaersId).Find(&list)
	return list, err
}
