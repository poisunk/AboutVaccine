package repo

import (
	"about-vaccine/internal/base/dao"
	"about-vaccine/internal/entity"
	"about-vaccine/internal/service/adverse_report"
)

type AdverseEventRepo struct {
	DB *dao.DB
}

func NewAdverseEventRepo(DB *dao.DB) adverse_report.AdverseEventRepo {
	return &AdverseEventRepo{
		DB: DB,
	}
}

func (repo *AdverseEventRepo) Get(id int64) (*entity.AdverseEvent, bool, error) {
	event := &entity.AdverseEvent{}
	exist, err := repo.DB.ID(id).Get(event)
	if err != nil {
		return nil, false, err
	}
	return event, exist, nil
}

func (repo *AdverseEventRepo) GetBriefList(page, pageSize int) ([]*entity.AdverseEvent, int64, error) {
	var list []*entity.AdverseEvent
	total, err := repo.DB.Limit(pageSize, (page-1)*pageSize).Cols("id", "uid", "create_date", "description").
		FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (repo *AdverseEventRepo) GetBriefListByUid(uid int64, page, pageSize int) ([]*entity.AdverseEvent, int64, error) {
	var list []*entity.AdverseEvent
	total, err := repo.DB.Where("uid = ?", uid).Cols("id", "uid", "create_date", "description").
		Limit(pageSize, (page-1)*pageSize).FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (repo *AdverseEventRepo) GetBriefListByKeyword(keyword string, page, pageSize int) ([]*entity.AdverseEvent, int64, error) {
	var list []*entity.AdverseEvent
	total, err := repo.DB.Where("description like ?", "%"+keyword+"%").
		Cols("id", "uid", "create_date", "description").
		Limit(pageSize, (page-1)*pageSize).FindAndCount(&list)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (repo *AdverseEventRepo) GetListByVaccineId(vid int64, page, pageSize int) ([]*entity.AdverseEvent, error) {
	var list []*entity.AdverseEvent
	err := repo.DB.SQL("SELECT id, uid, create_date, description FROM adverse_event WHERE id in (SELECT DISTINCT event_id FROM adverse_vaccine WHERE vaccine_id = ?) LIMIT ? OFFSET ?",
		vid, pageSize, (page-1)*pageSize).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *AdverseEventRepo) GetListByOAEId(oid int64, page, pageSize int) ([]*entity.AdverseEvent, error) {
	var list []*entity.AdverseEvent
	err := repo.DB.SQL("SELECT id, uid, create_date, description FROM adverse_event WHERE id in (SELECT DISTINCT event_id FROM adverse_symptom WHERE oae_id = ?) LIMIT ? OFFSET ?",
		oid, pageSize, (page-1)*pageSize).Find(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (repo *AdverseEventRepo) GetUid(id int64) (int64, bool, error) {
	var uid int64
	has, err := repo.DB.Table(&entity.AdverseEvent{}).ID(id).Cols("uid").Get(&uid)
	if err != nil {
		return 0, false, err
	}
	return uid, has, nil
}

func (repo *AdverseEventRepo) Create(event *entity.AdverseEvent) error {
	_, err := repo.DB.Insert(event)
	return err
}

func (repo *AdverseEventRepo) Delete(id int64) error {
	_, err := repo.DB.ID(id).Delete(&entity.AdverseEvent{})
	return err
}
