package service

import (
	"GoRoLingG/service/image_service"
	"GoRoLingG/service/redis_service"
	"GoRoLingG/service/user_service"
)

type ServiceGroup struct {
	ImageService image_service.ImageService
	UserService  user_service.UserService
	RedisService redis_service.RedisService
}

var Service = new(ServiceGroup)
