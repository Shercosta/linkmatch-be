package database

import (
	"fmt"
	"linkmatch-be/helpers"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() *gorm.DB {
	dbhost := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	fmt.Println("Connecting to database", dbhost, dbport, dbuser, dbname)

	dsn := "host=" + dbhost + " user=" + dbuser + " password=" + dbpass + " dbname=" + dbname + " port=" + dbport + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	helpers.LogFatalCaption("Error connecting to database : ", err)
	return db
}
