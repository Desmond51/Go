package service

import "github.com/Desmond51/Go/hello/entity"

type VideoService interface{
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct{
	videos []entity.Video
}
func New() VideoService {
	
}