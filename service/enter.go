package service

import "GoRoLingG/service/image_service"

type ServiceGroup struct {
	ImageService image_service.ImageService
}

var Service = new(ServiceGroup)
