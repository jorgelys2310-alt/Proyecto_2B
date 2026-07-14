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

func (controller *ReceiptController) FindAll(c *gin.Context) {
	response, err := controller.ReceiptService.FindAll()
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

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
