package service

import (
	"MyWeb/models"
	"errors"
	"log"
)

type VaersVaxServiceImpl struct {
}

func InitVaersVaxService() VaersVaxService {
	return &VaersVaxServiceImpl{}
}

func (vs *VaersVaxServiceImpl) CountVaersVaxByVaersId(vaersId int64) (count int64, err error) {
	count, err = models.CountVaersVaxByVaersId(vaersId)
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("查询VaersCount失败")
	}
	return count, nil
}

func (vs *VaersVaxServiceImpl) GetVaersVaxListByVaersId(vid int64) (list []*VaersVax, err error) {
	vaxList, err := models.GetVaersVaxListByVaersId(vid)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("查询VaersList失败")
	}
	for _, v := range vaxList {
		term, err := vs.GetVaersVaxTerm(v.VaxId)
		if err != nil {
			return nil, err
		}
		list = append(list, &VaersVax{
			ID:           v.ID,
			VaersId:      v.VaersId,
			Type:         term.Type,
			Manufacturer: term.Manufacturer,
			Name:         term.Name,
			Dose:         v.Dose,
			Route:        v.Route,
			Site:         v.Site,
			VaxId:        v.VaxId,
		})
	}
	return list, nil
}

func (vs *VaersVaxServiceImpl) GetVaersVaxTermList(keyword string, page, pageSize int) (list []*VaersVaxTerm, err error) {
	termList, err := models.GetVaersVaxTermListByName(keyword, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("查询VaersTermList失败")
	}
	for _, v := range termList {
		list = append(list, &VaersVaxTerm{
			ID:           v.ID,
			Type:         v.Type,
			Manufacturer: v.Manufacturer,
			Name:         v.Name,
		})
	}
	return list, nil
}

func (vs *VaersVaxServiceImpl) GetVaersIdListByVaxId(vaxId int64) (list []int64, err error) {
	vaxs, err := models.GetVaersIdByVaxId(vaxId)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("查询VaersList失败")
	}
	for _, v := range vaxs {
		list = append(list, v.VaersId)
	}
	return list, nil
}

func (vs *VaersVaxServiceImpl) GetVaersVaxTerm(id int64) (term *VaersVaxTerm, err error) {
	vt, err := models.GetVaersVaxTermById(id)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("查询VaersTerm失败")
	}
	term = &VaersVaxTerm{
		ID:           vt.ID,
		Type:         vt.Type,
		Manufacturer: vt.Manufacturer,
		Name:         vt.Name,
	}
	return term, nil
}
