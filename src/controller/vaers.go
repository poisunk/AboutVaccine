package controller

import (
	"MyWeb/config"
	"MyWeb/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// SearchVaers 检索Vaers数据
// 可选参数：vaccineId, symptomId, page, pageSize
func SearchVaers(c *gin.Context) {
	vaccineId := c.Query("vaccineId")
	symptomId := c.Query("symptomId")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	if len(vaccineId) == 0 && len(symptomId) == 0 {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "两个参数不能同时为空",
		})
		return
	}

	// 准备服务
	vaersService := service.InitVaersService()
	if len(vaccineId) == 0 {
		// 检索symptomId
		id, _ := strconv.ParseInt(symptomId, 10, 64)
		vaers, total, err := vaersService.GetVaersResultsBySymptomId(id, page, pageSize)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				Code:    config.FailureStatus,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, Response{
			Code:    config.SuccessStatus,
			Message: "查询成功",
			Data: gin.H{
				"results":  vaers,
				"total":    total,
				"page":     page,
				"pageSize": pageSize,
			},
		})
		return
	} else if len(symptomId) == 0 {
		// 检索vaccineId
		id, _ := strconv.ParseInt(vaccineId, 10, 64)
		vaers, total, err := vaersService.GetVaersResultsByVaccineId(id, page, pageSize)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				Code:    config.FailureStatus,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, Response{
			Code:    config.SuccessStatus,
			Message: "查询成功",
			Data: gin.H{
				"results":  vaers,
				"total":    total,
				"page":     page,
				"pageSize": pageSize,
			},
		})
		return
	} else {
		// 检索vaccineId, symptomId
		id, _ := strconv.ParseInt(vaccineId, 10, 64)
		id2, _ := strconv.ParseInt(symptomId, 10, 64)
		vaers, err := vaersService.GetVaersResults(id, id2)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				Code:    config.FailureStatus,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, Response{
			Code:    config.SuccessStatus,
			Message: "查询成功",
			Data:    vaers,
		})
	}
}

// GetVaers 获取Vaers数据
// 必选参数：vaersId
func GetVaers(c *gin.Context) {
	vaersId, _ := strconv.ParseInt(c.Param("vaersId"), 10, 64)

	// 准备服务
	vaersService := service.InitVaersService()
	vaers, err := vaersService.GetVaersByVaersId(vaersId)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "查询成功",
		Data:    vaers,
	})
}

// GetVaersVaccineList 获取Vaers的Vaccine列表
// 可选参数：page, pageSize, keyword
func GetVaersVaccineList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	keyword := c.DefaultQuery("keyword", "")

	// 准备服务
	vaersService := service.InitVaersVaxService()
	vaers, total, err := vaersService.GetVaersVaxTermList(keyword, page, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "查询成功",
		Data: gin.H{
			"vaccines": vaers,
			"page":     page,
			"pageSize": pageSize,
			"keyword":  keyword,
			"total":    total,
		},
	})
}

// GetVaersSymptomList 获取Vaers的Symptom列表
// 可选参数：page, pageSize, keyword
func GetVaersSymptomList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	keyword := c.DefaultQuery("keyword", "")

	// 准备服务
	vaersService := service.InitVaersSymptomService()
	vaers, total, err := vaersService.GetVaersSymptomTermList(keyword, page, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "查询成功",
		Data: gin.H{
			"symptoms": vaers,
			"page":     page,
			"pageSize": pageSize,
			"keyword":  keyword,
			"total":    total,
		},
	})
}
