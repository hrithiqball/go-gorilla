package services

import (
	"local_my_api/internal/models"
	"local_my_api/internal/repositories"
	"local_my_api/pkg/utils"
)

type ProductService interface {
	CreateProductService(product *models.Product) (*models.Product, error)
	GetProductListService(pagination utils.Pagination, businessID string) ([]models.Product, int64, error)
	GetProductService(id string) (*models.Product, error)
	UpdateProductService(id string, product *models.ProductUpdate) (*models.Product, error)
	DeleteProductService(id string) error
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{productRepository: repo}
}

func (s *productService) CreateProductService(p *models.Product) (*models.Product, error) {
	return s.productRepository.CreateProduct(p)
}

func (s *productService) GetProductListService(pagination utils.Pagination, businessID string) ([]models.Product, int64, error) {
	return s.productRepository.GetProductList(pagination, businessID)
}

func (s *productService) GetProductService(id string) (*models.Product, error) {
	return s.productRepository.GetProductByID(id)
}

func (s *productService) UpdateProductService(id string, p *models.ProductUpdate) (*models.Product, error) {
	return s.productRepository.UpdateProduct(id, p)
}

func (s *productService) DeleteProductService(id string) error {
	return s.productRepository.DeleteProduct(id)
}
