package helpers

import (
	"belajar-fiber/src/configs"
	"belajar-fiber/src/models"
)

func Migration() {
	configs.DB.AutoMigrate(&models.Category{}, &models.Product{}, &models.User{}, &models.Address{})
}
