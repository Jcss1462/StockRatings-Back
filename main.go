package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/gin-contrib/cors"
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

	// Obtener el puerto y validar que no esté vacío
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8081" // Valor por defecto si no está definido en el .env
	}

	// Conectar a la BD
	config.InitDB()

	//Inicializo el servidor
	server := gin.Default()

	//Configuro CORS
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Agregar Swagger (solo en desarrollo)
	if os.Getenv("APP_ENV") != "production" {
		utils.GenerateSwaggerDocs() // 🔹 Generar Swagger antes de servirlo
		server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		utils.OpenSwaggerUI(runtime.GOOS, appPort) // 🔹 Abrir automáticamente en el navegador
	}

	//registro las rutas
	routes.RegisterRoutes(server)

	//corro el servidor
	// Iniciar el servidor
	fmt.Println("🚀 Servidor corriendo en http://localhost:" + appPort)
	server.Run(":" + appPort)
}
