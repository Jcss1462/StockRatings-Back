package services

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/jcss1462/StockRatings-Back/dtos"
	"github.com/jcss1462/StockRatings-Back/models"
	"github.com/jcss1462/StockRatings-Back/repositories"
)

// UpdateStockFromAPI obtiene los datos de la API externa y actualiza la BD
func UpdateStockFromAPI(clientToken string) error {
	apiURL := os.Getenv("API_STOCK_URL")

	var nextPage = ""
	for {
		response, err := apiRequest(clientToken, apiURL, nextPage)
		if err != nil {
			return err
		}

		// Convertir DTOs a Stocks usando mapstructure
		var stocks []models.Stock
		for _, dto := range response.Items {
			stock, err := dto.ConvertToStock()
			if err != nil {
				return err
			}
			stocks = append(stocks, stock)
		}

		//Inserta los datos a medida que se reciben
		if err := repositories.InsertStocks(stocks); err != nil {
			return err
		}

		nextPage = response.Next_page
		if nextPage == "" {
			break
		}
	}

	return nil
}

// apiRequest realiza la solicitud a la API externa y devuelve la respuesta estructurada
func apiRequest(clientToken, apiUrl, nextPage string) (dtos.ApiStockResponseDto, error) {
	fullUrl := apiUrl
	if nextPage != "" {
		fullUrl += "?next_page=" + nextPage
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return dtos.ApiStockResponseDto{}, errors.New("error creando la solicitud HTTP")
	}

	req.Header.Set("Authorization", clientToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dtos.ApiStockResponseDto{}, errors.New("error al hacer la solicitud a la API")
	}

	if resp.StatusCode == http.StatusForbidden {
		return dtos.ApiStockResponseDto{}, errors.New("token de autenticacion invalido")
	}

	if resp.StatusCode != http.StatusOK {
		return dtos.ApiStockResponseDto{}, errors.New("error al hacer la solicitud a la API: StatusCode: " + strconv.Itoa(resp.StatusCode))
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dtos.ApiStockResponseDto{}, errors.New("error al leer la respuesta de la API")
	}

	var response dtos.ApiStockResponseDto
	if err := json.Unmarshal(body, &response); err != nil {
		return dtos.ApiStockResponseDto{}, errors.New("error al parsear el JSON recibido")
	}

	return response, nil
}

func GetAllStocks() ([]models.Stock, error) {
	return repositories.GetAllStocks()
}
