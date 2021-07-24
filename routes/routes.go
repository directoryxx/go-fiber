package routes

import (
	"github.com/directoryxx/go-fiber/controller"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/", controller.Home)

	auth := v1.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)
}
