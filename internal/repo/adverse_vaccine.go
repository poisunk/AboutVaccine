package repo

import (
	"about-vaccine/internal/base/dao"
	"about-vaccine/internal/entity"
	"about-vaccine/internal/service/adverse_report"
)

type AdverseVaccineRepo struct {
	DB *dao.DB
}

func NewAdverseVaccineRepo(DB *dao.DB) adverse_report.AdverseVaccineRepo {
	return &AdverseVaccineRepo{
		DB: DB,
	}
}

func (repo *AdverseVaccineRepo) GetById(id int64) (*entity.AdverseVaccine, bool, error) {
	vaccine := &entity.AdverseVaccine{}
	exist, err := repo.DB.ID(id).Get(vaccine)
	if err != nil {
		return nil, false, err
	}
	return vaccine, exist, nil
}

func (repo *AdverseVaccineRepo) GetListByEventId(eventId int64) ([]*entity.AdverseVaccine, error) {
	var list []*entity.AdverseVaccine
	//TODO
	err := repo.DB.Where("event_id = ?", eventId).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *AdverseVaccineRepo) Count() (int64, error) {
	total, err := repo.DB.Count(&entity.AdverseVaccine{})
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (repo *AdverseVaccineRepo) CreateList(vaccine []*entity.AdverseVaccine) error {
	_, err := repo.DB.Insert(vaccine)
	return err
}
