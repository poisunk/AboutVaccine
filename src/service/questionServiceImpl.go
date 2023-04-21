package service

import (
	"about-vaccine/src/models"
	"errors"
	"log"
	"strings"
)

type QuestionServiceImpl struct {
	QuestionnaireService
}

func InitQuestionService() QuestionService {
	return &QuestionServiceImpl{
		QuestionnaireService: InitQuestionnaireService(),
	}
}

const SplitSign = "$"

func (service *QuestionServiceImpl) GetQuestionTypeList() []string {
	// 拿到QuestionType中的所有key
	keys := make([]string, 0, len(QuestionType))
	for k := range QuestionType {
		keys = append(keys, k)
	}
	return keys
}

func (service *QuestionServiceImpl) CreateQuestion(q []*Question) error {
	var qs []*models.Question
	for _, qi := range q {
		if len(qi.Content) == 0 {
			return errors.New("问题内容不能为空！")
		}
		qs = append(qs, &models.Question{
			QuestionnaireId: qi.QuestionnaireId,
			Type:            qi.Type,
			Content:         qi.Content,
			IsRequired:      qi.IsRequired,
			Options:         strings.Join(qi.Options, SplitSign),
			Order:           qi.Order,
		})
	}
	if err := models.CreateQuestionList(qs); err != nil {
		log.Println(err.Error())
		return errors.New("创建问题失败！")
	}
	return nil
}

func (service *QuestionServiceImpl) DeleteQuestionById(id int64) error {
	if err := models.DeleteQuestionById(id); err != nil {
		log.Println(err.Error())
		return errors.New("删除问题失败！")
	}
	return nil
}

func (service *QuestionServiceImpl) GetQuestionList(qid int64) (q []*Question, err error) {
	var qs []*models.Question
	if qs, err = models.GetQuestionListByQId(qid); err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取问题失败！")
	}
	for _, qi := range qs {
		q = append(q, &Question{
			Id:              qi.Id,
			QuestionnaireId: qi.QuestionnaireId,
			Type:            qi.Type,
			Content:         qi.Content,
			IsRequired:      qi.IsRequired,
			Options:         strings.Split(qi.Options, SplitSign),
			Order:           qi.Order,
		})
	}
	return q, nil
}
