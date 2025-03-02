package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jcss1462/StockRatings-Back/config"
	"github.com/jcss1462/StockRatings-Back/dtos"
	"github.com/jcss1462/StockRatings-Back/models"
)

// Obtener todos los stocks
// Endpoint para actualizar la tabla Stock desde una API externa
func UpdateStockFromAPI(c *gin.Context) {
	// Obtener el token del header de la solicitud
	clientToken := c.GetHeader("Authorization")
	if clientToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorización requerido"})
		return
	}

	// URL de la API externa (cambia esto por la API real)
	apiURL := os.Getenv("API_STOCK_URL")

	var stockToInsert []models.Stock
	var nextPage = ""
	for { // Bucle infinito
		response := apiRequest(clientToken, apiURL, nextPage, c)
		stockToInsert = append(stockToInsert, response.Items...)
		nextPage = response.Next_page
		if nextPage == "" { // Condición de salida
			break
		}
	}

	// Limpiar la tabla actual y actualizar con los nuevos datos
	config.DB.Exec("DELETE FROM stocks") // Borra todos los registros
	if err := config.DB.Create(&stockToInsert).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar los datos en la BD"})
		return
	}

	// Responder con los datos insertados
	c.JSON(http.StatusOK, gin.H{
		"message": "Datos de Stock actualizados correctamente",
	})
}

func apiRequest(clientToken, apiUrl, nextPage string, c *gin.Context) dtos.ApiStockResponse {

	fullUrl := apiUrl
	if nextPage != "" {
		fullUrl += ("?next_page=" + nextPage)
	}

	// Crear una nueva solicitud HTTP GET
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creando la solicitud HTTP"})
		return dtos.ApiStockResponse{}
	}

	// Agregar el token al header de la solicitud a la API externa
	req.Header.Set("Authorization", clientToken)

	// Hacer la solicitud con un cliente HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hacer la solicitud a la API"})
		return dtos.ApiStockResponse{}
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer la respuesta de la API"})
		return dtos.ApiStockResponse{}
	}

	// Parsear el JSON recibido
	var response dtos.ApiStockResponse
	if err := json.Unmarshal(body, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al parsear el JSON recibido"})
		return dtos.ApiStockResponse{}
	}

	return response
}
