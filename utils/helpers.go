package utils

import (
	"log"
	"os"
	"strconv"

	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type Pagination struct {
	Page   int
	Size   int
	Offset int
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GenerateJWT(userID string) (string, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateRefreshToken(tokenString string) (string, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["sub"].(string)
		if !ok {
			return "", fmt.Errorf("invalid token claims")
		}
		return userID, nil
	}
	return "", fmt.Errorf("invalid token")
}

func ValidateJWT(tokenString string) (string, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["sub"].(string)
		if !ok {
			return "", fmt.Errorf("invalid token claims")
		}
		return userID, nil
	}
	return "", fmt.Errorf("invalid token")
}

func ParsePagination(pageStr string, sizeStr string) Pagination {
	page := 1
	size := 10

	if pageStr != "" {
		var err error
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}
	}

	if sizeStr != "" {
		var err error
		size, err = strconv.Atoi(sizeStr)
		if err != nil || size < 1 {
			size = 10
		}
	}

	offset := (page - 1) * size

	return Pagination{Page: page, Size: size, Offset: offset}
}
