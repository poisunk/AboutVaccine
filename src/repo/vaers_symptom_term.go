package repo

import (
	"about-vaccine/src/base/dao"
	"about-vaccine/src/entity"
)

type VaersSymptomTermRepo struct {
	DB *dao.DB
}

func NewVaersSymptomTermRepo(db *dao.DB) *VaersSymptomTermRepo {
	return &VaersSymptomTermRepo{
		DB: db,
	}
}

func (repo *VaersSymptomTermRepo) GetBySymptom(keyword string, page, pageSize int) ([]*entity.VaersResult, error) {
	list := make([]*entity.VaersResult, 0)
	err := repo.DB.Where("symptom LIKE ?", "%"+keyword+"%").Limit(pageSize, (page-1)*pageSize).Find(&list)
	return list, err
}
