package repo

import (
	"about-vaccine/src/base/dao"
	"about-vaccine/src/entity"
)

type OAETermRepo struct {
	DB *dao.DB
}

func NewOAETermRepo(db *dao.DB) *OAETermRepo {
	return &OAETermRepo{DB: db}
}

func (repo *OAETermRepo) GetByIRI(IRI string) (*entity.OAETerm, bool, error) {
	oaeTerm := new(entity.OAETerm)
	exist, err := repo.DB.Where("TermIRI = ?", IRI).Get(oaeTerm)
	if err != nil {
		return nil, false, err
	}
	return oaeTerm, exist, nil
}

func (repo *OAETermRepo) GetByLabel(label string, page, pageSize int) ([]*entity.OAETerm, int64, error) {
	oaeList := make([]*entity.OAETerm, 0)
	total, err := repo.DB.Where("TermLabel LIKE ?", "%"+label+"%").
		Limit(pageSize, (page-1)*pageSize).FindAndCount(&oaeList)
	if err != nil {
		return nil, 0, err
	}
	return oaeList, total, nil
}

func (repo *OAETermRepo) CountByLabel(label string) (int64, error) {
	total, err := repo.DB.Where("TermLabel LIKE ?", "%"+label+"%").Count()
	if err != nil {
		return 0, err
	}
	return total, nil
}
