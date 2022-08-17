package api

import (
	"cilicili.live/cilicili/service"
	"github.com/gin-gonic/gin"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	var service service.CreateVideoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListVideo 视频列表
func ListVideo(c *gin.Context) {
	var service service.ListVideoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowVideo 视频详情
func ShowVideo(c *gin.Context) {
	var service service.ShowVideoService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
