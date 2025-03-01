package main

import (
	"github.com/jcss1462/StockRatings-Back/config"
	"github.com/jcss1462/StockRatings-Back/models"
)

func main() {

	// Conectar a la BD
	config.InitDB()

	// Crear tablas automáticamente
	config.DB.AutoMigrate(&models.Stock{})
}
