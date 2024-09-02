package models

import (
	"time"

	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	ID              string    `gorm:"primaryKey" json:"id"`
	Name            string    `json:"name"`
	CoverPhoto      string    `json:"coverPhoto"`
	ProfilePhoto    string    `json:"profilePhoto"`
	Email           string    `gorm:"unique" json:"email"`
	Phone           string    `json:"phone"`
	Address         string    `json:"address"`
	Website         string    `json:"website"`
	BusinessOwnerID string    `json:"businessOwnerID"`
	BusinessOwner   User      `gorm:"foreignKey:BusinessOwnerID" json:"businessOwner"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Products        []Product `gorm:"foreignKey:BusinessID" json:"products"`
}

type BusinessUpdate struct {
	Name         *string `json:"name"`
	CoverPhoto   *string `json:"coverPhoto"`
	ProfilePhoto *string `json:"profilePhoto"`
	Phone        *string `json:"phone"`
	Address      *string `json:"address"`
	Website      *string `json:"website"`
}

type BusinessResponse struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	Phone           string            `json:"phone"`
	Email           string            `json:"email"`
	Website         string            `json:"website"`
	CoverPhoto      string            `json:"coverPhoto"`
	ProfilePhoto    string            `json:"profilePhoto"`
	CreatedAt       time.Time         `json:"createdAt"`
	UpdatedAt       time.Time         `json:"updatedAt"`
	Address         string            `json:"address"`
	BusinessOwnerID string            `json:"businessOwnerId"`
	Products        []ProductResponse `json:"products"`
}
