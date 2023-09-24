package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func initServer() *fiber.App {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		time.Sleep(time.Millisecond * 10)
		return c.SendString("hello world")
	})
	return app
}
func main() {
	app := initServer()

	app.Listen(":8080")
}
