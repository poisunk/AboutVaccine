package oae

import "about-vaccine/internal/entity"

type OAETermRepo interface {
	GetByIRI(string) (*entity.OAETerm, bool, error)
	GetBySimilarLabel(keyword string, page, pageSize int) ([]*entity.OAETerm, int64, error)
	GetName(int64) (string, bool, error)
}

type OAETermCommon struct {
	oaeTermRepo OAETermRepo
}

func NewOAETermCommon(oaeTermRepo OAETermRepo) *OAETermCommon {
	return &OAETermCommon{
		oaeTermRepo: oaeTermRepo,
	}
}

func (c *OAETermCommon) GetByIRI(iri string) (*entity.OAETerm, bool, error) {
	return c.oaeTermRepo.GetByIRI(iri)
}

func (c *OAETermCommon) GetBySimilarLabel(keyword string, page, pageSize int) ([]*entity.OAETerm, int64, error) {
	return c.oaeTermRepo.GetBySimilarLabel(keyword, page, pageSize)
}

func (c *OAETermCommon) GetName(oid int64) (string, bool, error) {
	return c.oaeTermRepo.GetName(oid)
}
