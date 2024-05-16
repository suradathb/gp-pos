// controllers/order_controller.go
package controllers

import (
	"go-pos/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newOrder models.Order
		if err := c.ShouldBindJSON(&newOrder); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		total := 0.0
		for i, item := range newOrder.Items {
			var product models.Product
			db.First(&product, item.ProductID)
			if product.Stock < item.Quantity {
				c.JSON(http.StatusBadRequest, gin.H{"message": "not enough stock"})
				return
			}
			product.Stock -= item.Quantity
			db.Save(&product)
			total += float64(item.Quantity) * product.Price
			newOrder.Items[i].Product = product
		}
		newOrder.Total = total
		newOrder.Paid = false
		db.Create(&newOrder)
		c.JSON(http.StatusCreated, newOrder)
	}
}
