package repository

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *domain.Category) error
	GetByID(id uuid.UUID) (*domain.Category, error)
	Update(category *domain.Category) error
	Delete(id uuid.UUID) error
	List() ([]*domain.Category, error)
	GetTree() ([]*domain.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *domain.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) GetByID(id uuid.UUID) (*domain.Category, error) {
	var category domain.Category
	err := r.db.Preload("Parent").First(&category, "id = ? AND deleted_at IS NULL", id).Error
	return &category, err
}

func (r *categoryRepository) Update(category *domain.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.Category{}, "id = ?", id).Error
}

func (r *categoryRepository) List() ([]*domain.Category, error) {
	var categories []*domain.Category
	err := r.db.Where("deleted_at IS NULL").Order("sort ASC").Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) GetTree() ([]*domain.Category, error) {
	var categories []*domain.Category
	err := r.db.Where("deleted_at IS NULL AND parent_id IS NULL").Order("sort ASC").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	// 递归加载子分类
	for _, cat := range categories {
		r.loadChildren(cat)
	}

	return categories, nil
}

func (r *categoryRepository) loadChildren(parent *domain.Category) {
	var children []*domain.Category
	r.db.Where("parent_id = ? AND deleted_at IS NULL", parent.ID).Order("sort ASC").Find(&children)

	for _, child := range children {
		r.loadChildren(child)
	}
}
