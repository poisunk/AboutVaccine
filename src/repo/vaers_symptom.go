package repo

import (
	"about-vaccine/src/base/dao"
	"about-vaccine/src/entity"
)

type VaersSymptomRepo struct {
	DB *dao.DB
}

func NewVaersSymptomRepo(db *dao.DB) *VaersSymptomRepo {
	return &VaersSymptomRepo{
		DB: db,
	}
}

func (repo *VaersSymptomRepo) GetByVaersId(vaersId int) ([]*entity.VaersResult, error) {
	list := make([]*entity.VaersResult, 0)
	err := repo.DB.Where("vaers_id = ?", vaersId).Find(&list)
	return list, err
}
