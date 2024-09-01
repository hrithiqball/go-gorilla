package handler

import (
	"local_my_api/internal/models"
	"local_my_api/internal/services"
	"local_my_api/pkg/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ProductHandler interface {
	CreateProductHandler(w http.ResponseWriter, r *http.Request)
	GetProductListHandler(w http.ResponseWriter, r *http.Request)
	GetProductHandler(w http.ResponseWriter, r *http.Request)
	UpdateProductHandler(w http.ResponseWriter, r *http.Request)
	DeleteProductHandler(w http.ResponseWriter, r *http.Request)
}

type productHandler struct {
	productService  services.ProductService
	businessService services.BusinessService
}

func NewProductHandler(productService services.ProductService, businessService services.BusinessService) ProductHandler {
	return &productHandler{productService: productService,
		businessService: businessService}
}

func (h *productHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(utils.UserIDKey).(string)
	if !ok {
		utils.ResponseWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := r.ParseForm(); err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	businessID := r.FormValue("businessId")
	name := r.FormValue("name")
	description := r.FormValue("description")
	productType := r.FormValue("type")
	priceStr := r.FormValue("price")
	stockStr := r.FormValue("stock")

	price := utils.ParseUint(priceStr)
	stock := utils.ParseInt(stockStr)

	featurePhotoFile, featurePhotoHeader, err := r.FormFile("featurePhoto")
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, "Failed to get feature photo", http.StatusBadRequest)
		return
	}

	var featurePhotoPath string
	if featurePhotoFile != nil {
		featurePhotoPath, err = saveFile(featurePhotoFile, featurePhotoHeader)
		if err != nil {
			http.Error(w, "Failed to save feature photo", http.StatusInternalServerError)
			return
		}
	}

	photos := r.MultipartForm.File["photos"]
	var photoPaths []string

	for _, header := range photos {
		file, err := header.Open()
		if err != nil {
			http.Error(w, "Failed to open photo file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		path, err := saveFile(file, header)
		if err != nil {
			http.Error(w, "Failed to save photo", http.StatusInternalServerError)
			return
		}
		photoPaths = append(photoPaths, path)
	}

	business, err := h.businessService.GetBusinessService(businessID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ResponseWithError(w, http.StatusNotFound, "Business not found")
			return
		}

		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if business.BusinessOwnerID != userID {
		utils.ResponseWithError(w, http.StatusForbidden, "Forbidden")
		return
	}

	product, err := h.productService.CreateProductService(&models.Product{
		Name:         name,
		Description:  description,
		Price:        price,
		Stock:        stock,
		Photos:       photoPaths,
		FeaturePhoto: featurePhotoPath,
		BusinessID:   businessID,
		Type:         productType,
	})
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	productResponse := toProductResponse(product)
	utils.RespondWithJson(w, http.StatusCreated, productResponse)
}

func (h *productHandler) GetProductListHandler(w http.ResponseWriter, r *http.Request) {
	var productList = []models.Product{}

	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")
	offsetStr := r.URL.Query().Get("offset")
	log.Printf("page: %s, size: %s, offset: %s", pageStr, sizeStr, offsetStr)

	pagination := utils.ParsePagination(pageStr, sizeStr)

	productList, count, err := h.productService.GetProductListService(pagination)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	productResponse := []models.ProductResponse{}
	for _, product := range productList {
		productResponse = append(productResponse, toProductResponse(&product))
	}

	response := struct {
		Products []models.ProductResponse `json:"productList"`
		Total    int64                    `json:"total"`
		Page     int                      `json:"page"`
		PageSize int                      `json:"pageSize"`
	}{
		Products: productResponse,
		Total:    count,
		Page:     pagination.Page,
		PageSize: pagination.Size,
	}

	utils.RespondWithJson(w, http.StatusOK, response)
}

func (h *productHandler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()

	product, err := h.productService.GetProductService(vars.Get("id"))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ResponseWithError(w, http.StatusNotFound, "Product not found")
			return
		}

		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	productResponse := toProductResponse(product)
	utils.RespondWithJson(w, http.StatusOK, productResponse)
}

func (h *productHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	productID := vars["id"]
	name := r.FormValue("name")
	description := r.FormValue("description")
	priceStr := r.FormValue("price")
	stockStr := r.FormValue("stock")

	price := utils.ParseInt(priceStr)
	stock := utils.ParseInt(stockStr)

	product, err := h.productService.UpdateProductService(productID, &models.ProductUpdate{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ResponseWithError(w, http.StatusNotFound, "Product not found")
			return
		}

		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	productResponse := toProductResponse(product)
	utils.RespondWithJson(w, http.StatusOK, productResponse)
}

func (h *productHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	businessID := r.FormValue("businessId")

	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		utils.ResponseWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	business, err := h.businessService.GetBusinessService(businessID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ResponseWithError(w, http.StatusNotFound, "Business not found")
			return
		}

		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if business.BusinessOwnerID != userID {
		utils.ResponseWithError(w, http.StatusForbidden, "Forbidden")
		return
	}

	err = h.productService.DeleteProductService(ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ResponseWithError(w, http.StatusNotFound, "Product not found")
			return
		}

		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJson(w, http.StatusOK, utils.Response{Message: "Product deleted", Status: "success"})
}

func toProductResponse(product *models.Product) models.ProductResponse {
	return models.ProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		Price:        product.Price,
		Stock:        product.Stock,
		Type:         product.Type,
		Photos:       product.Photos,
		FeaturePhoto: product.FeaturePhoto,
		BusinessID:   product.BusinessID,
	}
}
