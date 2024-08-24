package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           string `gorm:"primaryKey" json:"id"`
	Name         string `gorm:"not null" json:"name"`
	Email        string `gorm:"unique" json:"email"`
	PasswordHash string `gorm:"not null" json:"-"`
}

type UserUpdate struct {
	Name string
}
