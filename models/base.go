package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
	"os"
	"github.com/joho/godotenv"
	"fmt"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost :=os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbUri := fmt.Sprint("host=%s port=%s dbname=%s sslmode=disable user=%s password=%s", dbHost, dbPort, dbName, username, password)

	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)

	if err != nil {
		fmt.Print(err)
	}

	db = conn
	//db.Debug().AutoMigrate(...)
}

func GetDB() *gorm.DB {
	return db
}