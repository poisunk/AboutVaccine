package handler

import (
	"about-vaccine/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespBody struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PagedData struct {
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Total    int64       `json:"total"`
	Data     interface{} `json:"data"`
}

func HandleResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		c.JSON(http.StatusBadGateway, RespBody{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	if data == nil {
		c.JSON(http.StatusOK, RespBody{
			Code:    config.SuccessStatus,
			Message: "ok",
		})
		return
	}
	c.JSON(http.StatusOK, RespBody{
		Code:    config.SuccessStatus,
		Message: "ok",
		Data:    data,
	})
}
