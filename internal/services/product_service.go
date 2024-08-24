package services

import (
	"local_my_api/internal/models"
	"local_my_api/internal/repositories"
	"local_my_api/pkg/utils"
)

type ProductService interface {
	CreateProductService(product *models.Product) (*models.Product, error)
	GetProductListService(pagination utils.Pagination) ([]models.Product, int64, error)
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
	product, err := s.productRepository.CreateProduct(p)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) GetProductListService(pagination utils.Pagination) ([]models.Product, int64, error) {
	productList, count, err := s.productRepository.GetProductList(pagination)
	if err != nil {
		return nil, count, err
	}

	return productList, count, nil
}

func (s *productService) GetProductService(id string) (*models.Product, error) {
	product, err := s.productRepository.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) UpdateProductService(id string, p *models.ProductUpdate) (*models.Product, error) {
	product, err := s.productRepository.UpdateProduct(id, p)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) DeleteProductService(id string) error {
	return s.productRepository.DeleteProduct(id)
}
