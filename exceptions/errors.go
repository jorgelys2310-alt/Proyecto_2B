package exceptions

import "errors"

var (
	ErrNotFound           = errors.New("recurso no encontrado")
	ErrBadRequest         = errors.New("solicitud inválida")
	ErrInvalidCredentials = errors.New("credenciales inválidas")
	ErrEmailAlreadyExists = errors.New("el correo ya se encuentra registrado")
	ErrInsufficientStock  = errors.New("stock insuficiente")
)
