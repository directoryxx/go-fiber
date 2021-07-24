package main

import (
	"log"

	"github.com/directoryxx/go-fiber/routes"
	"github.com/directoryxx/go-fiber/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowMethods:  "GET,POST,PUT,DELETE,OPTIONS",
		ExposeHeaders: "Content-type,Authorization,accept",
		AllowOrigins:  "*",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	routes.Setup(app)

	log.Fatal(app.Listen(":8010"))
}
