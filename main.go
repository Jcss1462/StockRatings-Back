package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/jcss1462/StockRatings-Back/config"
	_ "github.com/jcss1462/StockRatings-Back/docs"
	"github.com/jcss1462/StockRatings-Back/routes"
	"github.com/jcss1462/StockRatings-Back/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Stock Ratings API
// @version 1.0
// @description API para manejar stocks y recomendaciones de inversión.
// @host localhost:8081
// @BasePath /
func main() {

	//cargo las variables de entorno
	utils.LoadEnvFile()

	// Conectar a la BD
	config.InitDB()

	//configuro e servidor y registro las rutas
	server := gin.Default()
	routes.RegisterRoutes(server)

	// Obtener el puerto y validar que no esté vacío
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8081" // Valor por defecto si no está definido en el .env
	}

	// Agregar Swagger solo en desarrollo
	if os.Getenv("APP_ENV") != "production" {
		utils.GenerateSwaggerDocs() // 🔹 Generar Swagger antes de servirlo
		server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		utils.OpenSwaggerUI(runtime.GOOS, appPort) // 🔹 Abrir automáticamente en el navegador
	}

	//corro el servidor
	// Iniciar el servidor
	fmt.Println("🚀 Servidor corriendo en http://localhost:" + appPort)
	server.Run(":" + appPort)
}
