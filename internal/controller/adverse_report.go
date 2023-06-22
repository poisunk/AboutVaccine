package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"vax/internal/base/handler"
	"vax/internal/schema"
	"vax/internal/service"
)

type AdverseReportController struct {
	AdverseEventService *service.AdverseReportService
}

func NewAdverseEventController(eventService *service.AdverseReportService) *AdverseReportController {
	return &AdverseReportController{
		AdverseEventService: eventService,
	}
}

// CreateAdverseEvent 创建不良反应报告
func (a *AdverseReportController) CreateAdverseEvent(c *gin.Context) {
	// 获取参数
	var adverseEvent = new(schema.AdverseEventAdd)
	err := c.ShouldBindJSON(adverseEvent)
	if err != nil {
		log.Println(err.Error())
		handler.HandleResponse(c, errors.New("json格式错误"), nil)
		return
	}
	token := c.Query("token")
	// 创建服务
	err = a.AdverseEventService.Create(adverseEvent, token)
	handler.HandleResponse(c, err, nil)
}

// GetAdverseEvent 查询不良反应报告
// @param id: 不良反应报告id
// @param page, pageSize
func (a *AdverseReportController) GetAdverseEvent(c *gin.Context) {
	// 获取page, pageSize
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	// 1. 通过id查询
	id := c.Query("id")
	if len(id) != 0 {
		// 根据id查询
		id, _ := strconv.ParseInt(id, 10, 64)
		adverseEvent, err := a.AdverseEventService.Get(id)
		handler.HandleResponse(c, err, adverseEvent)
		return
	}
	// 2. 通过keyword查询
	keyword := c.Query("keyword")
	if len(keyword) != 0 {
		//// 根据keyword查询
		adverseEvent, total, err := a.AdverseEventService.GetListByKeyword(keyword, page, pageSize)
		handler.HandleResponse(c, err, handler.PagedData{
			Total:    total,
			Page:     page,
			PageSize: pageSize,
			Data:     adverseEvent,
		})
		return
	}
	// 3. 通过vaccineId查询
	vaccineId := c.Query("vaccineId")
	if len(vaccineId) != 0 {
		// 根据vaccineId查询
		adverseEvent, err := a.AdverseEventService.GetListByVaccineId(vaccineId, page, pageSize)
		handler.HandleResponse(c, err, handler.PagedData{
			Page:     page,
			PageSize: pageSize,
			Data:     adverseEvent,
		})
		return
	}
	// 4. 通过oaeId查询
	oaeId := c.Query("oaeId")
	if len(oaeId) != 0 {
		// 根据oaeTerm查询
		adverseEvent, err := a.AdverseEventService.GetListByOAEId(oaeId, page, pageSize)
		handler.HandleResponse(c, err, handler.PagedData{
			Page:     page,
			PageSize: pageSize,
			Data:     adverseEvent,
		})
		return
	}
	// 获取所有
	adverseEventList, total, err := a.AdverseEventService.GetList(page, pageSize)
	handler.HandleResponse(c, err, handler.PagedData{
		Total:    total,
		Page:     page,
		PageSize: pageSize,
		Data:     adverseEventList,
	})
}

func (a *AdverseReportController) GetAdverseResult(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	// 1. 获取vaccineId, oaeId
	vaccineId := c.Query("vaccineId")
	oaeId := c.Query("oaeId")
	// 2. 根据vaccineId, oaeId查询
	adverseEvent, err := a.AdverseEventService.GetResult(vaccineId, oaeId, page, pageSize)
	handler.HandleResponse(c, err, handler.PagedData{
		Page:     page,
		PageSize: pageSize,
		Data:     adverseEvent,
	})
}

func (a *AdverseReportController) DeleteAdverseEvent(c *gin.Context) {
	// 获取id
	token := c.Query("token")
	s := c.Query("id")
	if len(s) == 0 {
		handler.HandleResponse(c, errors.New("id不能为空"), nil)
		return
	}
	id, _ := strconv.ParseInt(s, 10, 64)
	// 根据id删除
	err := a.AdverseEventService.Delete(id, token)
	handler.HandleResponse(c, err, nil)
}
