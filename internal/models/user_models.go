package models

import "gorm.io/gorm"

// User adalah model dari user
type User struct {
	gorm.Model
	Name         string
	Phone_Number string
	Email        string `gorm:"unique"`
	Password     string
	Status       string
}

// UserBody adalah model yang akan menampung data user dari body
type UserBody struct {
	Name         string `json:"name"`
	Phone_Number string `json:"phone_number"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Status       string `json:"status"`
}
