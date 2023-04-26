package controller

import (
	"about-vaccine/internal/base/handler"
	"about-vaccine/internal/config"
	"about-vaccine/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
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
	claim, err := u.UserService.Register(name, password)
	if err != nil {
		handler.HandleResponse(c, err, nil)
		return
	}
	t := c.DefaultQuery("type", "json")
	if t == "json" {
		handler.HandleResponse(c, err, claim)
	} else if t == "cookie" {
		c.SetCookie(config.UserClaimCookie, claim.Token, int(time.Hour.Milliseconds()*24),
			"/", "localhost", false, true)
		handler.HandleResponse(c, err, nil)
	}
}

// Login 用户登录
func (u *UserController) Login(c *gin.Context) {
	// 1. 尝试使用cookie登录
	cookie, _ := c.Cookie(config.UserClaimCookie)
	if len(cookie) != 0 {
		claim, err := u.UserService.LoginWithToken(cookie)
		t := c.DefaultQuery("type", "cookie")
		handler.HandleClaimResponse(c, err, t, claim)
		return
	}
	// 2. 尝试使用token登录
	token := c.DefaultQuery("token", "")
	if len(token) != 0 {
		claim, err := u.UserService.LoginWithToken(token)
		t := c.DefaultQuery("type", "token")
		handler.HandleClaimResponse(c, err, t, claim)
		return
	}
	// 3. 使用用户密码登录
	name := c.DefaultQuery("username", "")
	password := c.DefaultQuery("password", "")
	claim, err := u.UserService.Login(name, password)
	t := c.DefaultQuery("type", "token")
	handler.HandleClaimResponse(c, err, t, claim)
}

// Logout 用户注销
func (u *UserController) Logout(c *gin.Context) {
	cookie, _ := c.Cookie(config.UserClaimCookie)
	if len(cookie) != 0 {
		err := u.UserService.Logout(cookie)
		if err == nil {
			c.SetCookie(config.UserClaimCookie, cookie, -1,
				"/", "localhost", false, true)
			handler.HandleResponse(c, err, nil)
			return
		}
	}
	token := c.DefaultQuery("token", "")
	if len(token) != 0 {
		err := u.UserService.Logout(token)
		handler.HandleResponse(c, err, nil)
		return
	}
}

// SearchUser 搜索用户
func (u *UserController) SearchUser(c *gin.Context) {
	keyword := c.DefaultQuery("username", "")
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
