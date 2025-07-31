package main

import (
	"log"
	"task-manager/config"
	"task-manager/database"
	"task-manager/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config.LoadEnv()
	app := fiber.New()
	app.Use(logger.New())
	database.ConnectDB()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
