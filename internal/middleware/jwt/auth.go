package jwt

import (
	"about-vaccine/internal/base/handler"
	"errors"
	"github.com/gin-gonic/gin"
)

// Auth 验证token
func Auth(c *gin.Context) {
	auth := c.Query("token")
	var uid string
	if len(auth) == 0 {
		auth, _ = c.Cookie("user_token")
		if len(auth) == 0 {
			handler.HandleResponse(c, errors.New("token is empty"), nil)
			c.Abort()
			return
		}
	}
	token, err := ParseToken(auth)
	if err != nil {
		handler.HandleResponse(c, errors.New("token is invalid"), nil)
		c.Abort()
	} else {
		uid = token.Id
	}
	c.Set("userId", uid)
	c.Next()
}

// AuthWithoutLogin 验证token
func AuthWithoutLogin(c *gin.Context) {
	auth := c.Query("token")
	if len(auth) == 0 {
		auth, _ = c.Cookie("user_token")
	}
	var uid string
	token, err := ParseToken(auth)
	if err == nil {
		uid = token.Id
		c.Set("userId", uid)
	}
	c.Next()
}
