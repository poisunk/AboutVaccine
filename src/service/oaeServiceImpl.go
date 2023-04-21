package service

import (
	"about-vaccine/src/models"
	"about-vaccine/src/utile"
	"errors"
	"log"
)

type OAEServiceImpl struct {
}

func InitOAEService() OAEService {
	return &OAEServiceImpl{}
}

func (service *OAEServiceImpl) GetOAETermByLabel(label string, page int, pageSize int) (oaeList []*OAETerm, total int, err error) {
	terms, total, err := models.GetOaeTermByLabel(label, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取失败")
	}
	for _, term := range terms {
		t := &OAETerm{}
		err := utile.StructConv(term, t)
		if err != nil {
			log.Println(err.Error())
			return nil, 0, errors.New("转换错误")
		}
		oaeList = append(oaeList, t)
	}
	return oaeList, total, nil
}

func (service *OAEServiceImpl) GetOAETermByIRI(IRI string) (oaeTerm *OAETerm, err error) {
	term := new(models.OAETerm)
	if term, err = models.GetOaeTermByIRI(IRI); err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取失败")
	}
	oaeTerm = &OAETerm{}
	err = utile.StructConv(term, oaeTerm)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("转换错误")
	}
	return oaeTerm, nil
}

func (service *OAEServiceImpl) GetOAETermParentList(IRI string) (oaeList []*OAETerm, err error) {
	for {
		if len(IRI) == 0 {
			break
		}
		term := new(models.OAETerm)
		if term, err = models.GetOaeTermByIRI(IRI); err != nil {
			log.Println(err.Error())
			return nil, errors.New("获取失败")
		}
		t := new(OAETerm)
		err = utile.StructConv(term, t)
		if err != nil {
			log.Println(err.Error())
			return nil, errors.New("转换错误")
		}
		oaeList = append(oaeList, t)
		IRI = term.ParentTermIRI
	}
	// 翻转列表
	oaeList = ReverseList(oaeList)
	return oaeList, nil
}

func ReverseList(oaeList []*OAETerm) []*OAETerm {
	for i, j := 0, len(oaeList)-1; i < j; i, j = i+1, j-1 {
		oaeList[i], oaeList[j] = oaeList[j], oaeList[i]
	}
	return oaeList
}
