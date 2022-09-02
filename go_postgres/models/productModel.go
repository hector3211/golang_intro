package models

import (
	"gorm.io/gorm"
)

// We setup how we want our table/ model to look like
type Product struct {
	gorm.Model
	Name     string
	Quantity int
}
