// models/employee.go
package models

import (
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	BranchID uint   `json:"branch_id"`
}
