package router

import (
	"github.com/novalwardhana/golang-boiler-plate/package/crud/controller"
)

type Router struct {
	controller controller.Controller
}

func NewRouter(controller controller.Controller) *Router {
	return &Router{
		controller: controller,
	}
}
