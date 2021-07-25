package controller

import (
	"strconv"
	"time"

	"github.com/directoryxx/go-fiber/database"
	"github.com/directoryxx/go-fiber/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("username = ?", data["username"]).First(&user)

	if user.ID == 0 {
		c.Status(401)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "User not found",
		})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])) != nil {
		c.Status(401)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Incorrect password",
		})
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = strconv.Itoa(int(user.ID))
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Login successful",
		"token":   t,
		// "user": user,
	})
}
