package service

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"github.com/unmanned-store/backend/internal/repository"
)

type CategoryService interface {
	Create(category *domain.Category) error
	GetByID(id uuid.UUID) (*domain.Category, error)
	Update(category *domain.Category) error
	Delete(id uuid.UUID) error
	List() ([]*domain.Category, error)
	GetTree() ([]*domain.Category, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (s *categoryService) Create(category *domain.Category) error {
	return s.categoryRepo.Create(category)
}

func (s *categoryService) GetByID(id uuid.UUID) (*domain.Category, error) {
	return s.categoryRepo.GetByID(id)
}

func (s *categoryService) Update(category *domain.Category) error {
	return s.categoryRepo.Update(category)
}

func (s *categoryService) Delete(id uuid.UUID) error {
	return s.categoryRepo.Delete(id)
}

func (s *categoryService) List() ([]*domain.Category, error) {
	return s.categoryRepo.List()
}

func (s *categoryService) GetTree() ([]*domain.Category, error) {
	return s.categoryRepo.GetTree()
}
