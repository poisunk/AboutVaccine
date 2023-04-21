package models

import "about-vaccine/src/dao"

type VaccineType struct {
	Id   int64  `gorm:"int(11);column:id;primary_key" json:"id"`
	Type string `gorm:"varchar(255);column:type" json:"type"`
}

func GetVaccineTypeList(page int, pageSize int) (typeList []*VaccineType, err error) {
	if err = dao.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&typeList).Error; err != nil {
		return nil, err
	}
	return
}

func GetVaccineTypeNameById(Id int64) (name string, err error) {
	var v VaccineType
	if err = dao.DB.Model(VaccineType{}).Where("id = ?", Id).Find(&v).Error; err != nil {
		return "", err
	}
	return v.Type, nil
}

func CountVaccineType() (total int, err error) {
	if err = dao.DB.Model(VaccineType{}).Count(&total).Error; err != nil {
		return -1, err
	}
	return
}
