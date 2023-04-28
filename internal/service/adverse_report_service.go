package service

import (
	"about-vaccine/internal/middleware/jwt"
	"about-vaccine/internal/schema"
	"about-vaccine/internal/service/adverse_report"
	"errors"
	"log"
	"strconv"
)

type AdverseReportService struct {
	common         *adverse_report.AdverseReportCommon
	vaccineService *VaccineService
}

func NewAdverseReportService(
	common *adverse_report.AdverseReportCommon,
	vaccineService *VaccineService,
) *AdverseReportService {
	return &AdverseReportService{
		common:         common,
		vaccineService: vaccineService,
	}
}

func (a *AdverseReportService) Create(event *schema.AdverseEventAdd, token string) error {
	// 1. 解析token
	var uid *int64
	claim, err := jwt.ParseToken(token)
	if err != nil {
		uid = nil
	} else {
		*uid, _ = strconv.ParseInt(claim.Id, 10, 64)
	}
	// 2. 创建Event
	err = a.common.Create(event, uid)
	if err != nil {
		log.Println(err.Error())
		return errors.New("创建Event失败")
	}
	return nil
}

func (a *AdverseReportService) Get(id int64) (*schema.AdverseEventInfo, error) {
	e, _, err := a.common.Get(id)
	if err != nil {
		return nil, errors.New("不良反应报告不存在")
	}
	return e, nil
}

func (a *AdverseReportService) GetList(page, pageSize int) ([]*schema.AdverseEventInfo, int64, error) {
	el, total, err := a.common.GetList(page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取Event列表失败")
	}
	return el, total, nil
}

func (a *AdverseReportService) GetListByUid(uid int64, page, pageSize int) ([]*schema.AdverseEventInfo, int64, error) {
	el, total, err := a.common.GetListByUid(uid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取Event列表失败")
	}
	return el, total, nil
}

func (a *AdverseReportService) Delete(id int64) error {
	err := a.common.Delete(id)
	if err != nil {
		log.Println(err.Error())
		return errors.New("删除Event失败")
	}
	return nil
}
