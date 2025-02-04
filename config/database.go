package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("book-manager.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB = db
	DB.AutoMigrate(&models.User{}, &models.Book{})
}
