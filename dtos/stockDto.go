package dtos

import (
	"errors"
	"strconv"
	"strings"

	"github.com/jcss1462/StockRatings-Back/models"
)

type StockDto struct {
	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

// ConvertToStock convierte el DTO a la estructura Stock
func (s *StockDto) ConvertToStock() (models.Stock, error) {
	var stock models.Stock
	var err error

	stock.Ticker = s.Ticker

	// Conversión manual de los valores numéricos
	stock.TargetFrom, err = parsePrice(s.TargetFrom)
	if err != nil {
		return models.Stock{}, errors.New("error al convertir target_from")
	}

	stock.TargetTo, err = parsePrice(s.TargetTo)
	if err != nil {
		return models.Stock{}, errors.New("error al convertir target_to")
	}

	stock.Company = s.Company
	stock.Action = s.Action
	stock.Brokerage = s.Brokerage
	stock.RatingFrom = s.RatingFrom
	stock.RatingTo = s.RatingTo
	stock.Time = s.Time

	return stock, nil
}

// parsePrice convierte "$14.00" a float64
func parsePrice(priceStr string) (float64, error) {

	cleaned := ""

	if priceStr == "" {
		cleaned = "0.00"
	} else {
		cleaned = strings.ReplaceAll(priceStr, "$", "")
		cleaned = strings.ReplaceAll(cleaned, ",", "")
	}

	return strconv.ParseFloat(cleaned, 64)
}
