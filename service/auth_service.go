package service

import (
	"fmt"
	"local_my_api/utils"

	"golang.org/x/crypto/bcrypt"
)

func RegisterService(email, password, name string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	return CreateUserService(email, string(hash), name)
}

func LoginService(email, password string) (string, error) {
	user, _, err := GetUserByEmailService(email)
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

func RefreshTokenService(refreshToken string) (string, error) {
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
