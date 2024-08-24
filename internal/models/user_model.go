package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	Name         string
	Email        string `gorm:"unique"`
	PasswordHash string
}

type UserUpdate struct {
	Name string
}
