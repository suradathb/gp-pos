// models/order.go
package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Total      float64     `json:"total"`
	Paid       bool        `json:"paid"`
	EmployeeID uint        `json:"employee_id"`
	BranchID   uint        `json:"branch_id"`
	Items      []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
}
