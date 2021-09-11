package main

import (
	"os"

	config "github.com/joho/godotenv"
	dbConn "github.com/novalwardhana/golang-boiler-plate/config/database-connection"
	"gorm.io/gorm"
)

var dbMasterRead *gorm.DB
var dbMasterWrite *gorm.DB

func main() {

	if os.Getenv("GO_ENV") == "local" {
		if err := config.Load(".env"); err != nil {
			os.Exit(1)
		}
	}

	dbMasterRead = dbConn.DBMaster()
	dbMasterWrite = dbConn.DBMaster()

	forever := make(chan int)
	go func() {
		go StartHTTPHandler(dbMasterRead, dbMasterWrite)
	}()
	<-forever

}
