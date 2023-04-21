package models

import (
	"about-vaccine/src/dao"
	"github.com/jinzhu/gorm"
	"math"
)

type Question struct {
	Id              int64  `gorm:"int(10);column:id;primary_key" json:"id"`
	QuestionnaireId int64  `gorm:"int(11);column:questionnaire_id" json:"questionnaireId"`
	Type            string `gorm:"varchar(255);column:type" json:"type"`
	Content         string `gorm:"varchar(255);column:content" json:"content"`
	IsRequired      bool   `gorm:"tinyint(1);column:is_required" json:"isRequired"`
	Options         string `gorm:"text;column:options" json:"options"`
	Order           int64  `gorm:"int(11);column:order" json:"order"`
}

func GetQuestionListByQId(id int64) (q []*Question, err error) {
	s := gorm.Expr("IFNULL(order, ?) ASC, content ASC", math.MaxInt32)
	if err = dao.DB.Find(&q, "`questionnaire_id` = ?", id).Order(s).Error; err != nil {
		return nil, err
	}
	return
}

func CreateQuestionList(q []*Question) (err error) {
	t := dao.DB.Begin()
	for i := range q {
		if err = t.Create(q[i]).Error; err != nil {
			t.Rollback()
			return err
		}
	}
	t.Commit()
	return
}

func DeleteQuestionById(id int64) (err error) {
	if err = dao.DB.Delete(Question{}, "id = ?", id).Error; err != nil {
		return err
	}
	return
}

func UpdateQuestion(q *Question) (err error) {
	if err = dao.DB.Model(Question{}).Updates(q).Error; err != nil {
		return err
	}
	return
}
