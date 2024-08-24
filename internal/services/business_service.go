package services

import (
	"local_my_api/internal/models"
	"local_my_api/internal/repositories"
	"local_my_api/pkg/utils"
)

type BusinessService interface {
	CreateBusinessService(business *models.Business) (*models.Business, error)
	GetBusinessListService(pagination utils.Pagination) ([]models.Business, int64, error)
	GetBusinessService(id string) (*models.Business, error)
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
	business, err := s.businessRepository.CreateBusiness(b)
	if err != nil {
		return nil, err
	}

	return business, nil
}

func (s *businessService) GetBusinessListService(pagination utils.Pagination) ([]models.Business, int64, error) {
	businessList, count, err := s.businessRepository.GetBusinessList(pagination)
	if err != nil {
		return nil, 0, err
	}

	return businessList, count, nil
}

func (s *businessService) GetBusinessService(id string) (*models.Business, error) {
	business, err := s.businessRepository.GetBusinessByID(id)
	if err != nil {
		return nil, err
	}

	return business, nil
}

func (s *businessService) UpdateBusinessService(id string, b *models.BusinessUpdate) (*models.Business, error) {
	business, err := s.businessRepository.UpdateBusiness(id, b)
	if err != nil {
		return nil, err
	}

	return business, nil
}

func (s *businessService) DeleteBusinessService(id string) error {
	return s.businessRepository.DeleteBusiness(id)
}
