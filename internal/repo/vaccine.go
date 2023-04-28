package repo

import (
	"about-vaccine/internal/base/dao"
	"about-vaccine/internal/entity"
	"about-vaccine/internal/service/vaccine"
	"about-vaccine/internal/utile"
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

func (repo *VaccineRepo) GetListBySimilarName(name string, page, pageSize int) ([]*entity.Vaccine, int64, error) {
	var vaccines []*entity.Vaccine
	// TODO
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

func (repo *VaccineRepo) GetSimpleListBySimilarName(name string, page, pageSize int) ([]*entity.Vaccine, int64, error) {
	var vaccines []*entity.Vaccine
	// TODO
	total, err := repo.DB.Where("产品名称 LIKE ?", utile.HandleSearchWord(name)).
		Limit(pageSize, (page-1)*pageSize).Cols("id", "type", "产品名称").FindAndCount(&vaccines)
	if err != nil {
		return nil, 0, err
	}
	return vaccines, total, nil
}

func (repo *VaccineRepo) Update(vaccine *entity.Vaccine) error {
	_, err := repo.DB.ID(vaccine.Id).Update(vaccine)
	return err
}
