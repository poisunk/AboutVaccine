package vaers

import (
	"about-vaccine/internal/entity"
	"about-vaccine/internal/schema"
)

type VaersVaxTermRepo interface {
	GetListBySimilarName(keyword string, page, pageSize int) ([]*entity.VaersVaxTerm, int64, error)
}

type VaersVaxCommon struct {
	vaersVaxTermRepo VaersVaxTermRepo
}

func NewVaersVaxCommon(vaersVaxTermRepo VaersVaxTermRepo) *VaersVaxCommon {
	return &VaersVaxCommon{
		vaersVaxTermRepo: vaersVaxTermRepo,
	}
}

func (vc *VaersVaxCommon) GetVaxTermList(keyword string, page, pageSize int) ([]*schema.VaersVaxTerm, int64, error) {
	tl, total, err := vc.vaersVaxTermRepo.GetListBySimilarName(keyword, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	var list []*schema.VaersVaxTerm
	for _, v := range tl {
		vaccine := vc.FormatVaxTerm(v)
		list = append(list, vaccine)
	}
	return list, total, nil
}
