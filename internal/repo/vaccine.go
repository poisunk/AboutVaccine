package repo

import (
	"vax/internal/base/dao"
	"vax/internal/entity"
	"vax/internal/service/vaccine"
	"vax/internal/utile"
)

type VaccineRepo struct {
	DB *dao.DB
}

func NewVaccineRepo(db *dao.DB) vaccine.VaccineRepo {
	return &VaccineRepo{
		DB: db,
	}
}

func (repo *VaccineRepo) Get(id int64) (*entity.Vaccine, bool, error) {
	v := &entity.Vaccine{}
	exist, err := repo.DB.ID(id).Get(v)
	if err != nil {
		return nil, false, err
	}
	return v, exist, nil
}

func (repo *VaccineRepo) GetBriefListBySimilarName(name string, page, pageSize int) ([]*entity.Vaccine, int64, error) {
	var vaccines []*entity.Vaccine
	total, err := repo.DB.Where("product_name LIKE ?", utile.HandleSearchWord(name)).
		Cols("id", "type", "product_name", "production_company").
		Limit(pageSize, (page-1)*pageSize).FindAndCount(&vaccines)
	if err != nil {
		return nil, 0, err
	}
	return vaccines, total, nil
}

func (repo *VaccineRepo) GetBriefListByType(tid int64, page, pageSize int) ([]*entity.Vaccine, int64, error) {
	var vaccines []*entity.Vaccine
	total, err := repo.DB.Where("tid = ?", tid).
		Cols("id", "type", "product_name", "production_company").
		Limit(pageSize, (page-1)*pageSize).FindAndCount(&vaccines)
	if err != nil {
		return nil, 0, err
	}
	return vaccines, total, nil
}

func (repo *VaccineRepo) GetSimpleListBySimilarName(name string, page, pageSize int) ([]*entity.Vaccine, int64, error) {
	var vaccines []*entity.Vaccine
	total, err := repo.DB.Where("product_name LIKE ?", utile.HandleSearchWord(name)).
		Limit(pageSize, (page-1)*pageSize).Cols("id", "type", "product_name").FindAndCount(&vaccines)
	if err != nil {
		return nil, 0, err
	}
	return vaccines, total, nil
}

func (repo *VaccineRepo) GetName(id int64) (string, bool, error) {
	v := &entity.Vaccine{}
	exist, err := repo.DB.ID(id).Cols("product_name").Get(v)
	if err != nil {
		return "", false, err
	}
	return v.ProductName, exist, nil
}

func (repo *VaccineRepo) Update(vaccine *entity.Vaccine) error {
	_, err := repo.DB.ID(vaccine.Id).Update(vaccine)
	return err
}
