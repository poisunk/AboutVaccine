package models

import "MyWeb/dao"

type ResponseTerm struct {
	ID         int64  `json:"id"`
	QuestionId int64  `json:"questionId"`
	Answer     string `json:"answer"`
	ResponseID int64  `json:"responseId"`
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

func GetResponseTermByID(id int64) (r *ResponseTerm, err error) {
	r = &ResponseTerm{}
	if err = dao.DB.Where("id = ?", id).First(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func GetResponseTermByQID(qid int64) (r *ResponseTerm, err error) {
	r = &ResponseTerm{}
	if err = dao.DB.Where("question_id = ?", qid).First(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func GetResponseTermListByRID(rid int64) (r []*ResponseTerm, err error) {
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
