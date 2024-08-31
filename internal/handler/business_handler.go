package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"local_my_api/internal/models"
	"local_my_api/internal/services"
	"local_my_api/internal/validation"
	"local_my_api/pkg/utils"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

const uploadDir = "./public/uploads"

type BusinessHandler interface {
	CreateBusinessHandler(w http.ResponseWriter, r *http.Request)
	GetBusinessListHandler(w http.ResponseWriter, r *http.Request)
	GetBusinessHandler(w http.ResponseWriter, r *http.Request)
	UpdateBusinessHandler(w http.ResponseWriter, r *http.Request)
	DeleteBusinessHandler(w http.ResponseWriter, r *http.Request)
}

type businessHandler struct {
	businessService services.BusinessService
}

type BusinessResponse struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Phone           string    `json:"phone"`
	Email           string    `json:"email"`
	Website         string    `json:"website"`
	CoverPhoto      string    `json:"coverPhoto"`
	ProfilePhoto    string    `json:"profilePhoto"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Address         string    `json:"address"`
	BusinessOwnerID string    `json:"businessOwnerId"`
}

func NewBusinessHandler(service services.BusinessService) BusinessHandler {
	return &businessHandler{businessService: service}
}

func init() {
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		panic("failed to create upload directory" + err.Error())
	}
}

func (h *businessHandler) CreateBusinessHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(utils.UserIDKey).(string)
	if !ok {
		utils.ResponseWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	address := r.FormValue("address")
	website := r.FormValue("website")
	businessOwnerID := userID

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

	if err := validation.ValidateCreateBusinessFormInput(name, email, phone, address, website); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	business, err := h.businessService.CreateBusinessService(&models.Business{
		Name:            name,
		Email:           email,
		Phone:           phone,
		Address:         address,
		Website:         website,
		CoverPhoto:      coverPhotoPath,
		ProfilePhoto:    profilePhotoPath,
		BusinessOwnerID: businessOwnerID,
	})

	if err != nil || business == nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	businessResponse := toBusinessResponse(business)
	utils.RespondWithJson(w, http.StatusCreated, businessResponse)
}

func (h *businessHandler) GetBusinessListHandler(w http.ResponseWriter, r *http.Request) {
	var businessList = []models.Business{}

	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	pagination := utils.ParsePagination(pageStr, sizeStr)

	businessList, businessListCount, err := h.businessService.GetBusinessListService(pagination)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	businessListResponse := []BusinessResponse{}
	for _, business := range businessList {
		businessListResponse = append(businessListResponse, *toBusinessResponse(&business))
	}

	response := struct {
		Businesses []BusinessResponse `json:"businessList"`
		Total      int64              `json:"total"`
		Page       int                `json:"page"`
		PageSize   int                `json:"pageSize"`
	}{
		Businesses: businessListResponse,
		Total:      businessListCount,
		Page:       pagination.Page,
		PageSize:   pagination.Size,
	}

	utils.RespondWithJson(w, http.StatusOK, response)
}

func (h *businessHandler) GetBusinessHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	business, err := h.businessService.GetBusinessService(vars["id"])
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ResponseWithError(w, http.StatusNotFound, "Business not found")
			return
		}

		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	businessResponse := toBusinessResponse(business)
	utils.RespondWithJson(w, http.StatusOK, businessResponse)
}

func (h *businessHandler) UpdateBusinessHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paramsID := vars["id"]
	businessOwnerID := r.FormValue("businessOwnerId")

	userID, ok := r.Context().Value(utils.UserIDKey).(string)
	if !ok {
		utils.ResponseWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if businessOwnerID != userID {
		utils.ResponseWithError(w, http.StatusForbidden, "Forbidden")
		return
	}

	name := r.FormValue("name")
	// add other fields

	// validate form data

	business, err := h.businessService.UpdateBusinessService(paramsID, &models.BusinessUpdate{
		Name: &name,
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ResponseWithError(w, http.StatusNotFound, "Business not found")
			return
		}

		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	businessResponse := toBusinessResponse(business)
	utils.RespondWithJson(w, http.StatusOK, businessResponse)
}

func (h *businessHandler) DeleteBusinessHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paramsID := vars["id"]
	businessOwnerID := r.FormValue("businessOwnerId")

	userID, ok := r.Context().Value(utils.UserIDKey).(string)
	if !ok {
		utils.ResponseWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if businessOwnerID != userID {
		utils.ResponseWithError(w, http.StatusForbidden, "Forbidden")
		return
	}

	err := h.businessService.DeleteBusinessService(paramsID)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJson(w, http.StatusOK, nil)
}

func saveFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	hashedFileName := hex.EncodeToString(hash.Sum(nil)) + ext
	filePath := filepath.Join(uploadDir, hashedFileName)

	destFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer destFile.Close()

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	_, err = io.Copy(destFile, file)
	if err != nil {
		return "", err
	}

	return "/bucket/" + hashedFileName, nil
}

func toBusinessResponse(business *models.Business) *BusinessResponse {
	return &BusinessResponse{
		ID:              business.ID,
		Name:            business.Name,
		Phone:           business.Phone,
		Email:           business.Email,
		Website:         business.Website,
		CoverPhoto:      business.CoverPhoto,
		ProfilePhoto:    business.ProfilePhoto,
		CreatedAt:       business.CreatedAt,
		UpdatedAt:       business.UpdatedAt,
		Address:         business.Address,
		BusinessOwnerID: business.BusinessOwnerID,
	}
}
