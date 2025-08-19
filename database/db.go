package database

import (
	"fmt"
	"log"
	"task-manager/config"
	"task-manager/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open(config.DATABASE_URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("❌ Failed to connect to database")
	}
	fmt.Println("✅ Connected to database")
	DB.AutoMigrate(&models.User{}, &models.Task{})
	fmt.Println("Database migrated")
}
