package models

import "gorm.io/gorm"

// Barang adalah model dari barang
type Barang struct {
	gorm.Model
	Toko_Id  string
	Name     string
	Category string
	Amount   string
	Delete   int
}

// BarangBody adalah model yang akan menampung data dari body json
type BarangBody struct {
	Toko_Id  string `json:"toko_id"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Amount   string `json:"amount"`
	Delete   int
}
