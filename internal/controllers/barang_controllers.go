package controllers

import (
	"net/http"

	"github.com/bhtgaol/inventory/internal/initializers"
	"github.com/bhtgaol/inventory/internal/models"
	"github.com/gin-gonic/gin"
)

// CreateBarang adalah fungsi untuk menambahkan barang
func CreateBarang(c *gin.Context) {
	var body models.BarangBody
	var toko_id = c.Param("toko_id")

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	barang := models.Barang{
		Toko_Id:  toko_id,
		Name:     body.Name,
		Category: body.Category,
		Amount:   body.Amount,
	}

	result := initializers.DB.Create(&barang)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// SemuaBarang adalah fungsi yang memperlihatkan semua barang yang ada pada satu toko
func SemuaBarang(c *gin.Context) {
	var barang []models.Barang
	var toko_id = c.Param("toko_id")

	initializers.DB.Find(&barang, "Toko_ID = ?", toko_id)

	c.JSON(http.StatusOK, gin.H{
		"data": barang,
	})
}

// TemukanBarang adalah fungsi yang memperlihatkan satu barang dari satu toko
func TemukanBarang(c *gin.Context) {
	var barang models.Barang
	var barang_id = c.Param("barang_id")
	var toko_id = c.Param("toko_id")

	if err := initializers.DB.Find(&barang, "ID = ? AND Toko_ID = ?", barang_id, toko_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": barang,
	})
}

// UpdateToko adalah fungsi untuk melakukan perubahan data pada objek barang
func UpdateBarang(c *gin.Context) {
	var barang models.Barang
	var barang_id = c.Param("barang_id")
	var toko_id = c.Param("toko_id")

	if err := initializers.DB.Find(&barang, "ID = ? AND Toko_ID = ?", barang_id, toko_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found",
		})
		return
	}

	var body models.BarangBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	initializers.DB.Model(&barang).Updates(body)

	c.JSON(http.StatusOK, gin.H{
		"data": barang,
	})
}

// Deletetoko adalah fungsi untuk melakukan perubahan status delete barang
func DeleteBarang(c *gin.Context) {
	var barang models.Barang
	var barang_id = c.Param("barang_id")
	var toko_id = c.Param("toko_id")

	if err := initializers.DB.Find(&barang, "ID = ? AND Toko_ID = ?", barang_id, toko_id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found",
		})
		return
	}

	var body models.BarangBody
	body.Delete = 1

	initializers.DB.Model(&barang).Updates(body)

	c.JSON(http.StatusOK, gin.H{})
}
