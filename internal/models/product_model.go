package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID           string          `gorm:"primaryKey" json:"id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Price        decimal.Decimal `gorm:"type:numeric" json:"price"`
	Stock        int             `json:"stock"`
	Photos       []string        `gorm:"type:text[]" json:"photos"`
	FeaturePhoto string          `gorm:"type:text" json:"feature_photo"`
	Type         string          `json:"type"`
	BusinessID   string          `json:"businessId"`
	Business     Business        `gorm:"foreignKey:BusinessID" json:"business"`
}

type ProductUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
	Type        string `json:"type"`
}

type ProductResponse struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	Name         string   `json:"name"`
	Price        float64  `json:"price"`
	Type         string   `json:"type"`
	Stock        int      `json:"stock"`
	Photos       []string `json:"photos"`
	FeaturePhoto string   `json:"featurePhoto"`
	BusinessID   string   `json:"businessId"`
	Business     Business `json:"business"`
}
