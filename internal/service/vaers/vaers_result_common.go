package vaers

import (
	"about-vaccine/internal/entity"
	"about-vaccine/internal/schema"
	"log"
	"sync"
)

type VaersResultRepo interface {
	Get(vid int64, sid int64) (*entity.VaersResult, bool, error)
	GetListByVaccineId(vid int64, page, pageSize int) ([]*entity.VaersResult, int64, error)
	GetListBySymptomId(sid int64, page, pageSize int) ([]*entity.VaersResult, int64, error)
	SumByVaccineId(vid int64) (float64, error)
	SumBySymptomId(sid int64) (float64, error)
	Sum() (float64, error)
}

type VaersResultCommon struct {
	vaersResultRepo VaersResultRepo
}

func NewVaersResultCommon(vaersResultRepo VaersResultRepo) *VaersResultCommon {
	return &VaersResultCommon{
		vaersResultRepo: vaersResultRepo,
	}
}

func (vc *VaersResultCommon) GetResultByVaccineId(vid int64, page, pageSize int) ([]*schema.VaersResultInfo, int64, error) {
	rl, total, err := vc.vaersResultRepo.GetListByVaccineId(vid, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	var wg sync.WaitGroup
	var list []*schema.VaersResultInfo
	for _, v := range rl {
		wg.Add(1)
		go func(v *entity.VaersResult) {
			result := vc.completeResult(v)
			list = append(list, result)
			wg.Done()
		}(v)
	}
	wg.Wait()
	return list, total, nil
}

func (vc *VaersResultCommon) GetResultBySymptomId(sid int64, page, pageSize int) ([]*schema.VaersResultInfo, int64, error) {
	vl, total, err := vc.vaersResultRepo.GetListBySymptomId(sid, page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	var wg sync.WaitGroup
	var list []*schema.VaersResultInfo
	for _, v := range vl {
		wg.Add(1)
		go func(v *entity.VaersResult) {
			result := vc.completeResult(v)
			list = append(list, result)
			wg.Done()
		}(v)
	}
	wg.Wait()
	return list, total, nil
}

func (vc *VaersResultCommon) GetResult(vid int64, sid int64) (*schema.VaersResultInfo, bool, error) {
	v, has, err := vc.vaersResultRepo.Get(vid, sid)
	if err != nil {
		return nil, has, err
	}
	result := vc.completeResult(v)
	return result, has, nil
}

func (vc *VaersResultCommon) completeResult(v *entity.VaersResult) *schema.VaersResultInfo {
	// 总共需要四个数据a, b, vc, d
	// 			 	目标疫苗		其他疫苗
	// 目标不良反应	a			b
	// 其他不良反应	vc			d
	result := vc.FormatResultInfo(v)
	total_ac, err := vc.vaersResultRepo.SumByVaccineId(v.VaccineId)
	total_ab, err := vc.vaersResultRepo.SumBySymptomId(v.SymptomId)
	total_abcd, err := vc.vaersResultRepo.Sum()
	log.Println(total_ac, total_ab, total_abcd)
	if err != nil {
		return result
	}
	a := float64(v.Total)
	b, c := total_ab-a, total_ac-a
	d := total_abcd - a - b - c
	if a <= 0 || b <= 0 || c <= 0 || d <= 0 {
		return result
	}
	result.Prr = vc.calculatePrr(a, b, c, d)
	result.Chi = vc.calculateChi(a, b, c, d)
	return result
}

func (vc *VaersResultCommon) calculatePrr(a, b, c, d float64) float64 {
	result := (a / (a + c)) / (b / (b + d))
	return result
}

func (vc *VaersResultCommon) calculateChi(a, b, c, d float64) float64 {
	total := a + b + c + d
	result := (total * (a*d - b*c) * (a*d - b*c)) / ((b + d) * (a + c) * (a + b) * (d + c))
	return result
}
