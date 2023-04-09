package service

import (
	"MyWeb/models"
	"MyWeb/utile"
	"errors"
	"log"
)

type VaccineServiceImpl struct {
	VaccineTypeService
}

func InitVaccineService() VaccineService {
	service := &VaccineServiceImpl{
		VaccineTypeService: InitVaccineTypeService(),
	}
	return service
}

func (service *VaccineServiceImpl) CreateVaccine(vaccine Vaccine) error {
	v := &models.Vaccine{}
	err := utile.StructConv(vaccine, v)
	if err != nil {
		log.Println(err.Error())
		return errors.New("数据格式有误！")
	}
	if err = models.CreateVaccine(v); err != nil {
		log.Println(err.Error())
		return errors.New("数据库错误！")
	}
	return nil
}

func (service *VaccineServiceImpl) GetVaccine(id int64) (v *Vaccine, err error) {
	vaccine := &models.Vaccine{}
	if vaccine, err = models.GetVaccineById(id); err != nil {
		log.Println(err.Error())
		return nil, errors.New("数据库错误！")
	}
	v = &Vaccine{}
	err = utile.StructConv(vaccine, v)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("数据格式有误！")
	}
	// 查询疫苗类型
	v.Type, err = service.GetVaccineTypeById(vaccine.Tid)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (service *VaccineServiceImpl) GetVaccineList(page int, pageSize int, keyword string) (results []*Vaccine, total int, err error) {
	var vList []*models.Vaccine
	if vList, total, err = models.GetVaccineList(page, pageSize, keyword); err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("数据库错误！")
	}
	for _, v := range vList {
		vaccine := &Vaccine{}
		err = utile.StructConv(v, &vaccine)
		if err != nil {
			log.Println(err.Error())
			return nil, 0, errors.New("数据格式有误！")
		}
		// 查询疫苗类型
		vaccine.Type, err = service.GetVaccineTypeById(v.Tid)
		if err != nil {
			log.Println(err.Error())
			return nil, 0, errors.New("数据库错误！")
		}
		results = append(results, vaccine)
	}
	return results, total, nil
}

func (service *VaccineServiceImpl) DeleteVaccine(id int64) error {
	if err := models.DeleteVaccine(id); err != nil {
		log.Println(err.Error())
		return errors.New("数据库错误！")
	}
	return nil
}

func (service *VaccineServiceImpl) GetVaccineExampleList(page int, pageSize int, limit int) (results []*VaccineExample, total int, err error) {
	tList, err := models.GetVaccineTypeList(page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("数据库错误！")
	}
	for _, t := range tList {
		vList, err := models.GetVaccineListByTid(t.Id, limit+1)
		if err != nil {
			log.Println(err.Error())
			return nil, 0, errors.New("数据库错误！")
		}
		example := new(VaccineExample)
		for i, v := range vList {
			// 如果是达到limit
			if i == limit {
				example.More = true
				break
			}
			vaccine := &Vaccine{}
			err = utile.StructConv(v, &vaccine)
			if err != nil {
				log.Println(err.Error())
				return nil, 0, errors.New("数据格式有误！")
			}
			// 查询疫苗类型
			vaccine.Type = t.Type
			example.VaccineList = append(example.VaccineList, vaccine)
		}
		example.Type = t.Type
		results = append(results, example)
	}
	total, err = models.CountVaccineType()
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("数据库错误！")
	}
	return results, total, nil
}

func (service *VaccineServiceImpl) GetVaccineExampleByTid(tid int64, limit int) (example *VaccineExample, err error) {
	t, err := service.GetVaccineTypeById(tid)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("数据库错误！")
	}
	// 查询疫苗
	vList, err := models.GetVaccineListByTid(tid, limit+1)
	example = new(VaccineExample)
	for i, v := range vList {
		// 如果是达到limit
		if i == limit {
			example.More = true
			break
		}
		vaccine := &Vaccine{}
		err = utile.StructConv(v, &vaccine)
		if err != nil {
			log.Println(err.Error())
			return nil, errors.New("数据格式有误！")
		}
		// 查询疫苗类型
		vaccine.Type = t
		example.VaccineList = append(example.VaccineList, vaccine)
	}
	example.Type = t
	return example, nil
}
