package models

import (
	"about-vaccine/src/dao"
)

type OAETerm struct {
	Id              int    `gorm:"int(11);column:id;primary_key" json:"id"`
	TermIRI         string `gorm:"varchar(255);column:TermIRI" json:"termIRI"`
	TermLabel       string `gorm:"varchar(255);column:TermLabel" json:"termLabel"`
	ParentTermIRI   string `gorm:"varchar(255);column:ParentTermIRI" json:"parentTermIRI"`
	ParentTermLabel string `gorm:"varchar(255);column:ParentTermLabel" json:"parentTermLabel"`
	AlternativeTerm string `gorm:"varchar(255);column:AlternativeTerm" json:"alternativeTerm"`
	Definition      string `gorm:"varchar(255);column:Definition" json:"definition"`
}

func GetOaeTermByLabel(label string, page int, pageSize int) (oaeList []*OAETerm, total int, err error) {
	db := dao.DB.Model(OAETerm{}).Where("TermLabel LIKE ?", "%"+label+"%").Count(&total)
	if err = db.Offset((page - 1) * pageSize).Limit(50).Find(&oaeList).Error; err != nil {
		return nil, 0, err
	}
	return
}

func GetOaeTermByIRI(IRI string) (oaeTerm *OAETerm, err error) {
	oaeTerm = new(OAETerm)
	if err = dao.DB.First(oaeTerm, "TermIRI = ?", IRI).Error; err != nil {
		return nil, err
	}
	return
}
