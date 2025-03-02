package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jcss1462/StockRatings-Back/controllers"
)

func RegisterRoutes(server *gin.Engine) {

	stockGroup := server.Group("/stocks")
	{
		stockGroup.POST("/sync", controllers.UpdateStockFromAPI)
	}

}
