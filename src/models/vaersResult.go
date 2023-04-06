package models

import "MyWeb/dao"

type VaersResult struct {
	ID        int64  `json:"id"`
	VaccineId int64  `json:"vaccineId"`
	Name      string `json:"name"`
	SymptomId int64  `json:"symptomId"`
	Symptom   string `json:"symptom"`
	Total     int64  `json:"total"`
}

func (v *VaersResult) TableName() string {
	return "vaers_result"
}

func GetVaersResultListByVaccineId(vid int64, page, pageSize int) (list []*VaersResult, err error) {
	list = make([]*VaersResult, 0)
	err = dao.DB.Where("vaccine_id = ?", vid).Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return
}

func GetVaersResultListBySymptomId(sid int64, page, pageSize int) (list []*VaersResult, err error) {
	list = make([]*VaersResult, 0)
	err = dao.DB.Where("symptom_id = ?", sid).Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	return
}

func GetVaersResult(vid, sid int64) (vaers *VaersResult, err error) {
	vaers = &VaersResult{}
	err = dao.DB.Where("vaccine_id = ? AND symptom_id = ?", vid, sid).First(vaers).Error
	return
}

func CountVaersResult() (int64, error) {
	var count int64
	err := dao.DB.Model(&VaersResult{}).Count(&count).Error
	return count, err
}

func CountVaersResultByVaccineId(vid int64) (int64, error) {
	var count int64
	err := dao.DB.Model(&VaersResult{}).Where("vaccine_id = ?", vid).Count(&count).Error
	return count, err
}

func CountVaersResultBySymptomId(sid int64) (int64, error) {
	var count int64
	err := dao.DB.Model(&VaersResult{}).Where("symptom_id = ?", sid).Count(&count).Error
	return count, err
}
