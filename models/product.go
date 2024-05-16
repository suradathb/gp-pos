// models/product.go
package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	BranchID uint    `json:"branch_id"`
}