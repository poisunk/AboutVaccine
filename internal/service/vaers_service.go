package service

import (
	"errors"
	"log"
	"vax/internal/schema"
	"vax/internal/service/vaers"
)

type VaersService struct {
	vaersResultCommon  *vaers.VaersResultCommon
	vaersVaxCommon     *vaers.VaersVaxCommon
	vaersSymptomCommon *vaers.VaersSymptomCommon
}

func NewVaersService(
	vaersResultCommon *vaers.VaersResultCommon,
	vaersVaxCommon *vaers.VaersVaxCommon,
	vaersSymptomCommon *vaers.VaersSymptomCommon,
) *VaersService {
	return &VaersService{
		vaersResultCommon:  vaersResultCommon,
		vaersVaxCommon:     vaersVaxCommon,
		vaersSymptomCommon: vaersSymptomCommon,
	}
}

func (s *VaersService) GetResultByVaccineId(vid int64, page, pageSize int) ([]*schema.VaersResultInfo, int64, error) {
	list, total, err := s.vaersResultCommon.GetResultByVaccineId(vid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("查询失败")
	}
	return list, total, nil
}

func (s *VaersService) GetResultBySymptomId(sid int64, page, pageSize int) ([]*schema.VaersResultInfo, int64, error) {
	list, total, err := s.vaersResultCommon.GetResultBySymptomId(sid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("查询失败")
	}
	return list, total, nil
}

func (s *VaersService) GetResult(vid int64, sid int64) (*schema.VaersResultInfo, error) {
	v, has, err := s.vaersResultCommon.GetResult(vid, sid)
	if err != nil || !has {
		log.Println(err.Error())
		return nil, errors.New("记录不存在")
	}
	return v, nil
}

func (s *VaersService) GetVaccineTermList(keyword string, page, pageSize int) ([]*schema.VaersVaxTerm, int64, error) {
	list, total, err := s.vaersVaxCommon.GetVaxTermList(keyword, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("查询失败")
	}
	return list, total, nil
}

func (s *VaersService) GetSymptomTermList(keyword string, page, pageSize int) ([]*schema.VaersSymptomTerm, int64, error) {
	list, total, err := s.vaersSymptomCommon.GetSymptomTermList(keyword, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("查询失败")
	}
	return list, total, nil
}
