package repository

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *domain.Product) error
	GetByID(id uuid.UUID) (*domain.Product, error)
	Update(product *domain.Product) error
	Delete(id uuid.UUID) error
	List(page, pageSize int, filters map[string]interface{}) ([]*domain.Product, int64, error)
	CreateSKU(sku *domain.ProductSKU) error
	GetSKUByID(id uuid.UUID) (*domain.ProductSKU, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *domain.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetByID(id uuid.UUID) (*domain.Product, error) {
	var product domain.Product
	err := r.db.Preload("Category").Preload("SKUs").First(&product, "id = ? AND deleted_at IS NULL", id).Error
	return &product, err
}

func (r *productRepository) Update(product *domain.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Product{}, "id = ?", id).Error
}

func (r *productRepository) List(page, pageSize int, filters map[string]interface{}) ([]*domain.Product, int64, error) {
	var products []*domain.Product
	var total int64

	query := r.db.Model(&domain.Product{}).Where("deleted_at IS NULL")

	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Preload("Category").Offset(offset).Limit(pageSize).Find(&products).Error

	return products, total, err
}

func (r *productRepository) CreateSKU(sku *domain.ProductSKU) error {
	return r.db.Create(sku).Error
}

func (r *productRepository) GetSKUByID(id uuid.UUID) (*domain.ProductSKU, error) {
	var sku domain.ProductSKU
	err := r.db.Preload("Product").First(&sku, "id = ? AND deleted_at IS NULL", id).Error
	return &sku, err
}
