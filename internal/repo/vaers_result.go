package repo

import (
	"about-vaccine/internal/base/dao"
	"about-vaccine/internal/entity"
	"about-vaccine/internal/service/vaers"
)

type VaersResultRepo struct {
	DB *dao.DB
}

func NewVaersResultRepo(db *dao.DB) vaers.VaersResultRepo {
	return &VaersResultRepo{DB: db}
}

func (repo *VaersResultRepo) GetById(id int64) (*entity.VaersResult, bool, error) {
	r := &entity.VaersResult{}
	exist, err := repo.DB.ID(id).Get(r)
	if err != nil {
		return nil, false, err
	}
	return r, exist, nil
}

func (repo *VaersResultRepo) Get(vid, sid int64) (*entity.VaersResult, bool, error) {
	r := &entity.VaersResult{}
	exist, err := repo.DB.Where("vaccine_id = ? and symptom_id = ?", vid, sid).Get(r)
	if err != nil {
		return nil, false, err
	}
	return r, exist, nil
}

func (repo *VaersResultRepo) GetListByVaccineId(vid int64, page, pageSize int) ([]*entity.VaersResult, int64, error) {
	var list []*entity.VaersResult
	total, err := repo.DB.Where("vaccine_id = ?", vid).Limit(pageSize, (page-1)*pageSize).FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (repo *VaersResultRepo) GetListBySymptomId(sid int64, page, pageSize int) ([]*entity.VaersResult, int64, error) {
	var list []*entity.VaersResult
	total, err := repo.DB.Where("symptom_id = ?", sid).Limit(pageSize, (page-1)*pageSize).FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (repo *VaersResultRepo) SumByVaccineId(vid int64) (float64, error) {
	total, err := repo.DB.Where("vaccine_id = ?", vid).And("total > 1").
		Sum(&entity.VaersResult{}, "total")
	if err != nil {
		return 0, err
	}
	return total, err
}

func (repo *VaersResultRepo) SumBySymptomId(sid int64) (float64, error) {
	total, err := repo.DB.Where("symptom_id = ?", sid).And("total > 1").
		Sum(&entity.VaersResult{}, "total")
	if err != nil {
		return 0, err
	}
	return total, err
}

func (repo *VaersResultRepo) Sum() (float64, error) {
	total, err := repo.DB.Where("total > 1").Sum(&entity.VaersResult{}, "total")
	if err != nil {
		return 0, err
	}
	return total, err
}
