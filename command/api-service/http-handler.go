package main

import (
	"fmt"
	"os"

	globalENV "github.com/novalwardhana/golang-boilerplate/global/env"

	crudhandler "github.com/novalwardhana/golang-boilerplate/package/crud/handler"
	crudRepository "github.com/novalwardhana/golang-boilerplate/package/crud/repository"
	crudUsecase "github.com/novalwardhana/golang-boilerplate/package/crud/usecase"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func StartHTTPHandler(dbMasterRead *gorm.DB, dbMasterWrite *gorm.DB) {
	r := echo.New()

	/* crud basic function */
	crudRepository := crudRepository.NewRepository(dbMasterRead, dbMasterWrite)
	crudUsecase := crudUsecase.NewUsecase(crudRepository)
	crudhandler := crudhandler.NewHandler(crudUsecase)
	crudGroup := r.Group("/api/v1/crud")
	crudhandler.Mount(crudGroup)

	r.Start(fmt.Sprintf(":%s", os.Getenv(globalENV.PORT)))
}
