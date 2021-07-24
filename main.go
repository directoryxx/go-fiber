package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// if cant reach db try add "()"
	db, err := gorm.Open(mysql.Open("root:password@(mysql-server)/dev"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println(&db)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowMethods:  "GET,POST,PUT,DELETE,OPTIONS",
		ExposeHeaders: "Content-type,Authorization,accept",
		AllowOrigins:  "*",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	app.Get("/", home)

	log.Fatal(app.Listen(":8010"))
}

func home(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
