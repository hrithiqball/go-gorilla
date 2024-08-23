package handler

import (
	"encoding/json"
	"io"
	"local_my_api/db"
	"local_my_api/model"
	"local_my_api/service"
	"local_my_api/utils"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

const uploadDir = "./public/uploads"

func init() {
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		panic("failed to create upload directory" + err.Error())
	}
}

func CreateBusiness(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	address := r.FormValue("address")
	website := r.FormValue("website")
	businessOwnerID := r.FormValue("businessOwnerID")

	coverPhotoFile, coverPhotoheader, err := r.FormFile("coverPhoto")
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, "Failed to get cover photo", http.StatusBadRequest)
		return
	}
	var coverPhotoPath string
	if coverPhotoFile != nil {
		coverPhotoPath, err = saveFile(coverPhotoFile, coverPhotoheader)
		if err != nil {
			http.Error(w, "Failed to save cover photo", http.StatusInternalServerError)
			return
		}
	}

	profilePhotoFile, profilePhotoHeader, err := r.FormFile("profilePhoto")
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, "Failed to get profile photo", http.StatusBadRequest)
		return
	}
	var profilePhotoPath string
	if profilePhotoFile != nil {
		profilePhotoPath, err = saveFile(profilePhotoFile, profilePhotoHeader)
		if err != nil {
			http.Error(w, "Failed to save profile photo", http.StatusInternalServerError)
			return
		}
	}

	business, err := service.CreateBusiness(name, email, phone, address, website, coverPhotoPath, profilePhotoPath, businessOwnerID)
	if err != nil || business == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(business)
}

func GetBusinessList(w http.ResponseWriter, r *http.Request) {
	var businesseList = []model.Business{}

	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	pagination := utils.ParsePagination(pageStr, sizeStr)

	if err := db.DB.Offset(pagination.Offset).Limit(pagination.Size).Find(&businesseList).Error; err != nil {
		http.Error(w, "Error retrieving businesses", http.StatusInternalServerError)
		return
	}

	var businessListCount int64
	if err := db.DB.Model(&model.Business{}).Count(&businessListCount).Error; err != nil {
		http.Error(w, "Error counting businesses", http.StatusInternalServerError)
		return
	}

	response := struct {
		Businesses []model.Business `json:"businessList"`
		Total      int64            `json:"total"`
		Page       int              `json:"page"`
		PageSize   int              `json:"pageSize"`
	}{
		Businesses: businesseList,
		Total:      businessListCount,
		Page:       pagination.Page,
		PageSize:   pagination.Size,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func GetBusiness(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	business, err := model.GetBusiness(db.DB, vars["id"])
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Business not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error retrieving business", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(business); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func saveFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	defer file.Close()

	fileName := fileHeader.Filename
	filePath := filepath.Join(uploadDir, fileName)

	destFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
