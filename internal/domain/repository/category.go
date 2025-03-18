package repository

import (
	"context"
	"gorm.io/gorm"

	"github.com/EDDYCJY/go-gin-example/internal/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, category *domain.Category) error
	GetByID(ctx context.Context, id uint) (*domain.Category, error)
	List(ctx context.Context, parentID *uint) ([]*domain.Category, error)
	Update(ctx context.Context, category *domain.Category) error
	Delete(ctx context.Context, id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

// 实现所有接口方法...