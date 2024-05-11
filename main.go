package main

import (
	"belajar-fiber/src/configs"
	"belajar-fiber/src/helpers"
	"belajar-fiber/src/routes"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})
	configs.InitDB()
	helpers.Migration()
	routes.Router(app)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
