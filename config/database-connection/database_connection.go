package databaseConnection

import (
	"fmt"
	"os"
	"strconv"

	globalENV "github.com/novalwardhana/golang-boiler-plate/global/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBMaster() *gorm.DB {
	uri := os.Getenv(globalENV.DBMaster)
	db := CreateConnection(uri)
	return db
}

func CreateConnection(uri string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		fmt.Println("Please check database connection", err)
	}

	sql, err := db.DB()
	if err != nil {
		fmt.Println("Please check database connection")
	}

	maxPool, err := strconv.Atoi(os.Getenv(globalENV.DBMaxConnectionPool))
	if err != nil {
		maxPool = 5
	}
	sql.SetMaxOpenConns(maxPool)

	maxIdle, err := strconv.Atoi(os.Getenv(globalENV.DBMaxConnectionIdle))
	if err != nil {
		maxIdle = 2
	}
	sql.SetMaxIdleConns(maxIdle)

	return db
}
