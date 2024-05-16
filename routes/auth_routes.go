// routes/auth_routes.go
package routes

import (
	"go-pos/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api/v1/auth")
	{
		api.POST("/register", controllers.Register(db))
		api.POST("/login", controllers.Login(db))
	}
}
