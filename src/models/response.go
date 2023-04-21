package models

import (
	"about-vaccine/dao"
	"time"
)

type Response struct {
	Id              int64     `gorm:"int(11);column:id;primary_key" json:"id"`
	QuestionnaireId int64     `gorm:"int(11);column:questionnaire_id" json:"questionnaireId"`
	UserId          int64     `gorm:"int(11);column:user_id" json:"userId"`
	ResponseTime    time.Time `gorm:"autoCreateTime;datetime;column:response_time" json:"responseTime"`
}

func CreateResponse(r *Response) (err error) {
	if err = dao.DB.Create(r).Error; err != nil {
		return err
	}
	return
}

func GetResponseListByQid(qid int64, page int, pageSize int) (r []*Response, total int, err error) {
	db := dao.DB.Model(Response{}).Where("questionnaire_id = ?", qid).Count(&total)
	if err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&r).Error; err != nil {
		return nil, 0, err
	}
	return
}

func GetResponseListByUid(uid int64, page int, pageSize int) (r []*Response, total int, err error) {
	db := dao.DB.Model(Response{}).Where("user_id = ?", uid).Count(&total)
	if err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&r).Error; err != nil {
		return nil, 0, err
	}
	return
}

func DeleteResponse(id int64) (err error) {
	if err = dao.DB.Where("id = ?", id).Delete(&Response{}).Error; err != nil {
		return err
	}
	return
}

func GetResponseOwnerId(id int64) (uid int64, err error) {
	var response Response
	if err = dao.DB.Where("id = ?", id).First(&response).Error; err != nil {
		return 0, err
	}
	return response.UserId, nil
}
