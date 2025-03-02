package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jcss1462/StockRatings-Back/controllers"
	"github.com/jcss1462/StockRatings-Back/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	// Aplica el middleware globalmente
	server.Use(middlewares.ErrorHandlerMiddleware())

	stockGroup := server.Group("/stocks")
	{
		stockGroup.POST("/sync", controllers.SyncHandler)
		stockGroup.GET("/", controllers.GetAllStocksHandler)
		stockGroup.GET("/best-investment", controllers.GetBestInvestmentHandler)
	}

}
