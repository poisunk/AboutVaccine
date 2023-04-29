package repo

import (
	"about-vaccine/internal/base/dao"
	"about-vaccine/internal/entity"
	"about-vaccine/internal/service/adverse_report"
)

type AdverseSymptomRepo struct {
	DB *dao.DB
}

func NewAdverseSymptomRepo(DB *dao.DB) adverse_report.AdverseSymptomRepo {
	return &AdverseSymptomRepo{
		DB: DB,
	}
}

func (repo *AdverseSymptomRepo) GetListByEventId(eventId int64) ([]*entity.AdverseSymptom, error) {
	var list []*entity.AdverseSymptom
	err := repo.DB.Where("event_id = ?", eventId).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *AdverseSymptomRepo) CreateList(symptom []*entity.AdverseSymptom) error {
	_, err := repo.DB.Insert(symptom)
	return err
}
