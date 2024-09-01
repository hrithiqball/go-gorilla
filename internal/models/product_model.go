package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID           string   `gorm:"primaryKey" json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Price        uint     `json:"price"`
	Stock        int      `json:"stock"`
	Photos       []string `gorm:"type:text[]" json:"photos"`
	FeaturePhoto string   `gorm:"type:text" json:"feature_photo"`
	Type         string   `json:"type"`
	BusinessID   string   `json:"businessId"`
	Business     Business `gorm:"foreignKey:BusinessID" json:"business"`
}

type ProductUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
}

type ProductResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Stock       int    `json:"stock"`
	BusinessID  string `json:"businessId"`
}
