package models

import "time"

type Address struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Type        string    `json:"type" gorm:"not null" validate:"required"`
	PhoneNumber string    `json:"phone_number" gorm:"not null" validate:"required"`
	City        string    `json:"city" gorm:"not null" validate:"required"`
	PostalCode  string    `json:"postal_code" gorm:"not null" validate:"required"`
	Primary     bool      `json:"primary" gorm:"default:false"`
	UserID      uint      `json:"user_id" validate:"required"`
}
