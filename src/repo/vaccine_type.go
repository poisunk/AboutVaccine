package repo

import (
	"about-vaccine/src/base/dao"
	"about-vaccine/src/entity"
)

type VaccineTypeRepo struct {
	DB *dao.DB
}

func NewVaccineTypeRepo(db *dao.DB) *VaccineTypeRepo {
	return &VaccineTypeRepo{DB: db}
}

func (repo *VaccineTypeRepo) GetById(id int64) (*entity.VaccineType, bool, error) {
	v := &entity.VaccineType{}
	exist, err := repo.DB.ID(id).Get(v)
	if err != nil {
		return nil, false, err
	}
	return v, exist, nil
}

func (repo *VaccineTypeRepo) GetList(page, pageSize int) (typeList []*entity.VaccineType, err error) {
	err = repo.DB.Limit(pageSize, (page-1)*pageSize).Find(&typeList)
	if err != nil {
		return nil, err
	}
	return
}

func (repo *VaccineTypeRepo) Count() (total int64, err error) {
	total, err = repo.DB.Count(&entity.VaccineType{})
	if err != nil {
		return -1, err
	}
	return total, nil
}
