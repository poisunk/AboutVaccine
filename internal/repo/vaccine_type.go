package repo

import (
	"about-vaccine/internal/base/dao"
	"about-vaccine/internal/entity"
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

func (repo *VaccineTypeRepo) GetList(page, pageSize int) ([]*entity.VaccineType, int64, error) {
	typeList := make([]*entity.VaccineType, 0)
	total, err := repo.DB.Limit(pageSize, (page-1)*pageSize).FindAndCount(&typeList)
	if err != nil {
		return nil, 0, err
	}
	return typeList, total, nil
}

func (repo *VaccineTypeRepo) GetIdByType(typeStr string) (int64, bool, error) {
	var id int64
	has, err := repo.DB.Table(&entity.VaccineType{}).Where("type = ?", typeStr).Cols("id").Get(&id)
	if err != nil {
		return -1, false, err
	}
	return id, has, nil
}

func (repo *VaccineTypeRepo) Count() (total int64, err error) {
	total, err = repo.DB.Count(&entity.VaccineType{})
	if err != nil {
		return -1, err
	}
	return total, nil
}
