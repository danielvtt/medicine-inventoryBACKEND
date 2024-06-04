package database

import (
	"log"
	"medicine-inventory/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("medicine_inventory.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
}

func AutoMigrate() {
	DB.AutoMigrate(&models.Medicine{}, &models.User{}, &models.Withdraw{})
}
