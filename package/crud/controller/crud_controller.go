package controller

import (
	"github.com/novalwardhana/golang-boiler-plate/package/crud/service"
)

type controller struct {
	service service.Service
}

type Controller interface {
}

func NewController(service service.Service) Controller {
	return &controller{
		service: service,
	}
}
