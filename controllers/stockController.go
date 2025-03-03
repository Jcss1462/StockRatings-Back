package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcss1462/StockRatings-Back/services"
)

// SyncHandler actualiza los datos de los stocks desde una API externa
// @Summary Sincroniza los datos de los stocks
// @Description Llama a un servicio externo para actualizar los datos de los stocks. Requiere un token de autorización en el header.
// @Tags Stocks
// @Produce json
// @Param Authorization header string true "Token de autorización (Bearer)"
// @Router /stocks/sync [post]
func SyncHandler(c *gin.Context) {
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

// GetAllStocksHandler obtiene todos los stocks disponibles
// @Summary Obtiene todos los stocks
// @Description Retorna la lista completa de acciones disponibles en la base de datos.
// @Tags Stocks
// @Produce  json
// @Success 200 {array} models.Stock "Lista de acciones disponibles"
// @Router /stocks [get]
func GetAllStocksHandler(c *gin.Context) {
	stocks, err := services.GetAllStocks()
	if err != nil {
		c.Error(err) // Usa el middleware de errores
		return
	}
	c.JSON(http.StatusOK, stocks)
}

// GetBestInvestmentHandler obtiene la mejor inversión disponible
// @Summary Obtiene la mejor inversión
// @Description Retorna la acción con mayor potencial de crecimiento según los datos almacenados. Si no hay recomendaciones, devuelve `null`.
// @Tags Stocks
// @Produce json
// @Success 200 {object} models.Stock "La mejor acción recomendada"
// @Router /stocks/recommendation [get]
func GetBestInvestmentHandler(c *gin.Context) {
	bestStock, err := services.GetBestInvestment()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bestStock)
}
