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
	resultCommon   *adverse_report.AdverseResultCommon
	vaccineService *VaccineService
}

func NewAdverseReportService(
	common *adverse_report.AdverseReportCommon,
	resultCommon *adverse_report.AdverseResultCommon,
	vaccineService *VaccineService,
) *AdverseReportService {
	return &AdverseReportService{
		common:         common,
		resultCommon:   resultCommon,
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
	e, has, err := a.common.Get(id)
	if err != nil || !has {
		return nil, errors.New("不良反应报告不存在")
	}
	return e, nil
}

func (a *AdverseReportService) GetList(page, pageSize int) ([]*schema.AdverseEventBriefInfo, int64, error) {
	el, total, err := a.common.GetList(page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取Event列表失败")
	}
	return el, total, nil
}

func (a *AdverseReportService) GetListByKeyword(keyword string, page, pageSize int) ([]*schema.AdverseEventBriefInfo, int64, error) {
	el, total, err := a.common.GetListByKeyword(keyword, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取Event列表失败")
	}
	return el, total, nil
}

func (a *AdverseReportService) GetListByUid(uid int64, page, pageSize int) ([]*schema.AdverseEventBriefInfo, int64, error) {
	el, total, err := a.common.GetListByUid(uid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, 0, errors.New("获取Event列表失败")
	}
	return el, total, nil
}

func (a *AdverseReportService) GetListByVaccineId(vidStr string, page, pageSize int) ([]*schema.AdverseEventBriefInfo, error) {
	vid, err := strconv.ParseInt(vidStr, 10, 64)
	if err != nil {
		return nil, errors.New("vaccineId参数错误")
	}
	el, err := a.common.GetListByVaccineId(vid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取Event列表失败")
	}
	return el, nil
}

func (a *AdverseReportService) GetListByOAEId(oidStr string, page, pageSize int) ([]*schema.AdverseEventBriefInfo, error) {
	oid, err := strconv.ParseInt(oidStr, 10, 64)
	if err != nil {
		return nil, errors.New("oaeId参数错误")
	}
	el, err := a.common.GetListByOAEId(oid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取Event列表失败")
	}
	return el, nil
}

func (a *AdverseReportService) GetResult(vid, oid string, page, pageSize int) ([]*schema.AdverseResultInfo, error) {
	if len(vid) != 0 && len(oid) == 0 {
		id, _ := strconv.ParseInt(vid, 10, 64)
		return a.GetResultByVaccineId(id, page, pageSize)
	}
	if len(vid) == 0 && len(oid) != 0 {
		id, _ := strconv.ParseInt(oid, 10, 64)
		return a.GetResultByOAEId(id, page, pageSize)
	}
	return nil, errors.New("参数不能为空")
}

func (a *AdverseReportService) GetResultByVaccineId(vid int64, page, pageSize int) ([]*schema.AdverseResultInfo, error) {
	el, err := a.resultCommon.GetByVaccineId(vid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取Result列表失败")
	}
	return el, nil
}

func (a *AdverseReportService) GetResultByOAEId(oid int64, page, pageSize int) ([]*schema.AdverseResultInfo, error) {
	el, err := a.resultCommon.GetByOAEId(oid, page, pageSize)
	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("获取Result列表失败")
	}
	return el, nil
}

func (a *AdverseReportService) Delete(id int64, token string) error {
	claim, err := jwt.ParseToken(token)
	if err != nil {
		return errors.New("无效token")
	}
	owner, has, err := a.common.GetUid(id)
	if err != nil {
		return errors.New("不良反应报告不存在")
	}
	uid, _ := strconv.ParseInt(claim.Id, 10, 64)
	if uid != owner || !has {
		return errors.New("没有权限")
	}
	err = a.common.Delete(id)
	if err != nil {
		log.Println(err.Error())
		return errors.New("删除Event失败")
	}
	return nil
}
