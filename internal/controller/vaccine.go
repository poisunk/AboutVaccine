package controller

import (
	"about-vaccine/internal/base/handler"
	"about-vaccine/internal/service"
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

// GetVaccine 查询CFDA疫苗数据
// 可选参数: page, pageSize, productName
func (v *VaccineController) GetVaccine(c *gin.Context) {
	id := c.Query("id")
	// 根据id查询
	if len(id) != 0 {
		vaccine, err := v.service.Get(id)
		handler.HandleResponse(c, err, vaccine)
		return
	}

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

func (v *VaccineController) GetVaccineType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err == nil {
		v, err := v.service.GetTypeDetailInfo(id)
		handler.HandleResponse(c, err, v)
		return
	}
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
