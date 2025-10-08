package repository

import (
	"github.com/google/uuid"
	"github.com/unmanned-store/backend/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id uuid.UUID) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	GetByPhone(phone string) (*domain.User, error)
	GetByWechatOpenID(openID string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id uuid.UUID) error
	List(page, pageSize int, filters map[string]interface{}) ([]*domain.User, int64, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uuid.UUID) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("Store").First(&user, "id = ? AND deleted_at IS NULL", id).Error
	return &user, err
}

func (r *userRepository) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("username = ? AND deleted_at IS NULL", username).First(&user).Error
	return &user, err
}

func (r *userRepository) GetByPhone(phone string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("phone = ? AND deleted_at IS NULL", phone).First(&user).Error
	return &user, err
}

func (r *userRepository) GetByWechatOpenID(openID string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("wechat_openid = ? AND deleted_at IS NULL", openID).First(&user).Error
	return &user, err
}

func (r *userRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&domain.User{}, "id = ?", id).Error
}

func (r *userRepository) List(page, pageSize int, filters map[string]interface{}) ([]*domain.User, int64, error) {
	var users []*domain.User
	var total int64

	query := r.db.Model(&domain.User{}).Where("deleted_at IS NULL")

	// 应用过滤条件
	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	// 计算总数
	query.Count(&total)

	// 分页查询
	offset := (page - 1) * pageSize
	err := query.Preload("Store").Offset(offset).Limit(pageSize).Find(&users).Error

	return users, total, err
}
