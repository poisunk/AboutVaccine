package vaccine

import (
	"about-vaccine/internal/entity"
	"about-vaccine/internal/schema"
)

type VaccineTypeRepo interface {
	Get(id int64) (*entity.VaccineType, bool, error)
	GetList(page, pageSize int) ([]*entity.VaccineType, int64, error)
	GetIdByType(tp string) (int64, bool, error)
	GetTypeById(id int64) (string, bool, error)
}

type VaccineTypeCommon struct {
	vaccineTypeRepo VaccineTypeRepo
}

func NewVaccineTypeCommon(vaccineTypeRepo VaccineTypeRepo) *VaccineTypeCommon {
	return &VaccineTypeCommon{
		vaccineTypeRepo: vaccineTypeRepo,
	}
}

func (vc *VaccineTypeCommon) Get(id int64) (*schema.VaccineTypeDetailInfo, bool, error) {
	e, has, err := vc.vaccineTypeRepo.Get(id)
	if err != nil || !has {
		return nil, has, err
	}
	info := vc.FormatVaccineTypeDetailInfo(e)
	return info, has, nil
}

func (vc *VaccineTypeCommon) GetTypeList(page, pageSize int) ([]*schema.VaccineTypeInfo, int64, error) {
	entitys, total, err := vc.vaccineTypeRepo.GetList(page, pageSize)
	if err != nil {
		return nil, total, err
	}
	list := make([]*schema.VaccineTypeInfo, 0, len(entitys))
	for _, v := range entitys {
		list = append(list, vc.FormatVaccineTypeInfo(v))
	}
	return list, total, nil
}

func (vc *VaccineTypeCommon) GetIdByType(tp string) (int64, bool, error) {
	return vc.vaccineTypeRepo.GetIdByType(tp)
}

func (vc *VaccineTypeCommon) GetTypeById(id int64) (string, bool, error) {
	return vc.vaccineTypeRepo.GetTypeById(id)
}
