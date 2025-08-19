package routes

import (
	"task-manager/controllers"
	"task-manager/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Post("/refresh", controllers.Refresh)
	app.Post("/logout", controllers.Logout)

	app.Get("/profile", middleware.Protected(), controllers.Profile)
	tasks := app.Group("/tasks")
	tasks.Use(logger.New(), middleware.Protected())
	tasks.Get("/", controllers.GetTasks)
	tasks.Post("/", controllers.CreateTask)
	tasks.Get("/:id", controllers.GetTask)
	tasks.Put("/:id", controllers.UpdateTask)
	tasks.Delete("/:id", controllers.DeleteTask)
}
