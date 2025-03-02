package dtos

import "github.com/jcss1462/StockRatings-Back/models"

type ApiStockResponse struct {
	Items     []models.Stock
	Next_page string
}
