package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func IndexHandler(c *gin.Context) {
	names := []string{"Alice", "Bob", "Charlie", "David", "Alice", "Bob", "Charlie", "David"}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Items": names,
	})
}

func DownloadVaccineApk(c *gin.Context) {
	c.Header("content-type", "application/vnd.android.package-archive")
	c.Header("Content-Disposition", "attachment; filename=app-release.apk")
	c.Header("Content-Transfer-Encoding", "binary")
	c.File("./static/apk/app-release.apk")
}
