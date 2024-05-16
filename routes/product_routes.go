// routes/product_routes.go
package routes

import (
	"go-pos/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterProductRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		api.GET("/products", controllers.GetProducts(db))
		// Other product routes...
	}
}
