package controller

import (
	"about-vaccine/internal/base/handler"
	"about-vaccine/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type VaersController struct {
	service *service.VaersService
}

func NewVaersController(vaersService *service.VaersService) *VaersController {
	return &VaersController{
		service: vaersService,
	}
}

// SearchVaersResult 检索VaersResult数据
func (v *VaersController) SearchVaersResult(c *gin.Context) {
	vaccineId := c.Query("vaccineId")
	symptomId := c.Query("symptomId")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	if len(vaccineId) == 0 && len(symptomId) == 0 {
		handler.HandleResponse(c, errors.New("必须选择一个条件"), nil)
		return
	}
	if len(vaccineId) == 0 {
		// 通过symptomId检索
		id, _ := strconv.ParseInt(symptomId, 10, 64)
		vaers, total, err := v.service.GetResultBySymptomId(id, page, pageSize)
		handler.HandleResponse(c, err, handler.PagedData{
			Data:     vaers,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		})
		return
	} else if len(symptomId) == 0 {
		// 通过vaccineId检索
		id, _ := strconv.ParseInt(vaccineId, 10, 64)
		vaers, total, err := v.service.GetResultByVaccineId(id, page, pageSize)
		handler.HandleResponse(c, err, handler.PagedData{
			Data:     vaers,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		})
		return
	} else {
		// 检索vaccineId, symptomId
		vid, _ := strconv.ParseInt(vaccineId, 10, 64)
		sid, _ := strconv.ParseInt(symptomId, 10, 64)
		vaers, err := v.service.GetResult(vid, sid)
		handler.HandleResponse(c, err, vaers)
		return
	}
}

// GetVaersVaccineList 获取Vaers的Vaccine列表
// 可选参数：page, pageSize, keyword
func (v *VaersController) GetVaersVaccineList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	keyword := c.DefaultQuery("keyword", "")

	vaers, total, err := v.service.GetVaccineTermList(keyword, page, pageSize)
	handler.HandleResponse(c, err, handler.PagedData{
		Data:     vaers,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

// GetVaersSymptomList 获取Vaers的Symptom列表
// 可选参数：page, pageSize, keyword
func (v *VaersController) GetVaersSymptomList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	keyword := c.DefaultQuery("keyword", "")

	vaers, total, err := v.service.GetSymptomTermList(keyword, page, pageSize)
	handler.HandleResponse(c, err, handler.PagedData{
		Data:     vaers,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}
