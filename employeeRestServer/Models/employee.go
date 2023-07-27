package Models

import "gorm"

// Employees
type Employee struct {
	Name    string  `json:"Name"`
	Address Address `gorm:"foreignKey:UserId"`
	gorm.Model
}
