package services

import (
	"local_my_api/internal/models"
	"local_my_api/internal/repositories"
	"local_my_api/pkg/utils"
)

type UserService interface {
	GetUserListService(pagination utils.Pagination) ([]models.User, int64, error)
	GetUserService(id string) (*models.User, error)
	GetUserByEmailService(email string) (*models.User, error)
	UpdateUserService(id string, user *models.UserUpdate) (*models.User, error)
	DeleteUserService(id string) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{userRepository: repo}
}

func (s *userService) GetUserListService(pagination utils.Pagination) ([]models.User, int64, error) {
	return s.userRepository.GetUserList(pagination)
}

func (s *userService) GetUserService(id string) (*models.User, error) {
	return s.userRepository.GetByUserByID(id)
}

func (s *userService) GetUserByEmailService(email string) (*models.User, error) {
	return s.userRepository.GetUserByEmail(email)
}

func (s *userService) UpdateUserService(id string, userUpdate *models.UserUpdate) (*models.User, error) {
	return s.userRepository.UpdateUser(id, userUpdate)
}

func (s *userService) DeleteUserService(id string) error {
	return s.userRepository.DeleteUser(id)
}
