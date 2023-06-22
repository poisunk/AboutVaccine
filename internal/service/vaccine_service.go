package service

import (
	"errors"
	"log"
	"strconv"
	"vax/internal/schema"
	"vax/internal/service/vaccine"
)

type VaccineService struct {
	vaccineCommon     *vaccine.VaccineCommon
	vaccineTypeCommon *vaccine.VaccineTypeCommon
}

func NewVaccineService(
	vaccineCommon *vaccine.VaccineCommon,
	vaccineTypeCommon *vaccine.VaccineTypeCommon,
) *VaccineService {
	return &VaccineService{
		vaccineCommon:     vaccineCommon,
		vaccineTypeCommon: vaccineTypeCommon,
	}
}

func (s *VaccineService) Get(idStr string) (*schema.VaccineInfo, error) {
	id, _ := strconv.ParseInt(idStr, 10, 64)
	v, has, err := s.vaccineCommon.Get(id)
	if err != nil || !has {
		return nil, errors.New("疫苗不存在")
	}
	return v, nil
}

func (s *VaccineService) GetByName(keyword string, page, pageSize int) ([]*schema.VaccineBriefInfo, int64, error) {
	vs, total, err := s.vaccineCommon.GetList(keyword, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取Vaccine失败")
	}
	return vs, total, nil
}

func (s *VaccineService) GetByType(typeStr string, page, pageSize int) ([]*schema.VaccineBriefInfo, int64, error) {
	tid, has, err := s.vaccineTypeCommon.GetIdByType(typeStr)
	if err != nil || has == false {
		log.Println(err.Error())
		return nil, 0, errors.New("type不存在")
	}
	vaccineList, total, err := s.vaccineCommon.GetListByType(tid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取VaccineType失败")
	}
	return vaccineList, total, nil
}

func (s *VaccineService) GetTypeList(page, pageSize int) ([]*schema.VaccineTypeInfo, int64, error) {
	list, total, err := s.vaccineTypeCommon.GetTypeList(page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取VaccineType失败")
	}
	return list, total, nil
}

func (s *VaccineService) GetTypeDetailInfo(id int64) (*schema.VaccineTypeDetailInfo, error) {
	v, has, err := s.vaccineTypeCommon.Get(id)
	if err != nil || !has {
		return nil, errors.New("type不存在")
	}
	return v, nil
}
