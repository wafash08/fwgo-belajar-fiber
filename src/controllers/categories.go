package controllers

import (
	"belajar-fiber/src/helpers"
	"belajar-fiber/src/models"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FindAllCategories(c *fiber.Ctx) error {
	categories, err := models.FindAllCategories()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "Categories is not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   categories,
	})
}

func FindCategoryByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	category, err := models.FindCategoryByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "Category is not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   category,
	})
}

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	err := c.BodyParser(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}
	errors := helpers.ValidateStruct(category)
	if len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}
	err = models.CreateCategory(&category)
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

func UpdateCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var category models.Category
	err := c.BodyParser(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}
	// errors := helpers.ValidateStruct(category)
	// if len(errors) > 0 {
	// 	return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	// }
	err = models.UpdateCategory(id, &category)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": fmt.Sprintf("Failed to update category with ID %d because there is no category with such id", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := models.DeleteCategory(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": fmt.Sprintf("Failed to delete category with ID %d because there is no category with such id", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
	})
}
