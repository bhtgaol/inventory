package models

import "gorm.io/gorm"

// Toko adalah model dari toko
type Toko struct {
	gorm.Model
	User_Id string
	Name    string
	Street  string
	City    string
	State   string
	Country string
	Delete  int
}

// TokoBody adalah model yang akan menampung data dari body json
type TokoBody struct {
	User_Id string `json:"user_id"`
	Name    string `json:"name"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Delete  int
}
