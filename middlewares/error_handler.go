package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorHandlerMiddleware captura errores y responde con JSON
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Continúa con la ejecución de otros middlewares y handlers

		// Verifica si hay errores en la cola de errores de Gin
		if len(c.Errors) > 0 {

			for _, err := range c.Errors {
				if err.Type == gin.ErrorTypePublic {
					// Obtiene el código HTTP del campo Meta (si existe)
					statusCode, ok := err.Meta.(int)
					if !ok {
						statusCode = http.StatusBadRequest // Código por defecto
					}
					c.JSON(statusCode, gin.H{"error": err.Error()})
					return
				}
			}

			// Si no son errores públicos, devuelve un error genérico
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		}
	}
}
