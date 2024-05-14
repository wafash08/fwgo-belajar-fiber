package models

import (
	"belajar-fiber/src/configs"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name" gorm:"not null" validate:"required,min=3,max=50"`
	Color     string         `json:"color" gorm:"not null"`
	Image     string         `json:"image" gorm:"not null"`
	Products  []Product      `json:"products"`
}

func FindAllCategories() ([]*Category, error) {
	var categories []*Category
	err := configs.DB.Preload("Products").Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func FindCategoryByID(id int) (*Category, error) {
	var category Category
	err := configs.DB.Preload("Products").Take(&category, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func CreateCategory(c *Category) error {
	err := configs.DB.Create(&c).Error
	return err
}

func UpdateCategory(id int, category *Category) error {
	result := configs.DB.Model(&Category{}).Where("id = ?", id).Updates(category)
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}

	return result.Error
}

func DeleteCategory(id int) error {
	result := configs.DB.Delete(&Category{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}
	return result.Error
}
