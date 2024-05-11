package controllers

import (
	"belajar-fiber/src/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FindAllCategories(c *fiber.Ctx) error {
	categories, err := models.FindAllCategories()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Categories is not found",
		})
	}
	return c.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   categories,
	})
}

func FindCategoryByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	category, err := models.FindCategoryByID(id)
	fmt.Println("category >> ", category)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "Category is not found",
		})
	}
	return c.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   category,
	})
}

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	err := c.BodyParser(&category)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}
	err = models.CreateCategory(&category)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"code":   fiber.StatusCreated,
		"status": "created",
	})
}

func UpdateCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var category models.Category
	err := c.BodyParser(&category)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}
	err = models.UpdateCategory(id, &category)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": fmt.Sprintf("Failed to update product with ID %d", id),
		})
	}
	return c.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := models.DeleteCategory(id)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": fmt.Sprintf("Failed to delete product with ID %d", id),
		})
	}
	return c.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
	})
}
