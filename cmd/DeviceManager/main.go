package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/GoSenso/DeviceManager/config"
	"github.com/GoSenso/DeviceManager/routes"
)

func main() {
	app := fiber.New()

	log.Default().Print(config.LoadConfig())
	routes.SetupRoutes(app)
	app.Listen(":3000")
}
