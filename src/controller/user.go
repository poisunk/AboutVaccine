package controller

import (
	"about-vaccine/src/config"
	"about-vaccine/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Register 用户注册
func Register(c *gin.Context) {
	// 得到用户名与密码
	name := c.DefaultQuery("username", "")
	password := c.DefaultQuery("password", "")

	// 准备服务
	uService := service.InitUserService()
	// 注册，得到token
	token, err := uService.Register(service.User{
		Nickname: name,
		Password: password,
	})
	// 注册失败
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	// 注册成功
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "注册成功！",
		Data:    token,
	})
}

// Login 用户登录
func Login(c *gin.Context) {
	// 得到用户名与密码
	name := c.DefaultQuery("username", "")
	password := c.DefaultQuery("password", "")

	// 准备服务
	uService := service.InitUserService()
	// 登录，得到token
	token, err := uService.Login(service.User{
		Nickname: name,
		Password: password,
	})

	// 登录失败
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	// 登录成功
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "登录成功！",
		Data:    token,
	})
}

// Status 检查token
func Status(c *gin.Context) {
	token := c.Query("token")
	//准备服务
	uService := service.InitUserService()
	// 检验token
	newToken, err := uService.Status(token)
	// 无效token
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	// 有效token
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "已登录！",
		Data:    newToken,
	})
}

func Logout(c *gin.Context) {
	token := c.Query("token")
	// 准备服务
	uService := service.InitUserService()
	// 注销
	err := uService.Logout(token)
	// 注销失败
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	// 注销成功
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "注销成功！",
	})
}

func GetUserList(c *gin.Context) {
	// page, pageSize
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	// 准备服务
	uService := service.InitUserService()
	// 获取用户列表
	userList, total, err := uService.GetUserList(page, pageSize)
	// 获取失败
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    config.FailureStatus,
			Message: err.Error(),
		})
		return
	}
	// 获取成功
	more := pageSize*(page-1)+len(userList) < int(total)
	c.JSON(http.StatusOK, Response{
		Code:    config.SuccessStatus,
		Message: "获取用户列表成功！",
		Data: gin.H{
			"userList": userList,
			"total":    total,
			"more":     more,
			"page":     page,
			"pageSize": pageSize,
		},
	})
}
