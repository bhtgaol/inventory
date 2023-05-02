package models

import "gorm.io/gorm"

// Barang adalah model dari barang
type Barang struct {
	gorm.Model
	User_Id string `json:"user_id"`
	Name    string `json:"name"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
	Delete  int
}

// BarangBody adalah model yang akan menampung data dari body json
type BarangBody struct {
	Name    string
	Street  string
	City    string
	State   string
	Country string
	Delete  int
}
