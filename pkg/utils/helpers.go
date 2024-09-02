package utils

import (
	"log"
	"strconv"

	"github.com/joho/godotenv"
	gonanoid "github.com/matoous/go-nanoid/v2"
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

func GenerateNanoID() string {
	s, error := gonanoid.New()
	if error != nil {
		log.Fatalf("Error generating id")
		return ""
	}
	return s
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

func ParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -999
	}
	return i
}

func ParseUint(s string) uint {
	u64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0.0
	}
	return uint(u64)
}
