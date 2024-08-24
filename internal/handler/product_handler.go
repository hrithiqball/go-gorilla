package handler

import (
	"local_my_api/internal/models"
	"local_my_api/internal/services"
	"local_my_api/pkg/utils"
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

type ProductResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func NewProductHandler(productService services.ProductService, businessService services.BusinessService) ProductHandler {
	return &productHandler{productService: productService,
		businessService: businessService}
}

func (h *productHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	businessID := r.FormValue("businessId")
	name := r.FormValue("name")
	description := r.FormValue("description")
	priceStr := r.FormValue("price")
	stockStr := r.FormValue("stock")

	// validate form abd construct product

	price := utils.ParseInt(priceStr)
	stock := utils.ParseInt(stockStr)

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

		utils.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if business.BusinessOwnerID != userID {
		utils.ResponseWithError(w, http.StatusForbidden, "Forbidden")
		return
	}

	product, err := h.productService.CreateProductService(&models.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
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

	pageStr := r.FormValue("page")
	sizeStr := r.FormValue("size")

	pagination := utils.ParsePagination(pageStr, sizeStr)

	productList, count, err := h.productService.GetProductListService(pagination)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	productResponse := []ProductResponse{}
	for _, product := range productList {
		productResponse = append(productResponse, toProductResponse(&product))
	}

	response := struct {
		Products []ProductResponse `json:"productList"`
		Total    int64             `json:"total"`
		Page     int               `json:"page"`
		PageSize int               `json:"pageSize"`
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

func toProductResponse(product *models.Product) ProductResponse {
	return ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	}
}
