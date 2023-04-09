package service

import (
	"MyWeb/models"
	"errors"
	"log"
	"time"
)

type ResponseServiceImpl struct {
	QuestionnaireService
	QuestionService
	UserService
}

func InitResponseService() ResponseService {
	return &ResponseServiceImpl{
		QuestionnaireService: InitQuestionnaireService(),
		QuestionService:      InitQuestionService(),
		UserService:          InitUserService(),
	}
}

func (rs *ResponseServiceImpl) GetResponsesByQid(qid int64, page, pageSize int) (
	result []*Response, total int, err error) {
	response, total, err := models.GetResponseListByQid(qid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取问卷答案列表失败")
	}

	for _, r := range response {
		u, err := rs.GetUser(r.UserId)
		if err != nil {
			return nil, 0, err
		}
		temp := &Response{
			Id:              r.Id,
			QuestionnaireId: r.QuestionnaireId,
			UserId:          r.UserId,
			UserName:        u.Nickname,
			ResponseTime:    r.ResponseTime,
		}
		terms, err := rs.GetResponseTermsByRid(r.Id)
		if err != nil {
			return nil, 0, err
		}
		temp.ResponseTermList = terms
		result = append(result, temp)
	}
	return result, total, nil
}

func (rs *ResponseServiceImpl) GetResponseByUid(uid int64, page, pageSize int) (
	result []*Response, total int, err error) {
	response, total, err := models.GetResponseListByUid(uid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取问卷答案列表失败")
	}
	for _, r := range response {
		u, err := rs.GetUser(r.UserId)
		if err != nil {
			return nil, 0, err
		}
		temp := &Response{
			Id:              r.Id,
			QuestionnaireId: r.QuestionnaireId,
			UserId:          r.UserId,
			UserName:        u.Nickname,
			ResponseTime:    r.ResponseTime,
		}
		terms, err := rs.GetResponseTermsByRid(r.Id)
		if err != nil {
			return nil, 0, err
		}
		temp.ResponseTermList = terms
		result = append(result, temp)
	}
	return result, total, nil
}

func (rs *ResponseServiceImpl) DeleteResponse(id int64) (err error) {
	err = models.DeleteResponse(id)
	if err != nil {
		log.Println(err.Error())
		return errors.New("删除问卷答案失败")
	}
	return nil
}

func (rs *ResponseServiceImpl) CreateResponse(r *Response) (err error) {
	response := &models.Response{
		QuestionnaireId: r.QuestionnaireId,
		UserId:          r.UserId,
		ResponseTime:    time.Now(),
	}

	// 检查数据
	// 1. 回答questionId不能为空
	for _, v := range r.ResponseTermList {
		if v.QuestionId == 0 {
			return errors.New("问题Id不能为空")
		}
	}
	err = models.CreateResponse(response)
	if err != nil {
		log.Println(err.Error())
		return errors.New("创建问卷答案失败")
	}
	for _, v := range r.ResponseTermList {
		term := &models.ResponseTerm{
			QuestionId: v.QuestionId,
			Answer:     v.Answer,
			ResponseId: response.Id,
		}
		err = models.CreateResponseTerm(term)
		if err != nil {
			log.Println(err.Error())
			return errors.New("创建问卷答案失败")
		}
	}
	return nil
}

func (rs *ResponseServiceImpl) GetResponseTermsByRid(rid int64) (r []*ResponseTerm, err error) {
	response, err := models.GetResponseTermListByRId(rid)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取问卷答案列表失败")
	}
	for _, v := range response {
		r = append(r, &ResponseTerm{
			Id:         v.Id,
			QuestionId: v.QuestionId,
			Answer:     v.Answer,
		})
	}
	return r, nil
}

func (rs *ResponseServiceImpl) GetResponseOwnerId(rid int64) (uid int64, err error) {
	uid, err = models.GetResponseOwnerId(rid)
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("获取问卷答案列表失败")
	}
	return uid, nil
}
