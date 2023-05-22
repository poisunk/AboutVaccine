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

func (repo *AdverseVaccineRepo) GetListByEventId(eventId int64) ([]*entity.AdverseVaccine, error) {
	var list []*entity.AdverseVaccine
	err := repo.DB.Where("event_id = ?", eventId).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *AdverseVaccineRepo) GetListByOAEId(oid int64, page, pageSize int) ([]*entity.AdverseVaccine, error) {
	var list []*entity.AdverseVaccine
	err := repo.DB.SQL("SELECT DISTINCT vaccine_id FROM adverse_vaccine WHERE event_id IN (SELECT DISTINCT event_id FROM adverse_symptom WHERE oae_id = ?) LIMIT ? OFFSET ?",
		oid, pageSize, (page-1)*pageSize).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *AdverseVaccineRepo) CreateList(vaccine []*entity.AdverseVaccine) error {
	_, err := repo.DB.Insert(vaccine)
	return err
}
