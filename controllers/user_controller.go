package controllers

import (
	"net/http"

	"Proyecto_2B/dto"
	"Proyecto_2B/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(
	userService *services.UserService,
) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// Register godoc
// @Summary Registrar usuario
// @Description Crea un usuario nuevo y guarda su contraseña cifrada con BCrypt.
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.RegisterUserRequest true "Datos del nuevo usuario"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} dto.APIError
// @Failure 500 {object} dto.APIError
// @Router /users/register [post]
func (controller *UserController) Register(c *gin.Context) {
	var request dto.RegisterUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		respondValidationError(c, err)
		return
	}

	response, err := controller.UserService.Register(request)
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

// Login godoc
// @Summary Iniciar sesión
// @Description Valida las credenciales y devuelve un token JWT.
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Credenciales del usuario"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.APIError
// @Failure 401 {object} dto.APIError
// @Failure 500 {object} dto.APIError
// @Router /users/login [post]
func (controller *UserController) Login(c *gin.Context) {
	var request dto.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		respondValidationError(c, err)
		return
	}

	response, err := controller.UserService.Login(request)
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// FindByID godoc
// @Summary Buscar usuario por ID
// @Description Obtiene los datos públicos de un usuario.
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID del usuario"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.APIError
// @Failure 401 {object} dto.APIError
// @Failure 404 {object} dto.APIError
// @Router /users/{id} [get]
func (controller *UserController) FindByID(c *gin.Context) {
	id, valid := parseID(c, "id")
	if !valid {
		return
	}

	response, err := controller.UserService.FindByID(id)
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// Update godoc
// @Summary Actualizar usuario
// @Description Actualiza uno o varios datos de un usuario existente.
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID del usuario"
// @Param request body dto.UpdateUserRequest true "Datos que se actualizarán"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.APIError
// @Failure 401 {object} dto.APIError
// @Failure 404 {object} dto.APIError
// @Router /users/{id} [put]
func (controller *UserController) Update(c *gin.Context) {
	id, valid := parseID(c, "id")
	if !valid {
		return
	}

	var request dto.UpdateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		respondValidationError(c, err)
		return
	}

	response, err := controller.UserService.Update(id, request)
	if err != nil {
		respondError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

// Delete godoc
// @Summary Eliminar usuario
// @Description Elimina un usuario de la base de datos.
// @Tags Users
// @Security BearerAuth
// @Param id path int true "ID del usuario"
// @Success 204 "Usuario eliminado correctamente"
// @Failure 400 {object} dto.APIError
// @Failure 401 {object} dto.APIError
// @Failure 404 {object} dto.APIError
// @Router /users/{id} [delete]
func (controller *UserController) Delete(c *gin.Context) {
	id, valid := parseID(c, "id")
	if !valid {
		return
	}

	if err := controller.UserService.Delete(id); err != nil {
		respondError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
