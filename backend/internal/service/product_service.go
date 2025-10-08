package service

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"github.com/unmanned-store/backend/internal/repository"
)

type ProductService interface {
	Create(product *domain.Product) error
	GetByID(id uuid.UUID) (*domain.Product, error)
	Update(product *domain.Product) error
	Delete(id uuid.UUID) error
	List(page, pageSize int, categoryID *uuid.UUID, keyword string, status string) ([]*domain.Product, int64, error)
	UpdateStatus(id uuid.UUID, status string) error
	GetHotProducts(limit int) ([]*domain.Product, error)
	SearchProducts(keyword string, page, pageSize int) ([]*domain.Product, int64, error)
	
	// SKU管理
	CreateSKU(sku *domain.ProductSKU) error
	UpdateSKU(sku *domain.ProductSKU) error
	DeleteSKU(id uuid.UUID) error
	GetSKUByID(id uuid.UUID) (*domain.ProductSKU, error)
}

type productService struct {
	productRepo  repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

func NewProductService(productRepo repository.ProductRepository, categoryRepo repository.CategoryRepository) ProductService {
	return &productService{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *productService) Create(product *domain.Product) error {
	return s.productRepo.Create(product)
}

func (s *productService) GetByID(id uuid.UUID) (*domain.Product, error) {
	return s.productRepo.GetByID(id)
}

func (s *productService) Update(product *domain.Product) error {
	return s.productRepo.Update(product)
}

func (s *productService) Delete(id uuid.UUID) error {
	return s.productRepo.Delete(id)
}

func (s *productService) List(page, pageSize int, categoryID *uuid.UUID, keyword string, status string) ([]*domain.Product, int64, error) {
	filters := make(map[string]interface{})
	if categoryID != nil {
		filters["category_id"] = *categoryID
	}
	if status != "" {
		filters["status"] = status
	}
	
	// 如果有关键词搜索，需要在repository层处理
	if keyword != "" {
		filters["keyword"] = keyword
	}
	
	return s.productRepo.List(page, pageSize, filters)
}

func (s *productService) UpdateStatus(id uuid.UUID, status string) error {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return err
	}
	
	product.Status = status
	return s.productRepo.Update(product)
}

func (s *productService) GetHotProducts(limit int) ([]*domain.Product, error) {
	// 这里可以根据销量、浏览量等计算热销商品
	// 暂时返回最新的商品
	return s.productRepo.GetHotProducts(limit)
}

func (s *productService) SearchProducts(keyword string, page, pageSize int) ([]*domain.Product, int64, error) {
	filters := map[string]interface{}{
		"keyword": keyword,
	}
	return s.productRepo.List(page, pageSize, filters)
}

// SKU管理方法
func (s *productService) CreateSKU(sku *domain.ProductSKU) error {
	return s.productRepo.CreateSKU(sku)
}

func (s *productService) UpdateSKU(sku *domain.ProductSKU) error {
	return s.productRepo.UpdateSKU(sku)
}

func (s *productService) DeleteSKU(id uuid.UUID) error {
	return s.productRepo.DeleteSKU(id)
}

func (s *productService) GetSKUByID(id uuid.UUID) (*domain.ProductSKU, error) {
	return s.productRepo.GetSKUByID(id)
}
