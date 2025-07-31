package routes

import (
	"task-manager/controllers"
	"task-manager/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	auth := api.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)

	api.Get("/profile", middleware.Protected(), controllers.Profile)
}

//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQyNDA3NjAsInVzZXJfaWQiOjF9.OHMkrOVUZJdb9jPgEAgB52ZJYTjP5xlmpz_EXcNflr4
