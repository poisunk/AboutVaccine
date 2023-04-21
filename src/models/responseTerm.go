package models

import "about-vaccine/dao"

type ResponseTerm struct {
	Id         int64  `json:"id"`
	QuestionId int64  `json:"questionId"`
	Answer     string `json:"answer"`
	ResponseId int64  `json:"responseId"`
}

func (r *ResponseTerm) TableName() string {
	return "response_term"
}

func CreateResponseTerm(r *ResponseTerm) (err error) {
	if err = dao.DB.Create(r).Error; err != nil {
		return err
	}
	return nil
}

func GetResponseTermById(id int64) (r *ResponseTerm, err error) {
	r = &ResponseTerm{}
	if err = dao.DB.Where("id = ?", id).First(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func GetResponseTermByQId(qid int64) (r *ResponseTerm, err error) {
	r = &ResponseTerm{}
	if err = dao.DB.Where("question_id = ?", qid).First(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func GetResponseTermListByRId(rid int64) (r []*ResponseTerm, err error) {
	db := dao.DB.Model(ResponseTerm{}).Where("response_id = ?", rid)
	if err = db.Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func DeleteResponseTerm(id int64) (err error) {
	if err = dao.DB.Where("id = ?", id).Delete(&ResponseTerm{}).Error; err != nil {
		return err
	}
	return nil
}
