package repo

import (
	"about-vaccine/internal/base/dao"
	"about-vaccine/internal/entity"
)

type AdverseSymptomRepo struct {
	DB *dao.DB
}

func NewAdverseSymptomRepo(DB *dao.DB) *AdverseSymptomRepo {
	return &AdverseSymptomRepo{
		DB: DB,
	}
}

func (repo *AdverseSymptomRepo) GetById(id int64) (*entity.AdverseSymptom, bool, error) {
	symptom := &entity.AdverseSymptom{}
	exist, err := repo.DB.ID(id).Get(symptom)
	if err != nil {
		return nil, false, err
	}
	return symptom, exist, nil
}

func (repo *AdverseSymptomRepo) GetByEventId(eventId int64) ([]*entity.AdverseSymptom, error) {
	var list []*entity.AdverseSymptom
	err := repo.DB.Where("event_id = ?", eventId).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *AdverseSymptomRepo) Count() (int64, error) {
	total, err := repo.DB.Count(&entity.AdverseSymptom{})
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (repo *AdverseSymptomRepo) Create(symptom ...*entity.AdverseSymptom) error {
	_, err := repo.DB.Insert(symptom)
	return err
}
