package controller

import (
	"about-vaccine/src/config"
	"about-vaccine/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetVaccineTypeList 获取疫苗的所有类型
func GetVaccineTypeList(c *gin.Context) {
	// 获取page，pageSize
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	// 创建服务
	tService := service.InitVaccineTypeService()
	// 调用服务
	typeList, err := tService.GetVaccineTypeList(page, pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    config.FailureStatus,
			Message: "查询失败！",
		})
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "查询成功！",
		Data:    typeList,
	})
}
