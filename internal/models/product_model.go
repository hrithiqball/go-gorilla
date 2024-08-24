package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          string   `gorm:"primaryKey" json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Stock       int      `json:"stock"`
	BusinessID  string   `json:"businessID"`
	Business    Business `gorm:"foreignKey:BusinessID" json:"business"`
}

type ProductUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
}
