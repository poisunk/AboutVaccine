package service

import (
	"about-vaccine/src/entity"
	"about-vaccine/src/repo"
	"errors"
	"log"
)

type OaeTermService struct {
	OaeTermRepo *repo.OAETermRepo
}

func NewOaeTermService(oaeTermRepo *repo.OAETermRepo) *OaeTermService {
	return &OaeTermService{
		OaeTermRepo: oaeTermRepo,
	}
}

func (s *OaeTermService) GetByIRI(iri string) (*entity.OAETerm, error) {
	o, _, err := s.OaeTermRepo.GetByIRI(iri)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取OAETerm失败")
	}
	return o, nil
}

func (s *OaeTermService) GetByLabel(label string, page, pageSize int) ([]*entity.OAETerm, int64, error) {
	o, total, err := s.OaeTermRepo.GetByLabel(label, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取OAETerm失败")
	}
	return o, total, nil
}

func (s *OaeTermService) GetParents(IRI string) ([]*entity.OAETerm, error) {
	var list []*entity.OAETerm
	for len(IRI) != 0 {
		o, err := s.GetByIRI(IRI)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
		IRI = o.ParentTermIRI
	}
	list = s.Reverse(list)
	return list, nil
}

func (s *OaeTermService) Reverse(list []*entity.OAETerm) []*entity.OAETerm {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}
