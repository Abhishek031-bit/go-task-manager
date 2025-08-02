package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status" gorm:"default:'pending'"`
	UserID      uint   `json:"user_id"`
	User        User   `gorm:"foreignKey:UserID"`
}