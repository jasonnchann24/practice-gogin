package services

import "github.com/jasonnchann24/gogin-rest/entities"

type VideoService interface {
	FindAll() []entities.Video
	Save(entities.Video) entities.Video
}

type videoService struct {
	videos []entities.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entities.Video) entities.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entities.Video {
	return service.videos
}
