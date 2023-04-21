package controller

import (
	"about-vaccine/src/config"
	"about-vaccine/src/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"time"
)

func CreateQuestionnaire(c *gin.Context) {
	// 检查当前是否登录
	uid, _ := strconv.ParseInt(c.GetString("userId"), 10, 64)
	// 获取传送过来的数据
	var q service.Questionnaire
	if err := c.ShouldBindBodyWith(&q, binding.JSON); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "参数错误!",
		})
		return
	}

	// 问卷名不能为空
	if len(q.Name) == 0 {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "问卷名不能为空!",
		})
		return
	}
	// 准备数据
	q.CreateTime = time.Now()
	q.OwnerId = uid

	// 准备服务
	var qService = service.InitQuestionnaireService()
	err := qService.CreateQuestionnaire(&q)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "创建失败!",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "创建成功!",
		Data:    q,
	})
}

func GetQuestionnaireByUid(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    http.StatusBadRequest,
			Message: "参数错误！",
		})
		return
	}

	// 准备服务
	var qService = service.InitQuestionnaireService()
	q, err := qService.GetQuestionnaireByUid(int64(id))
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "获取失败!",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "获取成功!",
		Data:    q,
	})

}

func DeleteQuestionnaireById(c *gin.Context) {
	uid, _ := strconv.ParseInt(c.GetString("userId"), 10, 64)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	// 准备服务
	var qService = service.InitQuestionnaireService()
	o, err := qService.GetQuestionnaireOwnerIdById(id)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	// 检查是否有权限
	if o != uid {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "没有权限！",
		})
		return
	}

	// 删除
	err = qService.DeleteQuestionnaireById(int64(id))
	if err != nil {
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

func GetQuestionnaireList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	// 准备服务
	var qService = service.InitQuestionnaireService()
	q, total, err := qService.GetQuestionnaireList(page, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	more := int(total)-page*pageSize > 0
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "获取成功！",
		Data: gin.H{
			"page":     page,
			"pageSize": pageSize,
			"more":     more,
			"data":     q,
			"total":    total,
		},
	})
}
