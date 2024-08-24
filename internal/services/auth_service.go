package services

import (
	"fmt"
	"local_my_api/internal/models"
	"local_my_api/internal/repositories"
	"local_my_api/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterService(email, password, name string) error
	LoginService(email, password string) (string, error)
	RefreshTokenService(refreshToken string) (string, error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) AuthService {
	return &authService{userRepository: repo}
}

func (s *authService) RegisterService(email, password, name string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	return s.userRepository.CreateUser(&models.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(hash),
	})
}

func (s *authService) LoginService(email, password string) (string, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("user not found: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", fmt.Errorf("invalid password: %w", err)
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}

func (s *authService) RefreshTokenService(refreshToken string) (string, error) {
	userID, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token: %w", err)
	}

	token, err := utils.GenerateJWT(userID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}
