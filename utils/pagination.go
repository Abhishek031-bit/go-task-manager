package utils

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := c.QueryInt("page", 1)
		pageSize := c.QueryInt("page_size", 10)

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
