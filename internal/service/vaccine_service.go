package service

import (
	"about-vaccine/internal/entity"
	"about-vaccine/internal/repo"
	"about-vaccine/internal/schama"
	"errors"
	"log"
)

type VaccineService struct {
	VaccineRepo *repo.VaccineRepo
	TypeRepo    *repo.VaccineTypeRepo
}

func NewVaccineService(
	vaccineRepo *repo.VaccineRepo,
	typeRepo *repo.VaccineTypeRepo,
) *VaccineService {
	return &VaccineService{
		VaccineRepo: vaccineRepo,
		TypeRepo:    typeRepo,
	}
}

func (s *VaccineService) Get(id int64) (*schama.Vaccine, error) {
	v, has, err := s.VaccineRepo.GetByID(id)
	if err != nil || !has {
		return nil, errors.New("疫苗不存在")
	}
	if len(v.Type) == 0 {
		err := s.setupVaccineType(v)
		if err != nil {
			return nil, err
		}
	}
	vaccine := &schama.Vaccine{}
	vaccine.GetFormEntity(v)
	return vaccine, nil
}

func (s *VaccineService) GetByName(keyword string, page, pageSize int) ([]*schama.Vaccine, int64, error) {
	vs, total, err := s.VaccineRepo.GetListByProductName(keyword, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取Vaccine失败")
	}
	var list []*schama.Vaccine
	for _, v := range vs {
		if len(v.Type) == 0 {
			err := s.setupVaccineType(v)
			if err != nil {
				return nil, 0, err
			}
		}
		vaccine := &schama.Vaccine{}
		vaccine.GetFormEntity(v)
		list = append(list, vaccine)
	}
	return list, total, nil
}

func (s *VaccineService) GetByType(typeStr string, page, pageSize int) ([]*schama.Vaccine, int64, error) {
	tid, has, err := s.TypeRepo.GetIdByType(typeStr)
	if err != nil || has == false {
		log.Println(err.Error())
		return nil, 0, errors.New("type不存在")
	}
	vaccineList, total, err := s.VaccineRepo.GetListByType(tid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取VaccineType失败")
	}
	var list []*schama.Vaccine
	for _, v := range vaccineList {
		if len(v.Type) == 0 {
			err := s.setupVaccineType(v)
			if err != nil {
				return nil, 0, err
			}
		}
		vaccine := &schama.Vaccine{}
		vaccine.GetFormEntity(v)
		list = append(list, vaccine)
	}
	return list, total, nil
}

func (s *VaccineService) GetTypeList(page, pageSize int) ([]*schama.VaccineType, int64, error) {
	t, total, err := s.TypeRepo.GetList(page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取VaccineType失败")
	}
	var list []*schama.VaccineType
	for _, v := range t {
		tp := &schama.VaccineType{}
		tp.GetFormEntity(v)
		list = append(list, tp)
	}
	return list, total, nil
}

func (s *VaccineService) setupVaccineType(v *entity.Vaccine) error {
	t, _, err := s.TypeRepo.GetById(v.Tid)
	if err != nil {
		log.Println(err.Error())
		return errors.New("获取VaccineType失败")
	}
	v.Type = t.Type
	go func() {
		_ = s.VaccineRepo.Update(v)
	}()
	return nil
}
