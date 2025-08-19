package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status" gorm:"default:'pending'"`
	DueDate     *time.Time `json:"due_date"` // Pointer to time.Time to allow null values
	UserID      uint      `json:"user_id" gorm:"not null"`
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}