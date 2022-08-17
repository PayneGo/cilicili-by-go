package service

import (
	"cilicili.live/cilicili/model"
	"cilicili.live/cilicili/serializer"
)

// ShowVideoService 视频详情服务
type ShowVideoService struct {
}

// List 视频列表
func (service *ShowVideoService) Show(id string) serializer.Response {

	var video model.Video

	err := model.DB.First(&video, id).Error

	if err != nil {
		return serializer.Response{
			Code:  50004,
			Msg:   "视频不存在",
			Error: err.Error(),
		}
	}

	return serializer.BuildVideoResponse(video)
}
