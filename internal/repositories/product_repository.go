package repositories

import (
	"fmt"
	"local_my_api/internal/models"
	"local_my_api/pkg/utils"
	"log"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error)
	GetProductList(pagination utils.Pagination, businessID string) ([]models.Product, int64, error)
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

	return &product, r.db.Create(&product).Error
}

func (r *productRepository) GetProductList(pagination utils.Pagination, businessID string) ([]models.Product, int64, error) {
	productList := []models.Product{}
	totalProduct := int64(0)

	query := r.db.Offset(pagination.Offset).Limit(pagination.Size)

	if businessID != "" {
		query = query.Where(&models.Product{BusinessID: businessID})
	}

	if err := query.Preload(models.PreloadBusinessOwner).Find(&productList).Error; err != nil {
		return nil, totalProduct, fmt.Errorf("failed to retrieve product list: %w", err)
	}

	countQuery := r.db.Model(&models.Product{})
	if businessID != "" {
		countQuery = countQuery.Where(&models.Product{BusinessID: businessID})
	}

	if err := countQuery.Count(&totalProduct).Error; err != nil {
		return nil, totalProduct, fmt.Errorf("failed to count products: %w", err)
	}

	return productList, totalProduct, nil
}

func (r *productRepository) GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	log.Printf("id: %s", id)
	if err := r.db.Preload("Business").First(&product, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve product: %w", err)
	}

	return &product, nil
}

func (r *productRepository) UpdateProduct(id string, p *models.ProductUpdate) (*models.Product, error) {
	product, err := r.GetProductByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve product: %w", err)
	}

	if err := r.db.Model(product).Updates(p).Error; err != nil {
		return nil, fmt.Errorf("failed to update product: %w", err)
	}

	return product, nil
}

func (r *productRepository) DeleteProduct(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Product{}).Error; err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}
