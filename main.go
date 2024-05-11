package main

import (
	"belajar-fiber/src/routes"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})

	routes.Router(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
