package models

import (
	"MyWeb/dao"
	"MyWeb/utile"
)

type VaersVaxTerm struct {
	ID           int64  `json:"id"`
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
}

func (v *VaersVaxTerm) TableName() string {
	return "vaers_vax_terms"
}

func GetVaersVaxTermById(id int64) (v *VaersVaxTerm, err error) {
	if err = dao.DB.Where("id = ?", id).First(&v).Error; err != nil {
		return nil, err
	}
	return v, nil
}

func GetVaersVaxTermListByName(name string, page, pageSize int) (v []*VaersVaxTerm, err error) {
	if err = dao.DB.Where("name like ?", utile.HandleSearchWord(name)).Limit(pageSize).Offset((page - 1) * pageSize).Find(&v).Error; err != nil {
		return nil, err
	}
	return v, nil
}
