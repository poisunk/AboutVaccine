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

func (repo *VaersSymptomTermRepo) GetByName(keyword string, page, pageSize int) ([]*entity.VaersSymptomTerm, int64, error) {
	list := make([]*entity.VaersSymptomTerm, 0)
	total, err := repo.DB.Where("symptom LIKE ?", "%"+keyword+"%").Limit(pageSize, (page-1)*pageSize).FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}
	return list, total, err
}
