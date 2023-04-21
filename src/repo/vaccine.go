package repo

import (
	"about-vaccine/src/base/dao"
	"about-vaccine/src/entity"
)

type VaccineRepo struct {
	DB *dao.DB
}

func NewVaccineRepo(db *dao.DB) *VaccineRepo {
	return &VaccineRepo{
		DB: db,
	}
}

func (repo *VaccineRepo) GetByID(id int64) (*entity.Vaccine, bool, error) {
	vaccine := &entity.Vaccine{}
	exist, err := repo.DB.ID(id).Get(vaccine)
	if err != nil {
		return nil, false, err
	}
	return vaccine, exist, nil
}

func (repo *VaccineRepo) GetListByProductName(name string, page, pageSize int) ([]*entity.Vaccine, error) {
	var vaccines []*entity.Vaccine
	err := repo.DB.Where("product_name = ?", name).Limit(pageSize, (page-1)*pageSize).Find(&vaccines)
	if err != nil {
		return nil, err
	}
	return vaccines, nil
}
