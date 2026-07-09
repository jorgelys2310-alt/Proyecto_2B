package main

import (
	"Proyecto_2B/config"
	"Proyecto_2B/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Product{})
	config.DB.AutoMigrate(&models.Receipt{})
	config.DB.AutoMigrate(&models.ReceiptItem{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Ecommerce con Go + Gin funcionando",
		})
	})

	r.Run(":8080")
}