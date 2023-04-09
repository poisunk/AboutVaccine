package models

import (
	"MyWeb/dao"
	"time"
)

type Questionnaire struct {
	Id          int64     `gorm:"int(11);column:id;primary_key" json:"id"`
	Name        string    `gorm:"varchar(255);column:name" json:"name"`
	Description string    `gorm:"text;column:description" json:"description"`
	OwnerId     int64     `gorm:"int(11);column:owner_id" json:"ownerId"`
	CreateTime  time.Time `gorm:"autoCreateTime;datetime;column:create_time" json:"createTime"`
}

func (Questionnaire) TableName() (name string) {
	return "questionnaire"
}

func GetQuestionnaireList(page int, pageSize int) (q []*Questionnaire, err error) {
	if err = dao.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&q).Error; err != nil {
		return nil, err
	}
	return
}

func GetQuestionnaireListByUid(uid int64) (q []*Questionnaire, err error) {
	if err = dao.DB.Where("owner_id = ?", uid).Find(&q).Error; err != nil {
		return nil, err
	}
	return
}

func GetQuestionnaireById(id int64) (q *Questionnaire, err error) {
	q = &Questionnaire{}
	if err = dao.DB.Where("id = ?", id).First(q).Error; err != nil {
		return nil, err
	}
	return
}

func GetQuestionnaireOwnerIdById(id int64) (uid int64, err error) {
	var q Questionnaire
	if err = dao.DB.Model(Questionnaire{}).Select("owner_id").Where("id = ?", id).First(&q).Error; err != nil {
		return 0, err
	}
	return q.OwnerId, nil
}

func CreateQuestionnaire(q *Questionnaire) (err error) {
	if err = dao.DB.Create(q).Error; err != nil {
		return err
	}
	return nil
}

func CreateAQuestionnaire(q *Questionnaire) (err error) {
	if err = dao.DB.Create(q).Error; err != nil {
		return err
	}
	return nil
}

func DeleteQuestionnaireById(id int64) (err error) {
	if err = dao.DB.Delete(Questionnaire{}, "id = ?", id).Error; err != nil {

	}
	return nil
}

func CountQuestionnaire() (total int64, err error) {
	if err = dao.DB.Model(Questionnaire{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return
}
