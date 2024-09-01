package repositories

import (
	"fmt"
	"local_my_api/internal/models"
	"local_my_api/pkg/utils"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetProductList(pagination utils.Pagination) ([]models.Product, int64, error)
	GetProductByID(id string) (*models.Product, error)
	UpdateProduct(id string, product *models.ProductUpdate) (*models.Product, error)
	DeleteProduct(id string) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(p *models.Product) (*models.Product, error) {
	id := utils.GenerateNanoID()

	product := models.Product{
		ID:           id,
		Photos:       p.Photos,
		FeaturePhoto: p.FeaturePhoto,
		Type:         p.Type,
		Name:         p.Name,
		Description:  p.Description,
		Price:        p.Price,
		Stock:        p.Stock,
		BusinessID:   p.BusinessID,
	}

	return &product, r.db.Create(product).Error
}

func (r *productRepository) GetProductList(pagination utils.Pagination) ([]models.Product, int64, error) {
	productList := []models.Product{}
	totalProduct := int64(0)

	if err := r.db.Offset(pagination.Offset).Limit(pagination.Size).Find(&productList).Error; err != nil {
		return nil, totalProduct, fmt.Errorf("failed to retrieve product list: %w", err)
	}

	if err := r.db.Model(&models.Product{}).Count(&totalProduct).Error; err != nil {
		return nil, totalProduct, fmt.Errorf("failed to retrieve product list: %w", err)
	}

	return productList, totalProduct, nil
}

func (r *productRepository) GetProductByID(id string) (*models.Product, error) {
	product := models.Product{}
	if err := r.db.Preload("Business").Where("id = ?", id).First(&product).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve product: %w", err)
	}

	return &product, nil
}

func (r *productRepository) UpdateProduct(id string, p *models.ProductUpdate) (*models.Product, error) {
	product := models.Product{}
	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve product: %w", err)
	}

	if err := r.db.Model(&product).Updates(p).Error; err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return &product, nil
}

func (r *productRepository) DeleteProduct(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}
