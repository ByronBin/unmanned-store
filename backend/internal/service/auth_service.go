package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/unmanned-store/backend/internal/domain"
	"github.com/unmanned-store/backend/internal/middleware"
	"github.com/unmanned-store/backend/internal/repository"
	"github.com/unmanned-store/backend/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (string, string, error)
	Register(user *domain.User, password string) error
	RefreshToken(refreshToken string) (string, error)
	Logout(userID uuid.UUID) error
	VerifyToken(token string) (*middleware.Claims, error)
}

type authService struct {
	userRepo repository.UserRepository
	rdb      *redis.Client
	cfg      *config.Config
}

func NewAuthService(userRepo repository.UserRepository, rdb *redis.Client, cfg *config.Config) AuthService {
	return &authService{
		userRepo: userRepo,
		rdb:      rdb,
		cfg:      cfg,
	}
}

func (s *authService) Login(username, password string) (string, string, error) {
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return "", "", errors.New("用户名或密码错误")
	}

	if user.Status != "active" {
		return "", "", errors.New("账户已被禁用")
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", errors.New("用户名或密码错误")
	}

	// 生成访问令牌
	accessToken, err := s.generateToken(user, time.Duration(s.cfg.JWT.ExpireHours)*time.Hour)
	if err != nil {
		return "", "", err
	}

	// 生成刷新令牌
	refreshToken, err := s.generateToken(user, time.Duration(s.cfg.JWT.RefreshExpireHours)*time.Hour)
	if err != nil {
		return "", "", err
	}

	// 将刷新令牌存储到Redis
	ctx := context.Background()
	key := fmt.Sprintf("refresh_token:%s", user.ID)
	s.rdb.Set(ctx, key, refreshToken, time.Duration(s.cfg.JWT.RefreshExpireHours)*time.Hour)

	return accessToken, refreshToken, nil
}

func (s *authService) Register(user *domain.User, password string) error {
	// 检查用户名是否已存在
	if _, err := s.userRepo.GetByUsername(user.Username); err == nil {
		return errors.New("用户名已存在")
	}

	// 检查手机号是否已存在
	if user.Phone != "" {
		if _, err := s.userRepo.GetByPhone(user.Phone); err == nil {
			return errors.New("手机号已被注册")
		}
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// 设置默认值
	if user.Role == "" {
		user.Role = "customer"
	}
	user.Status = "active"

	return s.userRepo.Create(user)
}

func (s *authService) RefreshToken(refreshToken string) (string, error) {
	claims, err := s.VerifyToken(refreshToken)
	if err != nil {
		return "", err
	}

	// 验证Redis中的刷新令牌
	ctx := context.Background()
	key := fmt.Sprintf("refresh_token:%s", claims.UserID)
	storedToken, err := s.rdb.Get(ctx, key).Result()
	if err != nil || storedToken != refreshToken {
		return "", errors.New("刷新令牌无效")
	}

	// 获取用户信息
	user, err := s.userRepo.GetByID(claims.UserID)
	if err != nil {
		return "", err
	}

	// 生成新的访问令牌
	accessToken, err := s.generateToken(user, time.Duration(s.cfg.JWT.ExpireHours)*time.Hour)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (s *authService) Logout(userID uuid.UUID) error {
	ctx := context.Background()
	key := fmt.Sprintf("refresh_token:%s", userID)
	return s.rdb.Del(ctx, key).Err()
}

func (s *authService) VerifyToken(tokenString string) (*middleware.Claims, error) {
	claims := &middleware.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.Secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("令牌无效")
	}

	return claims, nil
}

func (s *authService) generateToken(user *domain.User, expireDuration time.Duration) (string, error) {
	now := time.Now()
	claims := middleware.Claims{
		UserID:  user.ID,
		Role:    user.Role,
		StoreID: user.StoreID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(expireDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.JWT.Secret))
}
