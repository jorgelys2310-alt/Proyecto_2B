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

// Create godoc
// @Summary Crear producto
// @Description Registra un producto nuevo en el catálogo.
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.ProductRequest true "Datos del producto"
// @Success 201 {object} dto.ProductResponse
// @Failure 400 {object} dto.APIError
// @Failure 401 {object} dto.APIError
// @Failure 500 {object} dto.APIError
// @Router /products [post]
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

// FindAll godoc
// @Summary Listar productos
// @Description Devuelve todos los productos registrados.
// @Tags Products
// @Produce json
// @Success 200 {array} dto.ProductResponse
// @Failure 500 {object} dto.APIError
// @Router /products [get]
func (controller *ProductController) FindAll(c *gin.Context) {
	response, err := controller.ProductService.FindAll()
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// FindByID godoc
// @Summary Buscar producto por ID
// @Description Devuelve un producto específico.
// @Tags Products
// @Produce json
// @Param id path int true "ID del producto"
// @Success 200 {object} dto.ProductResponse
// @Failure 400 {object} dto.APIError
// @Failure 404 {object} dto.APIError
// @Router /products/{id} [get]
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

// Update godoc
// @Summary Actualizar producto
// @Description Actualiza todos los datos de un producto.
// @Tags Products
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID del producto"
// @Param request body dto.ProductRequest true "Nuevos datos del producto"
// @Success 200 {object} dto.ProductResponse
// @Failure 400 {object} dto.APIError
// @Failure 401 {object} dto.APIError
// @Failure 404 {object} dto.APIError
// @Router /products/{id} [put]
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

// Delete godoc
// @Summary Eliminar producto
// @Description Elimina un producto del catálogo.
// @Tags Products
// @Security BearerAuth
// @Param id path int true "ID del producto"
// @Success 204 "Producto eliminado correctamente"
// @Failure 400 {object} dto.APIError
// @Failure 401 {object} dto.APIError
// @Failure 404 {object} dto.APIError
// @Router /products/{id} [delete]
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
