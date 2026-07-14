package routes

import (
	"Proyecto_2B/controllers"
	"Proyecto_2B/middleware"

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
		// Rutas públicas
		users.POST("/register", userController.Register)
		users.POST("/login", userController.Login)

		// Rutas protegidas
		protectedUsers := users.Group("")
		protectedUsers.Use(middleware.AuthMiddleware())
		{
			protectedUsers.GET("/:id", userController.FindByID)
			protectedUsers.PUT("/:id", userController.Update)
			protectedUsers.DELETE("/:id", userController.Delete)
		}
	}

	products := api.Group("/products")
	{
		// Consultas públicas
		products.GET("", productController.FindAll)
		products.GET("/:id", productController.FindByID)

		// Operaciones protegidas
		protectedProducts := products.Group("")
		protectedProducts.Use(middleware.AuthMiddleware())
		{
			protectedProducts.POST("", productController.Create)
			protectedProducts.PUT("/:id", productController.Update)
			protectedProducts.DELETE("/:id", productController.Delete)
		}
	}

	receipts := api.Group("/receipts")
	receipts.Use(middleware.AuthMiddleware())
	{
		receipts.POST("", receiptController.Create)
		receipts.GET("", receiptController.FindAll)
		receipts.GET("/user/:userId", receiptController.FindByUserID)
		receipts.GET("/:id", receiptController.FindByID)
		receipts.DELETE("/:id", receiptController.Delete)
	}
}
