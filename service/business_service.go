package service

import (
	"local_my_api/db"
	"local_my_api/model"
)

func CreateBusiness(name, email, phone, address, website, coverPhotoPath, profilePhotoPath, businessOwnerID string) (*model.Business, error) {
	business, err := model.CreateBusiness(db.DB, name, email, phone, address, website, coverPhotoPath, profilePhotoPath, businessOwnerID)
	if err != nil {
		return nil, err
	}

	return business, nil
}

func GetBusiness(id string) (*model.Business, error) {
	business, err := model.GetBusiness(db.DB, id)
	if err != nil {
		return nil, err
	}

	return business, nil
}
