package models

import "MyWeb/dao"

type VaersVax struct {
	ID      int64  `json:"id"`
	VaersId int64  `json:"vaersId"`
	Dose    string `json:"dose"`
	Route   string `json:"route"`
	Site    string `json:"site"`
	VaxId   int64  `json:"vaxId"`
}

func (v *VaersVax) TableName() string {
	return "vaers_vax"
}

func GetVaersVaxListByVaersId(vaersId int64) (list []*VaersVax, err error) {
	if err = dao.DB.Where("vaers_id = ?", vaersId).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func GetVaersVaxById(id int64) (v *VaersVax, err error) {
	if err = dao.DB.Where("id = ?", id).First(&v).Error; err != nil {
		return nil, err
	}
	return v, nil
}

func GetVaersVaxListByVaxId(vaxId int64, page, pageSize int) (list []*VaersVax, err error) {
	if err = dao.DB.Offset((page-1)*pageSize).Limit(pageSize).Where("vax_id = ?", vaxId).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
