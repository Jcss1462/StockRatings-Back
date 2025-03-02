package repositories

import (
	"time"

	"github.com/jcss1462/StockRatings-Back/config"
	"github.com/jcss1462/StockRatings-Back/models"
)

// InsertStocks inserta los datos en la BD a medida que se reciben
func InsertStocks(stocks []models.Stock) error {
	db := config.DB

	// Inicia una transacción para mayor seguridad
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	for _, stock := range stocks {
		err := tx.Exec(`
			INSERT INTO stocks (ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time, created_at, updated_at) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
			ON CONFLICT (ticker) DO UPDATE 
			SET target_from = EXCLUDED.target_from, 
				target_to = EXCLUDED.target_to, 
				company = EXCLUDED.company, 
				action = EXCLUDED.action, 
				brokerage = EXCLUDED.brokerage, 
				rating_from = EXCLUDED.rating_from, 
				rating_to = EXCLUDED.rating_to, 
				time = EXCLUDED.time,
				updated_at = NOW() -- Esto actualiza la fecha solo en caso de conflicto
			`, stock.Ticker, stock.TargetFrom, stock.TargetTo, stock.Company, stock.Action, stock.Brokerage, stock.RatingFrom, stock.RatingTo, stock.Time).Error

		if err != nil {
			tx.Rollback() // Revierte la transacción en caso de error
			return err
		}
	}

	return tx.Commit().Error
}

func GetAllStocks() ([]models.Stock, error) {
	db := config.DB
	var stocks []models.Stock

	err := db.Order("time DESC, target_to DESC").Find(&stocks).Error
	if err != nil {
		return nil, err
	}

	return stocks, nil
}

// GetBestInvestment obtiene la mejor inversión basada en TargetTo y Time
func GetBestInvestment() (*models.Stock, error) {
	db := config.DB
	var bestStock models.Stock

	err := db.Table("stocks").
		Where("DATE(time) = ?", time.Now().Format("2006-01-02")).
		Order("((target_to - target_from) / target_from) DESC").
		Limit(1).
		Find(&bestStock).Error

	var emptyStock models.Stock
	if bestStock == emptyStock {
		return nil, nil
	}

	return &bestStock, err
}
