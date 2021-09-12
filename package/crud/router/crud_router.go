package router

import (
	"fmt"
	"net/http"
	"strconv"

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
	group.PUT("/update/:id", r.update)
	group.DELETE("/delete", r.delete)
	group.GET("/info", r.info)
}

func (r *Router) add(c echo.Context) error {
	var newUser model.User
	var err error
	var response model.Response

	if err = c.Bind(&newUser); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Invalid payload"
		return c.JSON(http.StatusOK, response)
	}

	addResult := <-r.controller.Add(newUser)
	if addResult.Error != nil {
		response.StatusCode = http.StatusUnprocessableEntity
		response.Message = addResult.Error.Error()
		return c.JSON(http.StatusOK, response)
	}

	response.Message = "Succes add data"
	response.StatusCode = 200
	response.Data = addResult.Data
	return c.JSON(http.StatusOK, response)
}

func (r *Router) update(c echo.Context) error {
	var id int
	var user model.User
	var err error
	var response model.Response

	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "ID not valid"
		return c.JSON(http.StatusOK, response)
	}

	if err = c.Bind(&user); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Invalid payload"
		return c.JSON(http.StatusOK, response)
	}

	fmt.Println(id, user)

	return nil
}

func (r *Router) delete(c echo.Context) error {
	return nil
}

func (r *Router) info(c echo.Context) error {
	return nil
}
