package routes

import (
	"belajar-fiber/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	api := app.Group("/api/v1")
	categories := api.Group("/categories")
	categories.Get("/", controllers.FindAllCategories)
	categories.Get("/:id", controllers.FindCategoryByID)
	categories.Post("/", controllers.CreateCategory)
	categories.Put("/:id", controllers.UpdateCategory)
	categories.Delete("/:id", controllers.DeleteCategory)
}
