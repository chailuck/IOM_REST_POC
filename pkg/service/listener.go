package service

import (
	iomcontroller "gitlab.com/ft25/iom/engine/am/internal/controller"
	iomlistener "gitlab.com/ft25/iom/framework/service"
)

var TokenService *iomcontroller.TokenService

type RestService struct {
	iomlistener.RestService
}

func NewRestServiceController(service *iomcontroller.TokenService) {
	TokenService = service
}
