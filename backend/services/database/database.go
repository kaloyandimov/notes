package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"notes/backend/models"
)

var database *gorm.DB

func Connect() {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Fatalln("godotenv: could not load .env file")
	}

	dsn := os.Getenv("DATABASE_DNS")
	database, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatalln("gorm: could not connect to database", err)
	}

	if err = database.AutoMigrate(&models.User{}, &models.Note{}, &models.Role{}); err != nil {
		log.Fatalln("gorm: could not run auto migration")
	}

	log.Println("Database connection successful")
}
