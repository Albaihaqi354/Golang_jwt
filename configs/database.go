package configs

import (
	"golang_jwt/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	dsn := "host=localhost user=postgres password=Whobay123@ dbname=go_jwt port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed Connected To Database")
	}

	db.AutoMigrate(&models.User{})

	DB = db
	log.Println("Succesfully Connected To Database")
}
