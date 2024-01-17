package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This the root")
	})

	err := app.Listen("127.0.0.1:9991")
	if err != nil {
		fmt.Println(err)
		return
	}
}
