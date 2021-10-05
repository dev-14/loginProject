package main

import (
	"loginProject/database"
	"loginProject/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	//"gorm.io/gorm"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Route(app)

	app.Listen(":3000")

}
