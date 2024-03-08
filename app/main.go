package main

import (
	"app/app/database"
	"app/app/router"
	
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	router.SetupRoutes(app)

	app.Listen(":3000")
}

