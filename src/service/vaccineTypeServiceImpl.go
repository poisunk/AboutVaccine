package service

import (
	"MyWeb/models"
	"errors"
	"log"
)

type VaccineTypeServiceImpl struct {
	VaccineTypeService
}

func InitVaccineTypeService() VaccineTypeService {
	service := &VaccineTypeServiceImpl{}
	return service
}

func (service *VaccineTypeServiceImpl) GetVaccineTypeList(page int, pageSize int) (typeList []string, err error) {
	var l []*models.VaccineType
	if l, err = models.GetVaccineTypeList(page, pageSize); err != nil {
		log.Println(err.Error())
		return nil, errors.New("查询失败！")
	}
	typeList = make([]string, len(l))
	for i, v := range l {
		typeList[i] = v.Type
	}
	return typeList, nil
}

func (service *VaccineTypeServiceImpl) GetVaccineTypeById(id int64) (name string, err error) {
	if name, err = models.GetVaccineTypeNameById(id); err != nil {
		log.Println(err.Error())
		return "", errors.New("查询失败！")
	}
	return name, nil
}

func (service *VaccineTypeServiceImpl) CountVaccineType() (total int, err error) {
	if total, err = models.CountVaccineType(); err != nil {
		log.Println(err.Error())
		return -1, errors.New("查询失败！")
	}
	return total, nil
}
