package main

import (
	"fmt"
	"os"

	globalENV "github.com/novalwardhana/golang-boiler-plate/global/env"

	crudConteroller "github.com/novalwardhana/golang-boiler-plate/package/crud/controller"
	crudRouter "github.com/novalwardhana/golang-boiler-plate/package/crud/router"
	crudService "github.com/novalwardhana/golang-boiler-plate/package/crud/service"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func StartHTTPHandler(dbMasterRead *gorm.DB, dbMasterWrite *gorm.DB) {
	r := echo.New()

	/* crud basic function */
	crudService := crudService.NewService(dbMasterRead, dbMasterWrite)
	crudConteroller := crudConteroller.NewController(crudService)
	crudRouter := crudRouter.NewRouter(crudConteroller)
	crudGroup := r.Group("/api/v1/crud")
	crudRouter.Mount(crudGroup)

	r.Start(fmt.Sprintf(":%s", os.Getenv(globalENV.PORT)))
}
