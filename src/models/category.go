package models

import (
	"belajar-fiber/src/configs"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name" gorm:"not null"`
	Color     string         `json:"color" gorm:"not null"`
	Image     string         `json:"image" gorm:"not null"`
}

func FindAllCategories() ([]*Category, error) {
	var categories []*Category
	err := configs.DB.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func FindCategoryByID(id int) (*Category, error) {
	var category Category
	err := configs.DB.Take(&category, "id = ?", id).Error
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
	err := configs.DB.Model(&Category{}).Where("id = ?", id).Updates(category).Error
	return err
}

func DeleteCategory(id int) error {
	err := configs.DB.Delete(&Category{}, "id = ?", id).Error
	return err
}
