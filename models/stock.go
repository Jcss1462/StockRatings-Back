package models

import (
	"gorm.io/gorm"
)

// Stock representa un activo financiero
// @Description Modelo de una acción en la base de datos
type Stock struct {
	gorm.Model `swaggerignore:"true"` // Ignorar gorm.Model en Swagger
	Ticker     string                 `json:"ticker" gorm:"unique;not null" example:"AAPL"`
	TargetFrom float64                `json:"target_from" example:"150.00"`
	TargetTo   float64                `json:"target_to" example:"180.00"`
	Company    string                 `json:"company" example:"Apple Inc."`
	Action     string                 `json:"action" example:"Buy"`
	Brokerage  string                 `json:"brokerage" example:"Goldman Sachs"`
	RatingFrom string                 `json:"rating_from" example:"Neutral"`
	RatingTo   string                 `json:"rating_to" example:"Buy"`
	Time       string                 `json:"time" example:"2025-03-02"`
}
