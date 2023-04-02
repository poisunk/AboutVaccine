package service

type VaccineTypeService interface {
	GetVaccineTypeList(page int, pageSize int) (typeList []string, err error)
	GetVaccineTypeById(id int64) (name string, err error)
	CountVaccineType() (total int, err error)
}
