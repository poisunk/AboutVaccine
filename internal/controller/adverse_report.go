package controller

import (
	"about-vaccine/internal/base/handler"
	"about-vaccine/internal/schema"
	"about-vaccine/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type AdverseReportController struct {
	AdverseEventService *service.AdverseReportService
}

func NewAdverseEventController(eventService *service.AdverseReportService) *AdverseReportController {
	return &AdverseReportController{
		AdverseEventService: eventService,
	}
}

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

func (a *AdverseReportController) GetAdverseEvent(c *gin.Context) {
	// 获取page, pageSize
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	// 1. 获取id, 是否通过id查询
	s := c.Query("id")
	if len(s) != 0 {
		// 根据id查询
		id, _ := strconv.ParseInt(s, 10, 64)
		adverseEvent, err := a.AdverseEventService.Get(id)
		handler.HandleResponse(c, err, adverseEvent)
		return
	}
	// 2. 获取uid，是否通过uid查询
	s = c.Query("keyword")
	if len(s) != 0 {
		//// 根据keyword查询
		adverseEvent, total, err := a.AdverseEventService.GetListByKeyword(s, page, pageSize)
		handler.HandleResponse(c, err, handler.PagedData{
			Total:    total,
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
