package database

import (
	"log"
	"medicine-inventory/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("medicine_inventory.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	DB = db
}

func AutoMigrate() {
	DB.AutoMigrate(&models.Medicine{}, &models.User{}, &models.Withdraw{})
}
