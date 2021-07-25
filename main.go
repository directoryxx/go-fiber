package main

import (
	"log"

	"github.com/directoryxx/go-fiber/database"
	"github.com/directoryxx/go-fiber/routes"
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

	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	// route v1
	routes.V1(app)
	// route v2
	routes.V2(app)

	log.Fatal(app.Listen(":8010"))
}
