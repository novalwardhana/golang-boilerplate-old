package router

import (
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
	group.GET("/list", r.list)
	group.POST("/add", r.add)
	group.PUT("/update/:id", r.update)
	group.DELETE("/delete/:id", r.delete)
	group.GET("/info/:id", r.info)
}

func (r *Router) list(c echo.Context) error {
	var params model.Params
	var err error
	var response model.Response

	/* Validation parameter page */
	page := c.QueryParam("page")
	if params.Page, err = strconv.Atoi(page); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Invalid paramater page"
		return c.JSON(http.StatusOK, response)
	}

	/* Validation parameter limit */
	limit := c.QueryParam("limit")
	if params.Limit, err = strconv.Atoi(limit); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Invalid paramater limit"
		return c.JSON(http.StatusOK, response)
	}

	return nil
}

func (r *Router) add(c echo.Context) error {
	var newUser model.User
	var err error
	var response model.Response

	/* User payload validation */
	if err = c.Bind(&newUser); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Invalid payload"
		return c.JSON(http.StatusOK, response)
	}

	/* Add data process */
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

	/* User ID validation */
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "ID not valid"
		return c.JSON(http.StatusOK, response)
	}

	/* User payload validation */
	if err = c.Bind(&user); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Invalid payload"
		return c.JSON(http.StatusOK, response)
	}

	/* name, username, password validation */
	if len(user.Name) <= 5 || len(user.Username) <= 5 || len(user.Email) <= 5 || len(user.Password) <= 5 {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Invalid in name, username, email, or password"
		return c.JSON(http.StatusOK, response)
	}

	/* Update data process */
	updateResult := <-r.controller.Update(user, id)
	if updateResult.Error != nil {
		response.StatusCode = http.StatusUnprocessableEntity
		response.Message = updateResult.Error.Error()
		return c.JSON(http.StatusOK, response)
	}

	response.Message = "Succes update data"
	response.StatusCode = 200
	response.Data = updateResult.Data
	return c.JSON(http.StatusOK, response)
}

func (r *Router) info(c echo.Context) error {
	var id int
	var err error
	var response model.Response

	/* User ID validation */
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "ID not valid"
		return c.JSON(http.StatusOK, response)
	}

	/* Get data process */
	infoResult := <-r.controller.Info(id)
	if infoResult.Error != nil {
		response.StatusCode = http.StatusUnprocessableEntity
		response.Message = infoResult.Error.Error()
		return c.JSON(http.StatusOK, response)
	}

	response.StatusCode = 200
	response.Message = "Success get data"
	response.Data = infoResult.Data
	return c.JSON(http.StatusOK, response)
}

func (r *Router) delete(c echo.Context) error {
	var id int
	var err error
	var response model.Response

	/* User ID validation */
	if id, err = strconv.Atoi(c.Param("id")); err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "ID not valid"
		return c.JSON(http.StatusOK, response)
	}

	/* Delete data process */
	deleteResult := <-r.controller.Delete(id)
	if deleteResult.Error != nil {
		response.StatusCode = http.StatusUnprocessableEntity
		response.Message = deleteResult.Error.Error()
		return c.JSON(http.StatusOK, response)
	}

	response.StatusCode = 200
	response.Message = "Success delete data"
	return c.JSON(http.StatusOK, response)
}
