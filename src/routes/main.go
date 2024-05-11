package routes

import (
	"belajar-fiber/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	app.Get("/products", controllers.GetAllProducts)
	app.Get("/products/:id", controllers.GetProductById)
	app.Post("/products", controllers.CreateProduct)
	app.Put("/products/:id", controllers.UpdateProduct)
	app.Delete("/products/:id", controllers.DeleteProduct)
}
