package main

import (
	"Proyecto_2B/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Ecommerce con Go + Gin funcionando",
		})
	})

	r.Run(":8080")
}