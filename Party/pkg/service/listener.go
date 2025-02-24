package service

import (
	iomcontroller "internal/controller"

	iomlistener "gitlab.com/ft25/iom/framework/service"
)

var PartyService *iomcontroller.PartyService

type RestService struct {
	iomlistener.RestService
}

func NewRestServiceController(service *iomcontroller.PartyService) {
	PartyService = service
}
