package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"

	"github.com/EDDYCJY/go-gin-example/internal/domain"
	"github.com/EDDYCJY/go-gin-example/internal/repository"
	"github.com/EDDYCJY/go-gin-example/pkg/errors"
	"github.com/EDDYCJY/go-gin-example/pkg/jwt"
)

type UserService interface {
	Register(ctx context.Context, username, password, email string) error
	Login(ctx context.Context, username, password string) (string, error)
	GetUserInfo(ctx context.Context, id uint) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(ctx context.Context, username, password, email string) error {
	// 检查用户名是否已存在
	if _, err := s.userRepo.GetByUsername(ctx, username); err == nil {
		return errors.New(errors.CodeInvalidParams, "用户名已存在")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, errors.CodeServerError, "密码加密失败")
	}

	user := &domain.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return errors.Wrap(err, errors.CodeServerError, "创建用户失败")
	}

	return nil
}

func (s *userService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", errors.New(errors.CodeUnauthorized, "用户名或密码错误")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New(errors.CodeUnauthorized, "用户名或密码错误")
	}

	// 生成 JWT token
	token, err := jwt.GenerateToken(username)
	if err != nil {
		return "", errors.Wrap(err, errors.CodeServerError, "生成token失败")
	}

	return token, nil
}

func (s *userService) GetUserInfo(ctx context.Context, id uint) (*domain.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, errors.CodeNotFound, "用户不存在")
	}
	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, user *domain.User) error {
	if err := s.userRepo.Update(ctx, user); err != nil {
		return errors.Wrap(err, errors.CodeServerError, "更新用户信息失败")
	}
	return nil
}