package repositories

import (
	"fmt"
	"local_my_api/internal/models"
	"local_my_api/pkg/utils"
	"log"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetByUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserList(pagination utils.Pagination) ([]models.User, int64, error)
	GetUserBusiness(id string) ([]models.Business, error)
	UpdateUser(id string, user *models.UserUpdate) (*models.User, error)
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(u *models.User) error {
	u.ID = utils.GenerateNanoID()

	return r.db.Create(u).Error
}

func (r *userRepository) GetByUserByID(id string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserList(pagination utils.Pagination) ([]models.User, int64, error) {
	userList := []models.User{}
	totalUsers := int64(0)

	if err := r.db.Offset(pagination.Offset).Limit(pagination.Size).Find(&userList).Error; err != nil {
		return nil, totalUsers, fmt.Errorf("failed to retrieve user list: %w", err)
	}

	if err := r.db.Model(&models.User{}).Count(&totalUsers).Error; err != nil {
		return nil, totalUsers, fmt.Errorf("failed to count user list: %w", err)
	}

	return userList, totalUsers, nil
}

func (r *userRepository) GetUserBusiness(id string) ([]models.Business, error) {
	var businessList []models.Business
	query := r.db.Model(&models.Business{}).Where("business_owner_id = ?", id).Find(&businessList)

	err := query.Error
	if err != nil {
		log.Printf("Error retrieving user business: %v", err)
		return nil, fmt.Errorf("failed to retrieve business list: %w", err)
	}

	return businessList, nil
}

func (r *userRepository) UpdateUser(id string, u *models.UserUpdate) (*models.User, error) {
	user, err := r.GetByUserByID(id)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(user).Updates(u).Error; err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return user, nil
}

func (r *userRepository) DeleteUser(id string) error {
	user, err := r.GetByUserByID(id)
	if err != nil {
		return err
	}

	return r.db.Delete(&user).Error
}
