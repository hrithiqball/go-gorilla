package services

import (
	"local_my_api/internal/models"
	"local_my_api/internal/repositories"
	"local_my_api/pkg/utils"
)

type BusinessService interface {
	CreateBusinessService(business *models.Business) (*models.Business, error)
	GetBusinessListService(pagination utils.Pagination) ([]models.Business, int64, error)
	GetBusinessService(ID string) (*models.Business, error)
	UpdateBusinessService(id string, business *models.BusinessUpdate) (*models.Business, error)
	DeleteBusinessService(id string) error
}

type businessService struct {
	businessRepository repositories.BusinessRepository
}

func NewBusinessService(repo repositories.BusinessRepository) BusinessService {
	return &businessService{businessRepository: repo}
}

func (s *businessService) CreateBusinessService(b *models.Business) (*models.Business, error) {
	return s.businessRepository.CreateBusiness(b)
}

func (s *businessService) GetBusinessListService(pagination utils.Pagination) ([]models.Business, int64, error) {
	return s.businessRepository.GetBusinessList(pagination)
}

func (s *businessService) GetBusinessService(ID string) (*models.Business, error) {
	return s.businessRepository.GetBusinessByID(ID)
}

func (s *businessService) UpdateBusinessService(id string, b *models.BusinessUpdate) (*models.Business, error) {
	return s.businessRepository.UpdateBusiness(id, b)
}

func (s *businessService) DeleteBusinessService(id string) error {
	return s.businessRepository.DeleteBusiness(id)
}
