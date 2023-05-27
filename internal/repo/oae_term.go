package repo

import (
	"about-vaccine/internal/base/dao"
	"about-vaccine/internal/entity"
	"about-vaccine/internal/service/oae"
)

type OAETermRepo struct {
	DB *dao.DB
}

func NewOAETermRepo(db *dao.DB) oae.OAETermRepo {
	return &OAETermRepo{DB: db}
}

func (repo *OAETermRepo) GetByIRI(IRI string) (*entity.OAETerm, bool, error) {
	oaeTerm := &entity.OAETerm{}
	exist, err := repo.DB.Where("TermIRI = ?", IRI).Get(oaeTerm)
	if err != nil {
		return nil, false, err
	}
	return oaeTerm, exist, nil
}

func (repo *OAETermRepo) GetBySimilarLabel(label string, page, pageSize int) ([]*entity.OAETerm, int64, error) {
	oaeList := make([]*entity.OAETerm, 0)
	total, err := repo.DB.Where("TermLabel LIKE ?", "%"+label+"%").
		Limit(pageSize, (page-1)*pageSize).FindAndCount(&oaeList)
	if err != nil {
		return nil, 0, err
	}
	return oaeList, total, nil
}

func (repo *OAETermRepo) GetName(oid int64) (string, bool, error) {
	term := &entity.OAETerm{}
	exist, err := repo.DB.ID(oid).Get(term)
	if err != nil {
		return "", false, err
	}
	return term.TermLabel, exist, nil
}

func (repo *OAETermRepo) GetByID(id int64) (*entity.OAETerm, bool, error) {
	term := &entity.OAETerm{}
	exist, err := repo.DB.ID(id).Get(term)
	if err != nil {
		return nil, false, err
	}
	return term, exist, nil
}
