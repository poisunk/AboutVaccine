package controller

import (
	"about-vaccine/config"
	"about-vaccine/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetResponse(c *gin.Context) {
	uid, _ := strconv.ParseInt(c.GetString("userId"), 10, 64)
	qid, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	// 查询目标问卷
	qaService := service.InitQuestionnaireService()
	o, err := qaService.GetQuestionnaireOwnerIdById(qid)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	// 检查权限
	if uid != o {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "权限不足!",
		})
		return
	}

	rService := service.InitResponseService()
	r, total, err := rService.GetResponsesByQid(qid, page, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	more := pageSize*(page-1)+len(r) < total
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "获取成功!",
		Data: gin.H{
			"r":        r,
			"total":    total,
			"page":     page,
			"pageSize": pageSize,
			"more":     more,
		},
	})
}

func DeleteResponse(c *gin.Context) {
	uid, _ := strconv.ParseInt(c.GetString("userId"), 10, 64)
	qid, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	// 获取目标responseId
	s := c.Query("responseId")
	if s == "" {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "目标responseId不能为空!",
		})
		return
	}
	id, _ := strconv.ParseInt(s, 10, 64)
	// 查询目标问卷
	qaService := service.InitQuestionnaireService()
	o, err := qaService.GetQuestionnaireOwnerIdById(qid)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	// 检查权限
	if uid != o {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "权限不足!",
		})
		return
	}
	rService := service.InitResponseService()
	err = rService.DeleteResponse(id)
	if err != nil {

		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "删除成功!",
	})
}

func GetMineResponse(c *gin.Context) {
	uid, _ := strconv.ParseInt(c.GetString("userId"), 10, 64)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	rService := service.InitResponseService()
	r, total, err := rService.GetResponseByUid(uid, page, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	more := pageSize*(page-1)+len(r) < total
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "获取成功!",
		Data: gin.H{
			"r":        r,
			"total":    total,
			"page":     page,
			"pageSize": pageSize,
			"more":     more,
		},
	})
}

func CreateResponse(c *gin.Context) {
	uid, _ := strconv.ParseInt(c.GetString("userId"), 10, 64)
	qid, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	// 获取Response数据
	var r *service.Response
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	r.UserId = uid
	r.QuestionnaireId = qid
	rService := service.InitResponseService()
	err := rService.CreateResponse(r)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "创建成功!",
	})
}
