package main

import (
	"github.com/gofiber/fiber/v2"
	configs "github.com/harshalranjhani/gofibre/config"
	routes "github.com/harshalranjhani/gofibre/routes"
)

func main() {
	// create an instance of fiber
	app := fiber.New()

	configs.ConnectDB()

	routes.UserRoute(app)

	// listen on port
	app.Listen(":8080")

}
