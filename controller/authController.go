package controller

import (
	"fmt"

	"github.com/directoryxx/go-fiber/database"
	"github.com/directoryxx/go-fiber/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Username: data["username"],
		Password: password,
	}

	database.DB.Create(&user)

	fmt.Println(user)

	return c.JSON(user)
}
