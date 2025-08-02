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
	auth.Post("/refresh", middleware.Protected(), controllers.Refresh)

	api.Get("/profile", middleware.Protected(), controllers.Profile)

	task := api.Group("/tasks", middleware.Protected())
	task.Post("/", controllers.CreateTask)
	task.Get("/", controllers.GetTasks)
	task.Get("/:id", controllers.GetTask)
	task.Put("/:id", controllers.UpdateTask)
	task.Delete("/:id", controllers.DeleteTask)
}
