package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBCONFIG struct {
	HOST     string
	PORT     string
	USERNAME string
	PASSWORD string
	DATABASE string
}

var DB *gorm.DB
var err error

func ConnectDB() {
	dbConfig := DBCONFIG{
		HOST:     os.Getenv("DB_HOST"),
		PORT:     os.Getenv("DB_PORT"),
		USERNAME: os.Getenv("DB_USERNAME"),
		PASSWORD: os.Getenv("DB_PASSWORD"),
		DATABASE: os.Getenv("DB_DATABASE"),
	}
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", dbConfig.HOST, dbConfig.USERNAME, dbConfig.PASSWORD, dbConfig.DATABASE, dbConfig.PORT)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Database failed to connect")
	}
}
