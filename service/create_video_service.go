package service

import (
	"cilicili.live/cilicili/model"
	"cilicili.live/cilicili/serializer"
)

// CreateVideoService 视频投稿服务
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=5,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
}

// Create 创建视频
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}

	err := model.DB.Create(&video).Error

	if err != nil {
		return serializer.Response{
			Code:  50001,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}

	return serializer.BuildVideoResponse(video)
}
