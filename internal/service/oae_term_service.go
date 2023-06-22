package service

import (
	"errors"
	"log"
	"strconv"
	"vax/internal/entity"
	"vax/internal/service/oae"
)

type OaeTermService struct {
	common *oae.OAETermCommon
}

func NewOaeTermService(common *oae.OAETermCommon) *OaeTermService {
	return &OaeTermService{
		common: common,
	}
}

func (s *OaeTermService) GetByIRI(iri string) (*entity.OAETerm, error) {
	o, _, err := s.common.GetByIRI(iri)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取OAETerm失败")
	}
	return o, nil
}

func (s *OaeTermService) GetBySimilarLabel(label string, page, pageSize int) ([]*entity.OAETerm, int64, error) {
	o, total, err := s.common.GetBySimilarLabel(label, page, pageSize)
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

func (s *OaeTermService) GetByID(idStr string) (*entity.OAETerm, error) {
	if idStr == "" {
		return nil, errors.New("id参数不能为空")
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, errors.New("id格式错误")
	}
	o, exist, err := s.common.GetByID(id)
	if err != nil {
		return nil, errors.New("获取OAETerm失败")
	}
	if !exist {
		return nil, errors.New("目标OAETerm不存在")
	}
	return o, nil
}

func (s *OaeTermService) Reverse(list []*entity.OAETerm) []*entity.OAETerm {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}
