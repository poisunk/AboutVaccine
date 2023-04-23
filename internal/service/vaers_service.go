package service

import (
	"about-vaccine/internal/entity"
	"about-vaccine/internal/repo"
	"about-vaccine/internal/schama"
	"errors"
	"log"
	"sync"
)

type VaersService struct {
	vaers            *repo.VaersRepo
	vaersResult      *repo.VaersResultRepo
	vaersSymptom     *repo.VaersSymptomRepo
	vaersSymptomTerm *repo.VaersSymptomTermRepo
	vaersVax         *repo.VaersVaxRepo
	vaersVaxTerm     *repo.VaersVaxTermRepo
}

func NewVaersService(
	vaers *repo.VaersRepo,
	vaersResult *repo.VaersResultRepo,
	vaersSymptom *repo.VaersSymptomRepo,
	vaersSymptomTerm *repo.VaersSymptomTermRepo,
	vaersVax *repo.VaersVaxRepo,
	vaersVaxTerm *repo.VaersVaxTermRepo,
) *VaersService {
	return &VaersService{
		vaers:            vaers,
		vaersResult:      vaersResult,
		vaersSymptom:     vaersSymptom,
		vaersSymptomTerm: vaersSymptomTerm,
		vaersVax:         vaersVax,
		vaersVaxTerm:     vaersVaxTerm,
	}
}

func (s *VaersService) GetResultByVaccineId(vid int64, page, pageSize int) ([]*schama.VaersResult, int64, error) {
	rl, total, err := s.vaersResult.GetListByVaccineId(vid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("查询失败")
	}
	var wg sync.WaitGroup
	var list []*schama.VaersResult
	for _, v := range rl {
		wg.Add(1)
		go func(v *entity.VaersResult) {
			result := s.completeResult(v)
			list = append(list, result)
			wg.Done()
		}(v)
	}
	wg.Wait()
	return list, total, nil
}

func (s *VaersService) GetResultBySymptomId(sid int64, page, pageSize int) ([]*schama.VaersResult, int64, error) {
	vl, total, err := s.vaersResult.GetListBySymptomId(sid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("查询失败")
	}
	var wg sync.WaitGroup
	var list []*schama.VaersResult
	for _, v := range vl {
		wg.Add(1)
		go func(v *entity.VaersResult) {
			result := s.completeResult(v)
			list = append(list, result)
			wg.Done()
		}(v)
	}
	wg.Wait()
	return list, total, nil
}

func (s *VaersService) GetResult(vid int64, sid int64) (*schama.VaersResult, error) {
	v, has, err := s.vaersResult.Get(vid, sid)
	if err != nil || !has {
		log.Println(err.Error())
		return nil, errors.New("记录不存在")
	}
	result := s.completeResult(v)
	return result, nil
}

func (s *VaersService) GetVaccineTermList(keyword string, page, pageSize int) ([]*schama.VaersVaxTerm, int64, error) {
	tl, total, err := s.vaersVaxTerm.GetByName(keyword, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("查询失败")
	}
	var list []*schama.VaersVaxTerm
	for _, v := range tl {
		vaccine := &schama.VaersVaxTerm{}
		vaccine.GetFormEntity(v)
		list = append(list, vaccine)
	}
	return list, total, nil
}

func (s *VaersService) GetSymptomTermList(keyword string, page, pageSize int) ([]*schama.VaersSymptomTerm, int64, error) {
	sl, total, err := s.vaersSymptomTerm.GetByName(keyword, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("查询失败")
	}
	var list []*schama.VaersSymptomTerm
	for _, v := range sl {
		symptom := &schama.VaersSymptomTerm{}
		symptom.GetFormEntity(v)
		list = append(list, symptom)
	}
	return list, total, nil
}

func (s *VaersService) completeResult(v *entity.VaersResult) *schama.VaersResult {
	// 总共需要四个数据a, b, c, d
	// 			 	目标疫苗		其他疫苗
	// 目标不良反应	a			b
	// 其他不良反应	c			d
	result := &schama.VaersResult{}
	result.GetFormEntity(v)
	total_ac, err := s.vaersResult.SumByVaccineId(v.VaccineId)
	total_ab, err := s.vaersResult.SumBySymptomId(v.SymptomId)
	total_abcd, err := s.vaersResult.Sum()
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
	result.Prr = s.calculatePrr(a, b, c, d)
	result.Chi = s.calculateChi(a, b, c, d)
	return result
}

func (s *VaersService) calculatePrr(a, b, c, d float64) float64 {
	result := (a / (a + c)) / (b / (b + d))
	return result
}

func (s *VaersService) calculateChi(a, b, c, d float64) float64 {
	total := a + b + c + d
	result := (total * (a*d - b*c) * (a*d - b*c)) / ((b + d) * (a + c) * (a + b) * (d + c))
	return result
}
