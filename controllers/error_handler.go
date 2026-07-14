package controllers

import (
	"errors"
	"net/http"
	"time"

	"Proyecto_2B/dto"
	"Proyecto_2B/exceptions"

	"github.com/gin-gonic/gin"
)

func respondError(c *gin.Context, err error) {
	status := http.StatusInternalServerError
	errorName := "Internal Server Error"
	message := "Ocurrió un error interno en el servidor"

	switch {
	case errors.Is(err, exceptions.ErrNotFound):
		status = http.StatusNotFound
		errorName = "Not Found"
		message = err.Error()

	case errors.Is(err, exceptions.ErrBadRequest):
		status = http.StatusBadRequest
		errorName = "Bad Request"
		message = err.Error()

	case errors.Is(err, exceptions.ErrEmailAlreadyExists):
		status = http.StatusBadRequest
		errorName = "Bad Request"
		message = err.Error()

	case errors.Is(err, exceptions.ErrInvalidCredentials):
		status = http.StatusUnauthorized
		errorName = "Unauthorized"
		message = err.Error()

	case errors.Is(err, exceptions.ErrInsufficientStock):
		status = http.StatusBadRequest
		errorName = "Bad Request"
		message = err.Error()
	}

	c.JSON(status, dto.APIError{
		Timestamp: time.Now(),
		Status:    status,
		Error:     errorName,
		Message:   message,
		Path:      c.Request.URL.Path,
	})
}

func respondValidationError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, dto.APIError{
		Timestamp: time.Now(),
		Status:    http.StatusBadRequest,
		Error:     "Validation Error",
		Message:   err.Error(),
		Path:      c.Request.URL.Path,
	})
}

func parseID(c *gin.Context, parameter string) (uint, bool) {
	value := c.Param(parameter)

	var id uint64

	for _, character := range value {
		if character < '0' || character > '9' {
			respondValidationError(c, errors.New("el identificador debe ser numérico"))
			return 0, false
		}
	}

	if value == "" {
		respondValidationError(c, errors.New("identificador requerido"))
		return 0, false
	}

	for _, character := range value {
		id = id*10 + uint64(character-'0')
	}

	if id == 0 {
		respondValidationError(c, errors.New("el identificador debe ser mayor que cero"))
		return 0, false
	}

	return uint(id), true
}
