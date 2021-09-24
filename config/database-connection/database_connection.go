package databaseConnection

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	globalENV "github.com/novalwardhana/golang-boilerplate/global/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBMaster() *gorm.DB {
	uri := os.Getenv(globalENV.DBMaster)
	db := CreateConnection(uri)
	return db
}

func CreateConnection(uri string) *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: newLogger,
	})
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
