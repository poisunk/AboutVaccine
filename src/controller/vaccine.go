package controller

import (
	"MyWeb/config"
	"MyWeb/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// CreateAVaccineCFDA 提交一条CFDA疫苗数据
func CreateAVaccineCFDA(c *gin.Context) {
	// 接受传来的数据
	var v = new(service.Vaccine)
	err := c.BindJSON(v)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "提交数据格式有误！",
		})
		return
	}

	// 创建服务
	var vService = service.InitVaccineService()
	err = vService.CreateVaccine(*v)
	// 判断是否有错误
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "提交成功！",
	})
}

// GetVaccineCFDAList 查询CFDA疫苗数据
// 可选参数: page, pageSize, productName
func GetVaccineCFDAList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	// 得到请求的关键词
	productName := c.Query("productName")
	// 创建服务
	var vService = service.InitVaccineService()
	// 得到疫苗数据
	vList, total, err := vService.GetVaccineList(page, pageSize, productName)
	// 返回数据
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
			"page":        page,
			"pageSize":    pageSize,
			"total":       total,
			"more":        more,
			"vaccineList": vList,
		},
	})
}

func DeleteVaccineCFDAById(c *gin.Context) {
	s := c.Query("id")
	if len(s) == 0 {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: "id不能为空！",
		})
		return
	}
	id, _ := strconv.Atoi(s)

	// 创建服务
	var vService = service.InitVaccineService()
	// 删除疫苗
	if err := vService.DeleteVaccine(int64(id)); err != nil {
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
