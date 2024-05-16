// routes/order_routes.go
package routes

import (
	"go-pos/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterOrderRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		api.POST("/orders", controllers.CreateOrder(db))
		api.GET("/orders/:id", controllers.GetOrder(db))
		api.POST("/orders/:id/pay", controllers.PayOrder(db))
		api.GET("/sales-summary", controllers.GetSalesSummary(db))
	}
}
