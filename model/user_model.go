package model

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	Name         string
	Email        string `gorm:"unique"`
	PasswordHash string
}

func CreateUser(db *gorm.DB, email, passwordHash, name string) error {
	id, err := gonanoid.New()
	if err != nil {
		return err
	}
	user := User{ID: id, Name: name, Email: email, PasswordHash: passwordHash}
	return db.Create(&user).Error
}

func GetUser(db *gorm.DB, id string) (*User, error) {
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(db *gorm.DB, id, name string) (*User, error) {
	user, err := GetUser(db, id)
	if err != nil {
		return nil, err
	}

	user.Name = name

	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(db *gorm.DB, id string) error {
	user, err := GetUser(db, id)
	if err != nil {
		return err
	}

	return db.Delete(&user).Error
}
