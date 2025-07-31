package controllers

import (
	"task-manager/database"
	"task-manager/models"
	"task-manager/utils"

	"github.com/gofiber/fiber/v2"
)

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var input RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"err": "Invalid input"})
	}
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"err": "Failed to hash password"})
	}
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusConflict).JSON(&fiber.Map{"err": "Email already exists"})
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"err": "Failed to generate token"})
	}
	return c.JSON(&fiber.Map{"message": "User created successfully", "token": token})
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"err": "Invalid input"})
	}
	var user models.User
	result := database.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{"err": "User not found"})
	}
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{"err": "Invalid password"})
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"err": "Failed to generate token"})
	}
	return c.JSON(&fiber.Map{"message": "User logged in successfully", "token": token})
}

func Profile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{"err": "User not found"})
	}
	return c.JSON(&fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}
