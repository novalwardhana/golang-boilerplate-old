package main

import (
	"fmt"
	"os"

	globalENV "github.com/novalwardhana/golang-boilerplate/global/env"

	crudhandler "github.com/novalwardhana/golang-boilerplate/package/crud/handler"
	crudRepository "github.com/novalwardhana/golang-boilerplate/package/crud/repository"
	crudUsecase "github.com/novalwardhana/golang-boilerplate/package/crud/usecase"

	uploadFileHandler "github.com/novalwardhana/golang-boilerplate/package/upload-file/handler"
	uploadFileRepository "github.com/novalwardhana/golang-boilerplate/package/upload-file/repository"
	uploadFileUsecase "github.com/novalwardhana/golang-boilerplate/package/upload-file/usecase"

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

	/* Upload file function */
	uploadFileRepository := uploadFileRepository.NewRepository(dbMasterRead, dbMasterWrite)
	uploadFileUsecase := uploadFileUsecase.NewUsecase(uploadFileRepository)
	uploadFileHandler := uploadFileHandler.NewHandler(uploadFileUsecase)
	uploadFileGroup := r.Group("/api/v1/upload-file")
	uploadFileHandler.Mount(uploadFileGroup)

	r.Start(fmt.Sprintf(":%s", os.Getenv(globalENV.PORT)))
}
