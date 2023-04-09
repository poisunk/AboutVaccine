package controller

import (
	"MyWeb/config"
	"MyWeb/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetQuestionListByQId 获取问卷列表，通过id
func GetQuestionListByQId(c *gin.Context) {
	// 获取路径参数
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "参数错误!",
		})
		return
	}

	// 检查问卷是否存在
	qaService := service.InitQuestionnaireService()
	q, _ := qaService.GetQuestionnaireById(id)
	if q == nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "问卷不存在!",
		})
		return
	}

	// 初始化服务
	qService := service.InitQuestionService()
	list, err := qService.GetQuestionList(id)
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
		Data:    list,
	})
}

// CreateQuestion 创建问卷的问题列表
func CreateQuestion(c *gin.Context) {
	// 获取要创建的问卷的id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "参数错误!",
		})
		return
	}
	// 获取当前用户id
	uid, _ := strconv.Atoi(c.GetString("userId"))
	// 准备服务
	qaService := service.InitQuestionnaireService()
	o, err := qaService.GetQuestionnaireOwnerIdById(id)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	// 检查是否有权限
	if o != int64(uid) {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "没有权限！",
		})
		return
	}

	// 获取body中的问题列表
	var q []*service.Question
	if err := c.ShouldBindJSON(&q); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "question数据格式错误！",
		})
		return
	}
	qService := service.InitQuestionService()
	for _, v := range q {
		v.QuestionnaireId = id
	}
	err = qService.CreateQuestion(q)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "创建成功！",
	})
}

func GetQuestionTypeList(c *gin.Context) {
	qService := service.InitQuestionService()
	list := qService.GetQuestionTypeList()
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "获取成功!",
		Data:    list,
	})
}

func DeleteQuestion(c *gin.Context) {
	uid, _ := strconv.ParseInt(c.GetString("userId"), 10, 64)
	// 获取要删除的问题所在问卷的id
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	// 问题id
	qid, err := strconv.ParseInt(c.Param("qid"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "参数错误!",
		})
		return
	}

	// 准备服务
	qaService := service.InitQuestionnaireService()
	o, err := qaService.GetQuestionnaireOwnerIdById(id)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	// 检查是否有权限
	if uid != o {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "没有权限！",
		})
		return
	}

	qService := service.InitQuestionService()
	err = qService.DeleteQuestionById(qid)
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
