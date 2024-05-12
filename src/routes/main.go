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

	products := api.Group("/products")
	products.Get("/", controllers.FindAllProducts)
	products.Get("/:id", controllers.FindProductById)
	products.Post("/", controllers.CreateProduct)
	products.Put("/:id", controllers.UpdateProduct)
	products.Delete("/:id", controllers.DeleteProduct)

	users := api.Group("/users")
	users.Get("/", controllers.FindAllUsers)
	users.Get("/:id", controllers.FindUserById)
	users.Post("/register", controllers.RegisterUser)
	users.Post("/login", controllers.LoginWithEmailAndPassword)
	users.Put("/:id", controllers.UpdateUser)
	users.Delete("/:id", controllers.DeleteUser)
}
