package models

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey;autoIncrement"` // ID autoincremental
	Ticker     string `gorm:"unique;not null"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}
