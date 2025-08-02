package main

import (
	"log"
	"runtime"
	"task-manager/config"
	"task-manager/database"
	"task-manager/routes"
	_ "task-manager/workers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	config.LoadEnv()
	app := fiber.New()
	app.Use(logger.New())
	database.ConnectDB()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
