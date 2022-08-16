package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": "0",
		"msg":    "成功",
	})
}
