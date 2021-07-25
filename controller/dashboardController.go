package controller

import (
	"github.com/directoryxx/go-fiber/database"
	"github.com/directoryxx/go-fiber/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func DashboardHome(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	var userModels models.User

	database.DB.Where("id = ?", claims["id"]).First(&userModels)

	return c.JSON(fiber.Map{
		"user": userModels,
	})
}
