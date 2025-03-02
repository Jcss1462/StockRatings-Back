package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jcss1462/StockRatings-Back/config"
	"github.com/jcss1462/StockRatings-Back/routes"
	"github.com/jcss1462/StockRatings-Back/utils"
)

func main() {

	//cargo las variables de entorno
	utils.LoadEnvFile()

	// Conectar a la BD
	config.InitDB()

	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8081")
}
