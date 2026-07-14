package routes

import (
	"Proyecto_2B/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	userController *controllers.UserController,
	productController *controllers.ProductController,
	receiptController *controllers.ReceiptController,
) {
	api := router.Group("/api")

	users := api.Group("/users")
	{
		users.POST("/register", userController.Register)
		users.POST("/login", userController.Login)
		users.GET("/:id", userController.FindByID)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}

	products := api.Group("/products")
	{
		products.POST("", productController.Create)
		products.GET("", productController.FindAll)
		products.GET("/:id", productController.FindByID)
		products.PUT("/:id", productController.Update)
		products.DELETE("/:id", productController.Delete)
	}

	receipts := api.Group("/receipts")
	{
		receipts.POST("", receiptController.Create)
		receipts.GET("", receiptController.FindAll)
		receipts.GET("/:id", receiptController.FindByID)
		receipts.GET("/user/:userId", receiptController.FindByUserID)
		receipts.DELETE("/:id", receiptController.Delete)
	}
}
