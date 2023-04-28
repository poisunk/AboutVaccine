package vaccine

import (
	"about-vaccine/internal/entity"
	"about-vaccine/internal/schema"
)

type VaccineRepo interface {
	Get(id int64) (*entity.Vaccine, bool, error)
	GetListBySimilarName(keyword string, page, pageSize int) ([]*entity.Vaccine, int64, error)
	GetListByType(tid int64, page, pageSize int) ([]*entity.Vaccine, int64, error)
	GetSimpleListBySimilarName(keyword string, page, pageSize int) ([]*entity.Vaccine, int64, error)
	Update(v *entity.Vaccine) error
}

type VaccineTypeRepo interface {
	Get(id int64) (*entity.VaccineType, bool, error)
	GetList(page, pageSize int) ([]*entity.VaccineType, int64, error)
	GetIdByType(tp string) (int64, bool, error)
}

type VaccineCommon struct {
	vaccineRepo     VaccineRepo
	vaccineTypeRepo VaccineTypeRepo
}

func NewVaccineCommon(vaccineRepo VaccineRepo, vaccineTypeRepo VaccineTypeRepo) *VaccineCommon {
	return &VaccineCommon{
		vaccineRepo:     vaccineRepo,
		vaccineTypeRepo: vaccineTypeRepo,
	}
}

func (vc *VaccineCommon) GetList(keyword string, page, pageSize int) ([]*schema.VaccineInfo, int64, error) {
	entitys, total, err := vc.vaccineRepo.GetListBySimilarName(keyword, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	list := make([]*schema.VaccineInfo, 0, len(entitys))
	for _, v := range entitys {
		if len(v.Type) == 0 {
			err := vc.setupVaccineType(v)
			if err != nil {
				return nil, 0, err
			}
		}
		list = append(list, vc.FormatVaccineInfo(v))
	}
	return list, total, nil
}

func (vc *VaccineCommon) GetListByType(tid int64, page, pageSize int) ([]*schema.VaccineInfo, int64, error) {
	entitys, total, err := vc.vaccineRepo.GetListByType(tid, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	list := make([]*schema.VaccineInfo, 0, len(entitys))
	for _, v := range entitys {
		if len(v.Type) == 0 {
			err := vc.setupVaccineType(v)
			if err != nil {
				return nil, 0, err
			}
		}
		list = append(list, vc.FormatVaccineInfo(v))
	}
	return list, total, nil
}

func (vc *VaccineCommon) GetSimpleListBySimilarName(keyword string, page, pageSize int) ([]*schema.VaccineSimpleInfo, int64, error) {
	entitys, total, err := vc.vaccineRepo.GetListBySimilarName(keyword, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	list := make([]*schema.VaccineSimpleInfo, 0, len(entitys))
	for _, v := range entitys {
		if len(v.Type) == 0 {
			err := vc.setupVaccineType(v)
			if err != nil {
				return nil, 0, err
			}
		}
		list = append(list, vc.FormatVaccineSimpleInfo(v))
	}
	return list, total, nil
}

func (vc *VaccineCommon) GetTypeList(page, pageSize int) ([]*schema.VaccineTypeInfo, int64, error) {
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

func (vc *VaccineCommon) Get(id int64) (*schema.VaccineInfo, bool, error) {
	v, ok, err := vc.vaccineRepo.Get(id)
	if err != nil {
		return nil, false, err
	}
	return vc.FormatVaccineInfo(v), ok, nil
}

func (vc *VaccineCommon) GetTypeIdByType(tp string) (int64, bool, error) {
	return vc.vaccineTypeRepo.GetIdByType(tp)
}

func (vc *VaccineCommon) setupVaccineType(v *entity.Vaccine) error {
	t, _, err := vc.vaccineTypeRepo.Get(v.Tid)
	if err != nil {
		return err
	}
	v.Type = t.Type
	go func() {
		_ = vc.vaccineRepo.Update(v)
	}()
	return nil
}
