package controller

import (
	"github.com/directoryxx/go-fiber/models"
	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Register(c *fiber.Ctx) error {
	var user models.User

	user.Id = 1
	user.Username = "test"
	user.Password = "test"

	return c.JSON(user)
}
