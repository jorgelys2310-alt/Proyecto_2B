package controllers

import (
	"net/http"

	"Proyecto_2B/dto"
	"Proyecto_2B/services"

	"github.com/gin-gonic/gin"
)

type ReceiptController struct {
	ReceiptService *services.ReceiptService
}

func NewReceiptController(
	receiptService *services.ReceiptService,
) *ReceiptController {
	return &ReceiptController{
		ReceiptService: receiptService,
	}
}

// Create godoc
// @Summary Crear recibo
// @Description Crea una compra, calcula el total y descuenta automáticamente el stock.
// @Tags Receipts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateReceiptRequest true "Usuario y productos comprados"
// @Success 201 {object} dto.ReceiptResponse
// @Failure 400 {object} dto.APIError
// @Failure 401 {object} dto.APIError
// @Failure 404 {object} dto.APIError
// @Failure 500 {object} dto.APIError
// @Router /receipts [post]
func (controller *ReceiptController) Create(c *gin.Context) {
	var request dto.CreateReceiptRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		respondValidationError(c, err)
		return
	}

	response, err := controller.ReceiptService.Create(request)
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

// FindAll godoc
// @Summary Listar recibos
// @Description Devuelve todos los recibos registrados.
// @Tags Receipts
// @Produce json
// @Security BearerAuth
// @Success 200 {array} dto.ReceiptResponse
// @Failure 401 {object} dto.APIError
// @Failure 500 {object} dto.APIError
// @Router /receipts [get]
func (controller *ReceiptController) FindAll(c *gin.Context) {
	response, err := controller.ReceiptService.FindAll()
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// FindByID godoc
// @Summary Buscar recibo por ID
// @Description Devuelve un recibo con sus productos y subtotales.
// @Tags Receipts
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID del recibo"
// @Success 200 {object} dto.ReceiptResponse
// @Failure 400 {object} dto.APIError
// @Failure 401 {object} dto.APIError
// @Failure 404 {object} dto.APIError
// @Router /receipts/{id} [get]
func (controller *ReceiptController) FindByID(c *gin.Context) {
	id, valid := parseID(c, "id")
	if !valid {
		return
	}

	response, err := controller.ReceiptService.FindByID(id)
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// FindByUserID godoc
// @Summary Listar recibos de un usuario
// @Description Devuelve todos los recibos pertenecientes a un usuario.
// @Tags Receipts
// @Produce json
// @Security BearerAuth
// @Param userId path int true "ID del usuario"
// @Success 200 {array} dto.ReceiptResponse
// @Failure 400 {object} dto.APIError
// @Failure 401 {object} dto.APIError
// @Failure 404 {object} dto.APIError
// @Router /receipts/user/{userId} [get]
func (controller *ReceiptController) FindByUserID(c *gin.Context) {
	userID, valid := parseID(c, "userId")
	if !valid {
		return
	}

	response, err := controller.ReceiptService.FindByUserID(userID)
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// Delete godoc
// @Summary Eliminar recibo
// @Description Elimina un recibo y sus detalles.
// @Tags Receipts
// @Security BearerAuth
// @Param id path int true "ID del recibo"
// @Success 204 "Recibo eliminado correctamente"
// @Failure 400 {object} dto.APIError
// @Failure 401 {object} dto.APIError
// @Failure 404 {object} dto.APIError
// @Router /receipts/{id} [delete]
func (controller *ReceiptController) Delete(c *gin.Context) {
	id, valid := parseID(c, "id")
	if !valid {
		return
	}

	if err := controller.ReceiptService.Delete(id); err != nil {
		respondError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
