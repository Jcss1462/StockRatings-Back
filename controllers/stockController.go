package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcss1462/StockRatings-Back/services"
)

// UpdateStockFromAPI recibe la solicitud HTTP y delega la lógica al servicio
func Sync(c *gin.Context) {
	clientToken := c.GetHeader("Authorization")
	if clientToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorización requerido"})
		return
	}

	// Llama al servicio para actualizar los stocks
	err := services.UpdateStockFromAPI(clientToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Datos de Stock actualizados correctamente"})
}
