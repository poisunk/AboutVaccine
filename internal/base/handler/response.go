package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"vax/internal/config"
	"vax/internal/schema"
)

type RespBody struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PagedData struct {
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Total    int64       `json:"total,omitempty"`
	Data     interface{} `json:"data"`
}

func HandleResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		c.JSON(http.StatusOK, RespBody{
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

func HandleClaimResponse(c *gin.Context, err error, t string, claim *schema.UserClaim) {
	if t == "cookie" {
		if err == nil && claim != nil {
			c.SetCookie(config.UserClaimCookie, claim.Token, int(time.Hour.Seconds()*24),
				"/", "localhost", false, true)
		} else if err != nil {
			c.SetCookie(config.UserClaimCookie, "", -1, "/", "localhost", false, true)
		}
		HandleResponse(c, err, nil)
		return
	}
	c.SetCookie(config.UserClaimCookie, "", -1, "/", "localhost", false, true)
	HandleResponse(c, err, claim)
}
