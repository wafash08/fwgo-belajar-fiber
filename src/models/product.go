package models

import (
	"belajar-fiber/src/configs"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Product struct {
	ID          int            `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Name        string         `json:"name" gorm:"not null" validate:"required,min=3,max=100"`
	Brand       string         `json:"brand" gorm:"not null"`
	Rating      int            `json:"rating" gorm:"default:0"`
	Price       float64        `json:"price" gorm:"not null" validate:"required,min=0"`
	Color       string         `json:"color" gorm:"not null"`
	Size        int            `json:"size" gorm:"not null"`
	Quantity    int            `json:"quantity" validate:"required,min=0"`
	Image       string         `json:"image" gorm:"not null"`
	Condition   string         `json:"condition" gorm:"not null"`
	Description string         `json:"description"`
	CategoryID  uint           `json:"category_id"`
	Category    Category       `json:"category" gorm:"foreignKey:CategoryID"`
}

func FindAllProducts() ([]*Product, error) {
	var products []*Product
	err := configs.DB.Preload("Category").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func FindProductByID(id int) (*Product, error) {
	var product Product
	err := configs.DB.Preload("Category").Take(&product, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func CreateProduct(p *Product) error {
	err := configs.DB.Create(&p).Error
	return err
}

func UpdateProduct(id int, product *Product) error {
	result := configs.DB.Model(&Product{}).Where("id = ?", id).Updates(product)
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}

	return result.Error
}

func DeleteProduct(id int) error {
	result := configs.DB.Delete(&Product{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}
	return result.Error
}
