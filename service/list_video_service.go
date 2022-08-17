package service

import (
	"cilicili.live/cilicili/model"
	"cilicili.live/cilicili/serializer"
)

// ListVideoService 视频投稿服务
type ListVideoService struct {
}

// List 视频列表
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video

	err := model.DB.Find(&videos).Error

	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库查询失败",
			Error: err.Error(),
		}
	}

	return serializer.BuildListVideoResponse(videos)
}
