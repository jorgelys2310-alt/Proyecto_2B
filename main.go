package main

import (
	"log"

	"Proyecto_2B/config"
	"Proyecto_2B/controllers"
	"Proyecto_2B/models"
	"Proyecto_2B/repository"
	"Proyecto_2B/routes"
	"Proyecto_2B/services"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Receipt{},
		&models.ReceiptItem{},
	)

	if err != nil {
		log.Fatal("Error al ejecutar las migraciones: ", err)
	}

	err = config.DB.Exec(`
		ALTER TABLE receipts
		ADD CONSTRAINT fk_receipts_user
		FOREIGN KEY (created_by)
		REFERENCES users(user_id)
		ON UPDATE CASCADE
		ON DELETE RESTRICT
	`).Error

	if err != nil {
		log.Println("Relación receipts-users:", err)
	}

	err = config.DB.Exec(`
		ALTER TABLE receipt_items
		ADD CONSTRAINT fk_receipt_items_receipt
		FOREIGN KEY (receipt_id)
		REFERENCES receipts(receipt_id)
		ON UPDATE CASCADE
		ON DELETE CASCADE
	`).Error

	if err != nil {
		log.Println("Relación receipt_items-receipts:", err)
	}

	err = config.DB.Exec(`
		ALTER TABLE receipt_items
		ADD CONSTRAINT fk_receipt_items_product
		FOREIGN KEY (product_id)
		REFERENCES products(product_id)
		ON UPDATE CASCADE
		ON DELETE RESTRICT
	`).Error

	if err != nil {
		log.Println("Relación receipt_items-products:", err)
	}

	userRepository := repository.NewUserRepository(config.DB)
	productRepository := repository.NewProductRepository(config.DB)
	receiptRepository := repository.NewReceiptRepository(config.DB)

	userService := services.NewUserService(userRepository)
	productService := services.NewProductService(productRepository)

	receiptService := services.NewReceiptService(
		config.DB,
		receiptRepository,
		userRepository,
		productRepository,
	)

	userController := controllers.NewUserController(userService)
	productController := controllers.NewProductController(productService)
	receiptController := controllers.NewReceiptController(receiptService)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API Ecommerce con Go + Gin funcionando",
		})
	})

	routes.SetupRoutes(
		router,
		userController,
		productController,
		receiptController,
	)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
