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
	List(page, pageSize int, categoryID *uuid.UUID) ([]*domain.Product, int64, error)
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

func (s *productService) List(page, pageSize int, categoryID *uuid.UUID) ([]*domain.Product, int64, error) {
	filters := make(map[string]interface{})
	if categoryID != nil {
		filters["category_id"] = *categoryID
	}
	return s.productRepo.List(page, pageSize, filters)
}
