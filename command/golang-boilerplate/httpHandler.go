package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	globalENV "github.com/novalwardhana/golang-boiler-plate/global/env"
	"gorm.io/gorm"
)

func StartHTTPHandler(dbMasterRead *gorm.DB, dbMasterWrite *gorm.DB) {
	r := echo.New()
	r.Start(fmt.Sprintf(":%s", os.Getenv(globalENV.PORT)))
}
