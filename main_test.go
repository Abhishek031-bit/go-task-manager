package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"task-manager/controllers"
	"task-manager/database"
	"task-manager/models"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	database.ConnectDB()
	app := fiber.New()
	app.Post("/api/auth/register", controllers.Register)

	// Clean up the database after the test
	defer database.DB.Migrator().DropTable(&models.User{})

	// Create a new user
	user := controllers.RegisterInput{
		Name:     "test",
		Email:    "test@example.com",
		Password: "password",
	}
	payload, _ := json.Marshal(user)

	req := httptest.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var result map[string]any
	json.NewDecoder(resp.Body).Decode(&result)

	assert.Equal(t, "User created successfully", result["message"])
	assert.NotNil(t, result["token"])
}