package repo

import (
	"about-vaccine/src/base/dao"
	"about-vaccine/src/entity"
)

type VaersResultRepo struct {
	DB *dao.DB
}

func NewVaersResultRepo(db *dao.DB) *VaersResultRepo {
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

func (repo *VaersResultRepo) GetListByVaccineId(vaccineId int64) (list []*entity.VaersResult, err error) {
	list = make([]*entity.VaersResult, 0)
	err = repo.DB.Where("vaccine_id = ?", vaccineId).Find(&list)
	return
}

func (repo *VaersResultRepo) GetListBySymptomId(symptomId int64) (list []*entity.VaersResult, err error) {
	list = make([]*entity.VaersResult, 0)
	err = repo.DB.Where("symptom_id = ?", symptomId).Find(&list)
	return
}

func (repo *VaersResultRepo) Count() (int64, error) {
	total, err := repo.DB.Count(&entity.VaersResult{})
	return total, err
}

func (repo *VaersResultRepo) CountByVaccineId(vaccineId int64) (int64, error) {
	total, err := repo.DB.Where("vaccine_id = ?", vaccineId).Count(&entity.VaersResult{})
	return total, err
}

func (repo *VaersResultRepo) CountBySymptomId(symptomId int64) (int64, error) {
	total, err := repo.DB.Where("symptom_id = ?", symptomId).Count(&entity.VaersResult{})
	return total, err
}
