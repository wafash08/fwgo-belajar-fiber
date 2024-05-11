package main

import (
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

func (p *Product) Init() []Product {
	return []Product{
		{1, "Product A", 10.99, 100},
		{2, "Product B", 20.50, 50},
		{3, "Product C", 15.75, 75},
	}
}

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	var product Product
	products := product.Init()

	app.Get("/products", func(c *fiber.Ctx) error {
		return c.JSON(products)
	})

	app.Get("/products/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))

		var foundProduct Product
		for _, p := range products {
			if p.ID == id {
				foundProduct = p
				break
			}
		}

		return c.JSON(foundProduct)
	})

	app.Post("/products", func(c *fiber.Ctx) error {
		var newProduct Product
		err := c.BodyParser(&newProduct)
		if err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
			return err
		}

		newProduct.ID = len(products) + 1
		products = append(products, newProduct)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "product created successfully",
			"product": newProduct,
		})
	})

	app.Put("/products/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))

		var updatedProduct Product
		err := c.BodyParser(&updatedProduct)
		if err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid request body",
			})
			return err
		}

		foundIndex := slices.IndexFunc(products, func(p Product) bool {
			return p.ID == id
		})

		if foundIndex == -1 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("Product with ID %d is not found", id),
			})
		}

		updatedProduct.ID = id
		products[foundIndex] = updatedProduct
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("Product with ID %d updated successfully", id),
			"product": updatedProduct,
		})
	})

	app.Delete("/products/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))

		foundIndex := slices.IndexFunc(products, func(p Product) bool {
			return p.ID == id
		})

		if foundIndex == -1 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": fmt.Sprintf("Product with ID %d is not found", id),
			})
		}

		products = append(products[:foundIndex], products[foundIndex+1:]...)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("Product with ID %d deleted successfully", id),
		})
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
