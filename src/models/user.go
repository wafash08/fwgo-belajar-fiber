package models

import (
	"belajar-fiber/src/configs"
	"time"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" gorm:"not null" validate:"required"`
	Email     string    `json:"email" gorm:"not null;unique" validate:"required,email"`
	Password  string    `json:"password" gorm:"not null" validate:"required"`
	Role      string    `json:"role" gorm:"not null"`
	Addresses []Address `json:"addresses"`
}

func FindAllUsers() ([]*User, error) {
	var users []*User
	err := configs.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func FindUserByID(id int) (*User, error) {
	var user User
	err := configs.DB.Take(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByEmailAndPassword(email string, password string) (*User, error) {
	var user User
	err := configs.DB.Take(&user, "email = ? and password = ?", email, password).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByEmail(email string) (*User, error) {
	var user User
	result := configs.DB.Where("email = ?", email).Take(&user)
	return &user, result.Error
}

func CreateUser(u *User) error {
	err := configs.DB.Create(&u).Error
	return err
}

func UpdateUser(id int, user *User) error {
	result := configs.DB.Model(&User{}).Where("id = ?", id).Updates(user)
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}

	return result.Error
}

func DeleteUser(id int) error {
	result := configs.DB.Delete(&User{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}
	return result.Error
}
