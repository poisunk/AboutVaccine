package repo

import (
	"about-vaccine/src/base/dao"
	"about-vaccine/src/entity"
	"about-vaccine/src/utile"
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

func (repo *VaccineRepo) GetListByProductName(name string, page, pageSize int) ([]*entity.Vaccine, int64, error) {
	var vaccines []*entity.Vaccine
	total, err := repo.DB.Where("产品名称 LIKE ?", utile.HandleSearchWord(name)).
		Limit(pageSize, (page-1)*pageSize).FindAndCount(&vaccines)
	if err != nil {
		return nil, 0, err
	}
	return vaccines, total, nil
}

func (repo *VaccineRepo) GetListByType(tid int64, page, pageSize int) ([]*entity.Vaccine, int64, error) {
	var vaccines []*entity.Vaccine
	total, err := repo.DB.Where("tid = ?", tid).
		Limit(pageSize, (page-1)*pageSize).FindAndCount(&vaccines)
	if err != nil {
		return nil, 0, err
	}
	return vaccines, total, nil
}

func (repo *VaccineRepo) Update(vaccine *entity.Vaccine) error {
	_, err := repo.DB.ID(vaccine.Id).Update(vaccine)
	return err
}
