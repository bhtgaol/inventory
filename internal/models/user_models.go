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
