package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/novalwardhana/golang-boiler-plate/package/crud/controller"
	"github.com/novalwardhana/golang-boiler-plate/package/crud/model"
)

type Router struct {
	controller controller.Controller
}

func NewRouter(controller controller.Controller) *Router {
	return &Router{
		controller: controller,
	}
}

func (r *Router) Mount(group *echo.Group) {
	group.POST("/add", r.add)
	group.PUT("/update", r.update)
	group.DELETE("/delete", r.delete)
	group.GET("/info", r.info)
}

func (r *Router) add(c echo.Context) error {
	var newUser model.User
	var err error

	if err = c.Bind(&newUser); err != nil {
		return c.String(http.StatusOK, "xxx")
	}
	newUser.IsActive = true

	addResult := <-r.controller.Add(newUser)
	fmt.Println(addResult)

	return c.String(http.StatusOK, "Shinning through the city")
}

func (r *Router) update(c echo.Context) error {
	return nil
}

func (r *Router) delete(c echo.Context) error {
	return nil
}

func (r *Router) info(c echo.Context) error {
	return nil
}
