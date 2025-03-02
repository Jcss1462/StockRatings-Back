package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcss1462/StockRatings-Back/services"
)

// UpdateStockFromAPI recibe la solicitud HTTP y delega la lógica al servicio
func Sync(c *gin.Context) {
	clientToken := c.GetHeader("Authorization")
	if clientToken == "" {

		c.Error(&gin.Error{
			Err:  errors.New("token de autorización requerido"),
			Type: gin.ErrorTypePublic,
			Meta: http.StatusUnauthorized,
		})

		return
	}

	// Llama al servicio para actualizar los stocks
	err := services.UpdateStockFromAPI(clientToken)
	if err != nil {
		c.Error(&gin.Error{
			Err:  err,
			Type: gin.ErrorTypePublic,
			Meta: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Datos de Stock actualizados correctamente"})
}

func GetAllStocks(c *gin.Context) {
	stocks, err := services.GetAllStocks()
	if err != nil {
		c.Error(err) // Usa el middleware de errores
		return
	}
	c.JSON(http.StatusOK, stocks)
}
