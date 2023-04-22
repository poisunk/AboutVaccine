package controller

import (
	"about-vaccine/src/base/handler"
	"about-vaccine/src/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type VaccineController struct {
	service *service.VaccineService
}

func NewVaccineController(vaccineService *service.VaccineService) *VaccineController {
	return &VaccineController{
		service: vaccineService,
	}
}

// GetVaccineList 查询CFDA疫苗数据
// 可选参数: page, pageSize, productName
func (v *VaccineController) GetVaccineList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	typeStr := c.DefaultQuery("type", "")
	if len(typeStr) != 0 {
		list, total, err := v.service.GetByType(typeStr, page, pageSize)
		handler.HandleResponse(c, err, handler.PagedData{
			Data:     list,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		})
		return
	}
	productName := c.Query("productName")
	// 得到疫苗数据
	vList, total, err := v.service.GetByName(productName, page, pageSize)
	// 返回数据
	handler.HandleResponse(c, err, handler.PagedData{
		Data:     vList,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

func (v *VaccineController) GetVaccineTypeList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	list, total, err := v.service.GetTypeList(page, pageSize)
	handler.HandleResponse(c, err, handler.PagedData{
		Data:     list,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}
