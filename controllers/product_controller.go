package controllers

import (
	"net/http"

	"Proyecto_2B/dto"
	"Proyecto_2B/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *services.ProductService
}

func NewProductController(
	productService *services.ProductService,
) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (controller *ProductController) Create(c *gin.Context) {
	var request dto.ProductRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		respondValidationError(c, err)
		return
	}

	response, err := controller.ProductService.Create(request)
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (controller *ProductController) FindAll(c *gin.Context) {
	response, err := controller.ProductService.FindAll()
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (controller *ProductController) FindByID(c *gin.Context) {
	id, valid := parseID(c, "id")
	if !valid {
		return
	}

	response, err := controller.ProductService.FindByID(id)
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (controller *ProductController) Update(c *gin.Context) {
	id, valid := parseID(c, "id")
	if !valid {
		return
	}

	var request dto.ProductRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		respondValidationError(c, err)
		return
	}

	response, err := controller.ProductService.Update(id, request)
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (controller *ProductController) Delete(c *gin.Context) {
	id, valid := parseID(c, "id")
	if !valid {
		return
	}

	if err := controller.ProductService.Delete(id); err != nil {
		respondError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
