package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"vax/internal/base/handler"
	"vax/internal/service"
)

type OAETermController struct {
	service *service.OaeTermService
}

func NewOAETermController(termService *service.OaeTermService) *OAETermController {
	return &OAETermController{
		service: termService,
	}
}

// GetOaeTermsByLabel 得到与label有关的oae词条
func (o *OAETermController) GetOaeTermsByLabel(c *gin.Context) {
	label := c.DefaultQuery("label", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	// 查询
	oaeList, total, err := o.service.GetBySimilarLabel(label, page, pageSize)
	handler.HandleResponse(c, err, handler.PagedData{
		Data:     oaeList,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	})
}

// GetOaeTermByIRI 通过IRI链接查询oae词条
func (o *OAETermController) GetOaeTermByIRI(c *gin.Context) {
	IRI := c.DefaultQuery("IRI", "")
	if IRI == "" {
		handler.HandleResponse(c, errors.New("IRI不能为空"), nil)
		return
	}
	// 查询
	oaeTerm, err := o.service.GetByIRI(IRI)
	handler.HandleResponse(c, err, oaeTerm)
}

// GetOaeTermParents 通过IRI链接查询其所有的父类
func (o *OAETermController) GetOaeTermParents(c *gin.Context) {
	IRI := c.DefaultQuery("IRI", "")
	if IRI == "" {
		handler.HandleResponse(c, errors.New("IRI不能为空"), nil)
		return
	}
	list, err := o.service.GetParents(IRI)
	handler.HandleResponse(c, err, list)
}

func (o *OAETermController) GetOaeTermByID(c *gin.Context) {
	ID := c.Param("id")
	if ID == "" {
		handler.HandleResponse(c, errors.New("ID不能为空"), nil)
		return
	}
	// 查询
	oaeTerm, err := o.service.GetByID(ID)
	handler.HandleResponse(c, err, oaeTerm)
}
