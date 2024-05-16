// main.go
package main

import (
	"go-pos/models"
	"go-pos/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}

	db.AutoMigrate(&models.Product{}, &models.Branch{}, &models.Employee{}, &models.Order{}, &models.OrderItem{})

	r := gin.Default()

	routes.RegisterProductRoutes(r, db)
	routes.RegisterOrderRoutes(r, db)
	routes.RegisterAuthRoutes(r, db)

	r.Run(":8080")
}
