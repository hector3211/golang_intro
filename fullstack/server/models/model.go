package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
