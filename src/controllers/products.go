package controllers

import (
	"belajar-fiber/src/models"
	"fmt"
	"slices"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

var products = []models.Product{
	{ID: 1, Name: "Product A", Price: 10.99, Stock: 100},
	{ID: 2, Name: "Product B", Price: 20.50, Stock: 50},
	{ID: 3, Name: "Product C", Price: 15.75, Stock: 75},
}

func GetAllProducts(c *fiber.Ctx) error {
	return c.JSON(products)
}

func GetProductById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	foundIndex := slices.IndexFunc(products, func(p models.Product) bool {
		return p.ID == id
	})

	if foundIndex == -1 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("Product with ID %d is not found", id),
		})
	}

	foundProduct := products[foundIndex]

	return c.JSON(foundProduct)
}

func CreateProduct(c *fiber.Ctx) error {
	var newProduct models.Product
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
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var updatedProduct models.Product
	err := c.BodyParser(&updatedProduct)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
		return err
	}

	foundIndex := slices.IndexFunc(products, func(p models.Product) bool {
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
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	foundIndex := slices.IndexFunc(products, func(p models.Product) bool {
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
}
