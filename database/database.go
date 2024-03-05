package database

import (
	"campus-api/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Couldn't connect to the database")
	} else {
		log.Println("Connect to DB Succesfully")
	}
	DB = database
	database.AutoMigrate(
		&models.EmailData{},
		&models.User{},
		&models.Admin{},
	)
}
