package db

import (
	"fmt"
	"leiloa/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=leiloa port=5432 sslmode=disable",
		os.Getenv("HOST"), os.Getenv("USERDATABASE"), os.Getenv("DBPASSWORD"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connec to database: %v", err)
	}
}

func CreateDB() {
	DB.AutoMigrate(&models.User{}, &models.Auction{}, &models.Proposal{}, &models.Category{})
}
