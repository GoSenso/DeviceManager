package routes

import "github.com/gofiber/fiber/v2"

const (
	DeviceBasePath = "/device"
	RegisterDevice = "/register"
	DeleteDevice   = "/delete"
	GetDevice      = "/device"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	deviceApi := api.Group(DeviceBasePath)

	deviceApi.Post(RegisterDevice)
	deviceApi.Delete(DeleteDevice)
	deviceApi.Get(GetDevice)
	deviceApi.Get(GetDevice + "/:id")
}
