package controller

import (
	"about-vaccine/src/base/handler"
	"about-vaccine/src/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// Register 用户注册
// @(username)
func (u *UserController) Register(c *gin.Context) {
	// 得到用户名与密码
	name := c.DefaultQuery("username", "")
	password := c.DefaultQuery("password", "")
	// 注册，得到token
	token, err := u.UserService.Register(name, password)
	if err != nil {
		handler.HandleResponse(c, err, nil)
		return
	}
	t := c.DefaultQuery("type", "json")
	if t == "json" {
		handler.HandleResponse(c, err, gin.H{
			token: token,
		})
	} else if t == "cookie" {
		c.SetCookie("user_token", token, 3600, "/", "localhost", false, true)
		handler.HandleResponse(c, err, nil)
	}
}

// Login 用户登录
func (u *UserController) Login(c *gin.Context) {
	// 1. 如果有token，先尝试使用token登录
	token := c.DefaultQuery("token", "")
	if token != "" {
		newToken, err := u.UserService.LoginWithToken(token)
		if err == nil {
			handler.HandleResponse(c, err, gin.H{
				token: newToken,
			})
			return
		}
	}
	// 2. 使用cookie登录
	cookie, err := c.Cookie("user_token")
	if cookie != "" {
		newToken, err := u.UserService.LoginWithToken(cookie)
		if err == nil {
			handler.HandleResponse(c, err, gin.H{
				token: newToken,
			})
			return
		}
	}
	// 3. 使用用户名与密码登录
	name := c.DefaultQuery("username", "")
	password := c.DefaultQuery("password", "")
	token, err = u.UserService.Login(name, password)
	handler.HandleResponse(c, err, gin.H{
		token: token,
	})
}

// Logout 用户注销
func (u *UserController) Logout(c *gin.Context) {
	token := c.Query("token")
	err := u.UserService.Logout(token)
	handler.HandleResponse(c, err, nil)
}

// SearchUser 搜索用户
func (u *UserController) SearchUser(c *gin.Context) {
	keyword := c.DefaultQuery("keyword", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	users, total, err := u.UserService.SearchUserByName(keyword, page, pageSize)
	handler.HandleResponse(c, err, handler.PagedData{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		Data:     users,
	})
}
