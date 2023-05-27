package repo

import (
	"about-vaccine/internal/base/dao"
	"about-vaccine/internal/entity"
	"about-vaccine/internal/service/adverse_report"
)

type AdverseResultRepo struct {
	DB *dao.DB
}

func NewAdverseResultRepo(DB *dao.DB) adverse_report.AdverseResultRepo {
	return &AdverseResultRepo{
		DB: DB,
	}
}

func (repo *AdverseResultRepo) Count(vid, oid int64) (int64, error) {
	sql := "SELECT COUNT(DISTINCT(event_id)) FROM adverse_symptom WHERE oae_id = ? AND event_id in (SELECT event_id FROM adverse_vaccine WHERE vaccine_id = ?);"
	return repo.DB.SQL(sql, oid, vid).Count(&entity.AdverseSymptom{})
}

func (repo *AdverseResultRepo) CountByVaccineId(vid int64) (int64, error) {
	return repo.DB.Where("id in (SELECT DISTINCT event_id FROM adverse_vaccine WHERE vaccine_id = ?)", vid).Count(&entity.AdverseEvent{})
}

func (repo *AdverseResultRepo) CountByOAEId(oid int64) (int64, error) {
	return repo.DB.Where("id in (SELECT DISTINCT event_id FROM adverse_symptom WHERE oae_id = ?)", oid).Count(&entity.AdverseEvent{})
}
