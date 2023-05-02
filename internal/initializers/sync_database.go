package initializers

import "github.com/bhtgaol/inventory/internal/models"

// SyncDatabase adalah fungsi yang akan melakukan auto migrate
func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
