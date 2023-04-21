package controller

import (
	"about-vaccine/src/config"
	"about-vaccine/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetOaeTermsByLabel 得到与label有关的oae词条
func GetOaeTermsByLabel(c *gin.Context) {
	label := c.DefaultQuery("label", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	// 创建服务
	var oService = service.InitOAEService()
	// 查询
	oaeList, total, err := oService.GetOAETermByLabel(label, page, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	more := total-page*pageSize > 0
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "查询成功！",
		Data: gin.H{
			"page":     page,
			"pageSize": pageSize,
			"total":    total,
			"more":     more,
			"oaeTerms": oaeList,
		},
	})
}

// GetOaeTermByIRI 通过IRI链接查询oae词条
func GetOaeTermByIRI(c *gin.Context) {
	IRI := c.DefaultQuery("IRI", "")
	if IRI == "" {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "请求参数不能为空！",
		})
		return
	}

	// 创建服务
	var oService = service.InitOAEService()
	// 查询
	oaeTerm, err := oService.GetOAETermByIRI(IRI)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "查询成功！",
		Data:    oaeTerm,
	})
}

// GetOaeParentTermsByIRI 通过IRI链接查询其所有的父类
func GetOaeParentTermsByIRI(c *gin.Context) {
	IRI := c.DefaultQuery("IRI", "")
	if IRI == "" {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "请求参数不能为空！",
		})
		return
	}

	// 创建服务
	var oService = service.InitOAEService()
	var parentTerms []*service.OAETerm
	parentTerms, err := oService.GetOAETermParentList(IRI)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "查询成功！",
		Data:    parentTerms,
	})
}
