package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rajakumardev/gotodo/initializers"
)

func main() {

	// Load configuration
	config, err := initializers.LoadConfig(".")

	if err != nil {
		log.Fatal("Initialization Failed while reading configuration file...")
	}

	PORT := config.PORT
	fmt.Printf("Booting up gotodo api server on PORT %s...", PORT)
	app := fiber.New()
	app.Get("/ping", func(c *fiber.Ctx) error {

		response := map[string]string{
			"method": "ping",
			"status": "200",
		}
		return c.JSON(response)
	})
	app.Listen(PORT)
}
