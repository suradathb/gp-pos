// models/branch.go
package models

import (
	"gorm.io/gorm"
)

type Branch struct {
	gorm.Model
	Name string `json:"name"`
}
