package main

import (
	"golang-auth/database"

	"golang-auth/routes"

	// FRAMEWORK WEB APP, SAMA SEPERTI EXPRESSJS DI NODE JS
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// CONNECT TO DATABASE
	database.Connect()

	// CREATE INITIAL FIBER
	app := fiber.New()

	// ALLOW CORS METHOD
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// IMPORT ALL ROUTES
	routes.Setup(app)

	app.Listen(":4000")
}
