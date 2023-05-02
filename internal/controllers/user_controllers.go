package controllers

import (
	"net/http"

	"github.com/bhtgaol/inventory/internal/initializers"
	"github.com/bhtgaol/inventory/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register adalah fungsi untuk registrasi user baru
func Register(c *gin.Context) {
	var body models.UserBody

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.MinCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed hash password",
		})
	}

	user := models.User{
		Name:         body.Name,
		Phone_Number: body.Phone_Number,
		Email:        body.Email,
		Password:     string(hash),
		Status:       body.Status,
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// Login adalah fungsi untuk masuk kedalam aplikasi
func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	// token not implemented yet

	c.SetSameSite(http.SameSiteLaxMode)

	// cookie not implemented yet c.SetCookie()

	c.JSON(http.StatusOK, gin.H{})
}
