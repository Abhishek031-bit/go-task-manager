package database

import (
	"task-manager/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("task.db"), &gorm.Config{})
	if err != nil {
		panic("❌Failed to connect database")
	}
	println("✅Connected to database")
	db.AutoMigrate(&models.User{})
	DB = db
}
