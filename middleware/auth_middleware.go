package middleware

import (
	"net/http"
	"strings"
	"time"

	"Proyecto_2B/dto"
	"Proyecto_2B/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		if authorizationHeader == "" {
			unauthorized(
				c,
				"Se requiere el encabezado Authorization",
			)
			return
		}

		parts := strings.SplitN(
			authorizationHeader,
			" ",
			2,
		)

		if len(parts) != 2 ||
			!strings.EqualFold(parts[0], "Bearer") {

			unauthorized(
				c,
				"El formato debe ser: Bearer <token>",
			)
			return
		}

		tokenString := strings.TrimSpace(parts[1])

		if tokenString == "" {
			unauthorized(c, "El token está vacío")
			return
		}

		claims, err := utils.ValidateToken(tokenString)

		if err != nil {
			unauthorized(c, "Token inválido o expirado")
			return
		}

		c.Set("userId", claims.UserID)
		c.Set("userEmail", claims.Email)

		c.Next()
	}
}

func unauthorized(c *gin.Context, message string) {
	c.AbortWithStatusJSON(
		http.StatusUnauthorized,
		dto.APIError{
			Timestamp: time.Now(),
			Status:    http.StatusUnauthorized,
			Error:     "Unauthorized",
			Message:   message,
			Path:      c.Request.URL.Path,
		},
	)
}
