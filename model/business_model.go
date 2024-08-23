package model

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Business struct {
	gorm.Model
	ID              string `gorm:"primaryKey"`
	Name            string
	CoverPhoto      string
	ProfilePhoto    string
	Email           string `gorm:"unique"`
	Phone           string `gorm:"unique"`
	Address         string
	Website         string
	BusinessOwnerID string
	BusinessOwner   User `gorm:"foreignKey:BusinessOwnerID"`
}

func CreateBusiness(db *gorm.DB, name, email, phone, address, website, coverPhotoPath, profilePhotoPath string, ownerID string) (*Business, error) {
	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	business := Business{
		ID:              id,
		Name:            name,
		Email:           email,
		Phone:           phone,
		Address:         address,
		Website:         website,
		CoverPhoto:      coverPhotoPath,
		ProfilePhoto:    profilePhotoPath,
		BusinessOwnerID: ownerID}

	return &business, db.Create(&business).Error
}

func GetBusiness(db *gorm.DB, id string) (*Business, error) {
	var business Business
	if err := db.Preload("BusinessOwner").First(&business, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &business, nil
}
