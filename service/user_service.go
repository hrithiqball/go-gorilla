package service

import (
	"fmt"
	"local_my_api/db"
	"local_my_api/model"
	"local_my_api/utils"
	"net/http"

	"gorm.io/gorm"
)

func CreateUserService(email, passwordHash, name string) error {
	err := model.CreateUser(db.DB, email, passwordHash, name)
	if err != nil {
		return err
	}

	return nil
}

func GetUserService(id string) (*model.User, int, error) {
	user, err := model.GetUser(db.DB, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusNotFound, err
		}

		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}

func GetUserByEmailService(email string) (*model.User, int, error) {
	user, err := model.GetUserByEmail(db.DB, email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusNotFound, err
		}

		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}

func GetUserListService(pagination utils.Pagination) ([]model.User, int64, int, error) {
	userList := []model.User{}
	totalUsers := int64(0)

	if err := db.DB.Offset(pagination.Offset).Limit(pagination.Size).Find(&userList).Error; err != nil {
		return userList, totalUsers, http.StatusInternalServerError, fmt.Errorf("failed to retrieve users: %w", err)
	}

	if err := db.DB.Model(&model.User{}).Count(&totalUsers).Error; err != nil {
		return userList, totalUsers, http.StatusInternalServerError, fmt.Errorf("failed to count users: %w", err)
	}

	return userList, totalUsers, http.StatusOK, nil
}

func UpdateUserService(id, name string) (*model.User, int, error) {
	user, err := model.UpdateUser(db.DB, id, name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusNotFound, err
		}

		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil
}

func DeleteUserService(id string) (int, error) {
	err := model.DeleteUser(db.DB, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return http.StatusNotFound, err
		}

		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
