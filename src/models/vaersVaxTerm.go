package models

import (
	"MyWeb/dao"
	"MyWeb/utile"
)

type VaersVaxTerm struct {
	Id           int64  `json:"id"`
	Type         string `json:"type"`
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
}

func (v *VaersVaxTerm) TableName() string {
	return "vaers_vax_term"
}

func GetVaersVaxTermById(id int64) (v *VaersVaxTerm, err error) {
	v = &VaersVaxTerm{}
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

func CountVaersTermByName(name string) (count int64, err error) {
	if err = dao.DB.Model(&VaersVaxTerm{}).Where("name like ?", utile.HandleSearchWord(name)).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func CountVaersTerm() (count int64, err error) {
	if err = dao.DB.Model(&VaersVaxTerm{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
