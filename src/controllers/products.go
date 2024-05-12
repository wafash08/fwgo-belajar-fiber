package controllers

import (
	"belajar-fiber/src/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FindAllProducts(c *fiber.Ctx) error {
	products, err := models.FindAllProducts()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "There is no product.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   products,
	})
}

func FindProductById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product, err := models.FindProductByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "Product is not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   product,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	err = models.CreateProduct(&product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":   fiber.StatusCreated,
		"status": "created",
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}
	err = models.UpdateProduct(id, &product)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": fmt.Sprintf("Failed to update product with ID %d because there is no product with such id", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"status":  "ok",
		"message": fmt.Sprintf("Product with id %d successfully updated.", id),
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := models.DeleteProduct(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": fmt.Sprintf("Failed to delete product with ID %d because there is no product with such id", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"status":  "ok",
		"message": fmt.Sprintf("Product with id %d successfully deleted.", id),
	})
}
