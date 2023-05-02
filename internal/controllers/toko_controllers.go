package controllers

import (
	"net/http"

	"github.com/bhtgaol/inventory/internal/initializers"
	"github.com/bhtgaol/inventory/internal/models"
	"github.com/gin-gonic/gin"
)

// CreateToko adalah fungsi untuk melakukan create objek toko
func CreateToko(c *gin.Context) {
	var body models.TokoBody
	// untuk user id, di ambil dari cookie, untuk sementara static
	var user_id = "1"

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	toko := models.Toko{
		User_Id: user_id,
		Name:    body.Name,
		Street:  body.Street,
		City:    body.City,
		State:   body.State,
		Country: body.Country,
	}

	result := initializers.DB.Create(&toko)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// SemuaToko adalah fungsi yang memperlihatkan semua toko dari user
func SemuaToko(c *gin.Context) {
	var toko []models.Toko
	// untuk user id, di ambil dari cookie, untuk sementara static
	var user_id = "1"

	initializers.DB.Find(&toko, "User_ID = ?", user_id)

	c.JSON(http.StatusOK, gin.H{
		"data": toko,
	})
}

// TemukanToko adalah fungsi yang memperlihatkan satu toko dari usser
func TemukanToko(c *gin.Context) {
	var toko models.Toko
	// untuk user id, di ambil dari cookie, untuk sementara static
	var user_id = "1"
	var toko_id = c.Param("toko_id")

	if err := initializers.DB.Find(&toko, "ID = ? AND User_ID = ?", toko_id, user_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": toko,
	})
}

// UpdateToko adalah fungsi untuk melakukan perubahan data pada objek toko
func UpdateToko(c *gin.Context) {
	var toko models.Toko
	// untuk user id, di ambil dari cookie, untuk sementara statik
	var user_id = "1"
	var toko_id = c.Param("toko_id")

	if err := initializers.DB.Find(&toko, "ID = ? AND User_ID = ?", toko_id, user_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found",
		})
		return
	}

	var body models.TokoBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	initializers.DB.Model(&toko).Updates(body)

	c.JSON(http.StatusOK, gin.H{
		"data": toko,
	})
}

// Deletetoko adalah fungsi untuk melakukan perubahan status delete toko
func DeleteToko(c *gin.Context) {
	var toko models.Toko
	// untuk user id, di ambil dari cookie, untuk sementara statik
	var user_id = "1"
	var toko_id = c.Param("book_id")

	if err := initializers.DB.Find(&toko, "ID = ? AND User_ID = ?", toko_id, user_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found",
		})
		return
	}
	var body models.TokoBody
	body.Delete = 1

	initializers.DB.Model(&toko).Updates(body)

	c.JSON(http.StatusOK, gin.H{})
}
