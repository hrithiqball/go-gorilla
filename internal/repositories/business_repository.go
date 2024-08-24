package repositories

import (
	"fmt"
	"local_my_api/internal/models"
	"local_my_api/pkg/utils"

	"gorm.io/gorm"
)

type BusinessRepository interface {
	CreateBusiness(business *models.Business) (*models.Business, error)
	GetBusinessList(pagination utils.Pagination) ([]models.Business, int64, error)
	GetBusinessByID(id string) (*models.Business, error)
	UpdateBusiness(id string, business *models.BusinessUpdate) (*models.Business, error)
	DeleteBusiness(id string) error
}

type businessRepository struct {
	db *gorm.DB
}

func NewBusinessRepository(db *gorm.DB) BusinessRepository {
	return &businessRepository{db: db}
}

func (r *businessRepository) CreateBusiness(b *models.Business) (*models.Business, error) {
	id := utils.GenerateNanoID()

	business := models.Business{
		ID:              id,
		Name:            b.Name,
		Email:           b.Email,
		Phone:           b.Phone,
		Address:         b.Address,
		Website:         b.Website,
		CoverPhoto:      b.CoverPhoto,
		ProfilePhoto:    b.ProfilePhoto,
		BusinessOwnerID: b.BusinessOwnerID,
	}

	return &business, r.db.Create(business).Error
}

func (r *businessRepository) GetBusinessList(pagination utils.Pagination) ([]models.Business, int64, error) {
	businessList := []models.Business{}
	totalBusiness := int64(0)

	if err := r.db.Offset(pagination.Offset).Limit(pagination.Size).Find(&businessList).Error; err != nil {
		return nil, totalBusiness, fmt.Errorf("failed to retrieve business list: %w", err)
	}

	if err := r.db.Model(&models.Business{}).Count(&totalBusiness).Error; err != nil {
		return nil, totalBusiness, fmt.Errorf("failed to retrieve business list: %w", err)
	}

	return businessList, totalBusiness, nil
}

func (r *businessRepository) GetBusinessByID(id string) (*models.Business, error) {
	var business models.Business
	if err := r.db.Preload("BusinessOwner").First(&business, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &business, nil
}

func (r *businessRepository) UpdateBusiness(id string, b *models.BusinessUpdate) (*models.Business, error) {
	business, err := r.GetBusinessByID(id)
	if err != nil {
		return nil, fmt.Errorf("business not found")
	}

	if err := r.db.Model(business).Updates(b).Error; err != nil {
		return nil, fmt.Errorf("failed to update business: %w", err)
	}

	return business, nil
}

func (r *businessRepository) DeleteBusiness(id string) error {
	business, err := r.GetBusinessByID(id)
	if err != nil {
		return fmt.Errorf("business not found")
	}

	return r.db.Delete(business).Error
}
