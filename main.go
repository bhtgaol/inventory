package main

import (
	"github.com/bhtgaol/inventory/internal/controllers"
	"github.com/bhtgaol/inventory/internal/initializers"
	"github.com/gin-gonic/gin"
)

// ini adalah fungsi yang akan di jalankan terlebih dahulu sebelum proses yang lain
func init() {
	initializers.LoadDotENV()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	//user route
	r.POST("/users/register", controllers.Register)
	r.POST("/users/login", controllers.Login)

	//toko route
	r.POST("/tokos", controllers.CreateToko)
	r.GET("/tokos", controllers.SemuaToko)
	r.GET("/tokos/toko_id", controllers.TemukanToko)
	r.PATCH("/tokos/toko_id", controllers.UpdateToko)
	r.DELETE("/tokos/toko_id", controllers.DeleteToko)

	//barang route

	r.Run()
}
