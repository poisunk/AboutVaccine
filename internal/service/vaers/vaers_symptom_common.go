package vaers

import (
	"vax/internal/entity"
	"vax/internal/schema"
)

type VaersSymptomTermRepo interface {
	GetListBySimilarName(keyword string, page, pageSize int) ([]*entity.VaersSymptomTerm, int64, error)
}

type VaersSymptomCommon struct {
	vaersSymptomTermRepo VaersSymptomTermRepo
}

func NewVaersSymptomCommon(vaersSymptomTermRepo VaersSymptomTermRepo) *VaersSymptomCommon {
	return &VaersSymptomCommon{
		vaersSymptomTermRepo: vaersSymptomTermRepo,
	}
}

func (vc *VaersSymptomCommon) GetSymptomTermList(keyword string, page, pageSize int) ([]*schema.VaersSymptomTerm, int64, error) {
	tl, total, err := vc.vaersSymptomTermRepo.GetListBySimilarName(keyword, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	var list []*schema.VaersSymptomTerm
	for _, v := range tl {
		symptom := vc.FormatSymptomTerm(v)
		list = append(list, symptom)
	}
	return list, total, nil
}
