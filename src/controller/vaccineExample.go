package controller

import (
	"about-vaccine/src/config"
	"about-vaccine/src/models"
	"about-vaccine/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VaccineExample struct {
	Type         string            `json:"type"`
	VaccineCFDAs []*models.Vaccine `json:"vaccines"`
	HasMore      bool              `json:"hasMore"`
}

// GetVaccineListExampleByTId 通过类型获得示例疫苗列表
func GetVaccineListExampleByTId(c *gin.Context) {
	// 获取查询数量，默认为5
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))
	// 检查tid
	// 获取查询类型的tid，且不能为空
	tid, err := strconv.ParseInt(c.Param("tid"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    config.FailureStatus,
			Message: "参数错误！",
		})
		return
	}
	// 初始化服务
	vService := service.InitVaccineService()
	// 查询相应类型的疫苗数据
	example, err := vService.GetVaccineExampleByTid(tid, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "查询成功！",
		Data:    example,
	})
}

// GetVaccineListExample 获取所有类型的疫苗展示数据
func GetVaccineListExample(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))
	// 每个类型最大数据量
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	// 初始化服务
	vService := service.InitVaccineService()
	// 查询
	list, total, err := vService.GetVaccineExampleList(page, pageSize, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	more := (page-1)*pageSize+len(list) < total
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "查询成功！",
		Data: gin.H{
			"exampleList": list,
			"total":       total,
			"page":        page,
			"pageSize":    pageSize,
			"more":        more,
		},
	})

}
