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

type VaccineCommon struct {
	vaccineRepo       VaccineRepo
	vaccineTypeCommon *VaccineTypeCommon
}

func NewVaccineCommon(vaccineRepo VaccineRepo, vaccineTypeCommon *VaccineTypeCommon) *VaccineCommon {
	return &VaccineCommon{
		vaccineRepo:       vaccineRepo,
		vaccineTypeCommon: vaccineTypeCommon,
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

func (vc *VaccineCommon) Get(id int64) (*schema.VaccineInfo, bool, error) {
	v, ok, err := vc.vaccineRepo.Get(id)
	if err != nil {
		return nil, false, err
	}
	return vc.FormatVaccineInfo(v), ok, nil
}

func (vc *VaccineCommon) setupVaccineType(v *entity.Vaccine) error {
	t, has, err := vc.vaccineTypeCommon.GetTypeById(v.Tid)
	if err != nil || !has {
		return err
	}
	v.Type = t
	go func() {
		_ = vc.vaccineRepo.Update(v)
	}()
	return nil
}
