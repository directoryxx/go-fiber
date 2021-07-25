package routes

import (
	v2route "github.com/directoryxx/go-fiber/controller/v2"
	"github.com/gofiber/fiber/v2"
)

func V2(app *fiber.App) {
	api := app.Group("/api")
	v2 := api.Group("/v2")
	v2.Get("/", v2route.Home)
}
