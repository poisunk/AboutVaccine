package controller

import (
	"MyWeb/config"
	"MyWeb/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateAdverseEvent(c *gin.Context) {
	// 获取参数
	var adverseEvent = new(service.AdverseEvent)
	if err := c.ShouldBindJSON(adverseEvent); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "参数错误！",
		})
		return
	}
	// 检查登录信息，并得到用户Id
	uid, _ := strconv.ParseInt(c.GetString("userId"), 10, 64)
	adverseEvent.Uid = uid

	// 检查数据
	if adverseEvent.Description == "" {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "不良反应描述不能为空！",
		})
		return
	}

	// 创建服务
	var aService = service.InitAdverseEventService()
	if err := aService.CreateAdverseEvent(*adverseEvent); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "创建成功！",
		Data:    adverseEvent,
	})
}

func GetAdverseEvent(c *gin.Context) {
	// 获取id
	s := c.Query("id")
	// 初始化服务
	var aService = service.InitAdverseEventService()
	if len(s) == 0 {
		// 获取page, pageSize
		page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
		pageSize, _ := strconv.ParseInt(c.DefaultQuery("pageSize", "20"), 10, 64)
		// 获取所有
		adverseEventList, total, err := aService.GetAdverseEventList(page, pageSize)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				Code:    config.FailureStatus,
				Message: err.Error(),
			})
			return
		}
		more := pageSize*(page-1)+int64(len(adverseEventList)) < total
		c.JSON(http.StatusOK, Response{
			Code:    config.SuccessStatus,
			Message: "查询成功！",
			Data: gin.H{
				"eventList": adverseEventList,
				"page":      page,
				"pageSize":  pageSize,
				"total":     total,
				"more":      more,
			},
		})
		return
	}
	id, _ := strconv.ParseInt(s, 10, 64)
	// 根据id查询
	var adverseEvent *service.AdverseEvent
	var err error
	if adverseEvent, err = aService.GetAdverseEvent(id); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "查询成功！",
		Data:    adverseEvent,
	})
}

func DeleteAdverseEvent(c *gin.Context) {
	// 获取id
	s := c.Query("id")
	if len(s) == 0 {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "id不能为空！",
		})
		return
	}
	id, _ := strconv.ParseInt(s, 10, 64)
	// 获取服务
	var aService = service.InitAdverseEventService()
	// 根据id删除
	if err := aService.DeleteAdverseEvent(id); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "删除成功！",
	})
}
