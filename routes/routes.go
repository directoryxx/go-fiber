package routes

import (
	"github.com/directoryxx/go-fiber/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controller.Home)
}
