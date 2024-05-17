package controllers

import (
	"belajar-fiber/src/helpers"
	"belajar-fiber/src/models"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func FindAllUsers(c *fiber.Ctx) error {
	users, err := models.FindAllUsers()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "There is no user",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   users,
	})
}

func FindUserById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user, err := models.FindUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "User is not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   user,
	})
}

type LoginResponse struct {
	ID        uint             `json:"id"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Role      string           `json:"role"`
	Addresses []models.Address `json:"addresses"`
	Token     string           `json:"token"`
}

func LoginWithEmailAndPassword(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	userFromDB, err := models.FindUserByEmail(user.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"code": fiber.StatusNotFound, "message": "Email is not found"})
	}
	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "Email or password is wrong",
		})
	}

	token, err := helpers.GenerateToken(os.Getenv("SECRET_KEY"), user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": fiber.StatusInternalServerError, "message": "Failed to generate token"})
	}

	loginResponse := LoginResponse{
		ID:        userFromDB.ID,
		CreatedAt: userFromDB.CreatedAt,
		UpdatedAt: userFromDB.UpdatedAt,
		Name:      userFromDB.Name,
		Role:      userFromDB.Role,
		Email:     userFromDB.Email,
		Addresses: userFromDB.Addresses,
		Token:     token,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   loginResponse,
	})
}

func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	errors := helpers.ValidateStruct(user)
	if len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	err = models.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"code":    fiber.StatusConflict,
			"message": fmt.Sprintf("email %v has already been used", user.Email),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":   fiber.StatusCreated,
		"status": "created",
		"data":   user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var user models.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}
	err = models.UpdateUser(id, &user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": fmt.Sprintf("Failed to update user with ID %d because there is no product with such id", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"status":  "ok",
		"message": fmt.Sprintf("User with id %d successfully updated.", id),
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := models.DeleteUser(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": fmt.Sprintf("Failed to delete user with ID %d because there is no user with such id", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"status":  "ok",
		"message": fmt.Sprintf("User with id %d successfully deleted.", id),
	})
}
