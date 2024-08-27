package models

import (
	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	ID              string `gorm:"primaryKey" json:"id"`
	Name            string `json:"name"`
	CoverPhoto      string `json:"coverPhoto"`
	ProfilePhoto    string `json:"profilePhoto"`
	Email           string `gorm:"unique" json:"email"`
	Phone           string `json:"phone"`
	Address         string `json:"address"`
	Website         string `json:"website"`
	BusinessOwnerID string `json:"businessOwnerID"`
	BusinessOwner   User   `gorm:"foreignKey:BusinessOwnerID" json:"businessOwner"`
}

type BusinessUpdate struct {
	Name         *string `json:"name"`
	CoverPhoto   *string `json:"coverPhoto"`
	ProfilePhoto *string `json:"profilePhoto"`
	Phone        *string `json:"phone"`
	Address      *string `json:"address"`
	Website      *string `json:"website"`
}
