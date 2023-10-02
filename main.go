package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	const PORT string = ":8080"
	fmt.Println("Booting up gotodo api server...")
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
