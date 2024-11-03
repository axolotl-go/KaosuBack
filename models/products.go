package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model

	Name        string  `gorm:"not null" json:"name"`
	Description string  `json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Images      string  `gorm:"not null" json:"images"`
	Stock       int     `gorm:"not null" json:"stock"`
	Category    string  `json:"category"`
	IsAvailable bool    `gorm:"default:false" json:"is_available"`
}
