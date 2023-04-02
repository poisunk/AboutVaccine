package jwt

import (
	"MyWeb/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

const NoBodySign = 0

// Auth 验证token，提取uid
func Auth(c *gin.Context) {
	auth := c.Query("token")
	var uid string
	if len(auth) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    config.FailureStatus,
			"massage": "token is empty",
		})
		c.Abort()
	} else {
		token, err := ParseToken(auth)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    config.FailureStatus,
				"massage": "token is invalid",
			})
			c.Abort()
		} else {
			uid = token.Id
		}
	}
	c.Set("userId", uid)
	c.Next()
}
