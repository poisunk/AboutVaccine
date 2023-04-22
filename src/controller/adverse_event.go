package controller

import (
	"about-vaccine/src/base/handler"
	"about-vaccine/src/schama"
	"about-vaccine/src/service"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AdverseEventController struct {
	AdverseEventService *service.AdverseEventService
}

func NewAdverseEventController(eventService *service.AdverseEventService) *AdverseEventController {
	return &AdverseEventController{
		AdverseEventService: eventService,
	}
}

func (a *AdverseEventController) CreateAdverseEvent(c *gin.Context) {
	// 获取参数
	var adverseEvent = new(schama.AdverseEvent)
	err := c.ShouldBindJSON(adverseEvent)
	if err != nil {
		handler.HandleResponse(c, errors.New("json格式错误"), nil)
		return
	}
	uid := c.GetString("userId")
	if len(uid) != 0 {
		adverseEvent.Uid, _ = strconv.ParseInt(uid, 10, 64)
	}
	// 创建服务
	err = a.AdverseEventService.Create(adverseEvent)
	handler.HandleResponse(c, err, nil)
}

func (a *AdverseEventController) GetAdverseEvent(c *gin.Context) {
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
	s = c.GetString("userId")
	if len(s) != 0 {
		// 根据uid查询
		uid, _ := strconv.ParseInt(s, 10, 64)
		adverseEvent, total, err := a.AdverseEventService.GetListByUid(uid, page, pageSize)
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

func (a *AdverseEventController) DeleteAdverseEvent(c *gin.Context) {
	// 获取id
	s := c.Query("id")
	if len(s) == 0 {
		handler.HandleResponse(c, errors.New("id不能为空"), nil)
		return
	}
	id, _ := strconv.ParseInt(s, 10, 64)
	// 根据id删除
	err := a.AdverseEventService.Delete(id)
	handler.HandleResponse(c, err, nil)
}
