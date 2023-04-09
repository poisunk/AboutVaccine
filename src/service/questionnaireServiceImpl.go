package service

import (
	"MyWeb/models"
	"MyWeb/utile"
	"errors"
	"log"
)

type QuestionnaireServiceImpl struct {
	UserService
}

func InitQuestionnaireService() QuestionnaireService {
	return &QuestionnaireServiceImpl{
		UserService: InitUserService(),
	}
}

func (s *QuestionnaireServiceImpl) GetQuestionnaireList(page, pageSize int) ([]*Questionnaire, int64, error) {
	var result []*Questionnaire
	list, err := models.GetQuestionnaireList(page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取问卷列表失败")
	}
	for _, v := range list {
		t := &Questionnaire{}
		_ = utile.StructConv(v, t)
		u, err := s.GetUser(t.OwnerId)
		if err != nil {
			return nil, 0, err
		}
		t.OwnerName = u.Nickname
		result = append(result, t)
	}
	total, err := models.CountQuestionnaire()
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取问卷列表失败")
	}
	return result, total, nil
}

func (s *QuestionnaireServiceImpl) GetQuestionnaireByUid(uid int64) ([]*Questionnaire, error) {
	q, err := models.GetQuestionnaireListByUid(uid)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取问卷失败")
	}
	var result []*Questionnaire
	for _, v := range q {
		t := &Questionnaire{}
		_ = utile.StructConv(v, t)
		u, err := s.GetUser(t.OwnerId)
		if err != nil {
			return nil, err
		}
		t.OwnerName = u.Nickname
		result = append(result, t)
	}
	return result, nil
}

func (s *QuestionnaireServiceImpl) CreateQuestionnaire(questionnaire *Questionnaire) error {
	q := new(models.Questionnaire)
	_ = utile.StructConv(questionnaire, q)
	err := models.CreateQuestionnaire(q)
	if err != nil {
		log.Println(err.Error())
		return errors.New("创建问卷失败")
	}
	questionnaire.Id = q.Id
	return nil
}

func (s *QuestionnaireServiceImpl) DeleteQuestionnaireById(id int64) error {
	err := models.DeleteQuestionnaireById(id)
	if err != nil {
		log.Println(err.Error())
		return errors.New("删除问卷失败")
	}
	return nil
}

func (s *QuestionnaireServiceImpl) GetQuestionnaireById(id int64) (*Questionnaire, error) {
	q, err := models.GetQuestionnaireById(id)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取问卷失败")
	}
	t := &Questionnaire{}
	_ = utile.StructConv(q, t)
	return t, nil
}

func (s *QuestionnaireServiceImpl) GetQuestionnaireOwnerIdById(id int64) (uid int64, err error) {
	uid, err = models.GetQuestionnaireOwnerIdById(id)
	if err != nil {
		log.Println(err.Error())
		return 0, errors.New("获取问卷失败")
	}
	return uid, nil
}
