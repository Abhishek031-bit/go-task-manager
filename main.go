package main

import (
	"log"
	"runtime"
	"task-manager/config"
	"task-manager/database"
	"task-manager/routes"
	_ "task-manager/workers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	config.LoadEnv()
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	database.ConnectDB()
	routes.SetupRoutes(app)
	port := "3000"
	log.Fatal(app.Listen(":" + port))
}
