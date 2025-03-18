package service

import (
	"context"

	"github.com/EDDYCJY/go-gin-example/internal/domain"
	"github.com/EDDYCJY/go-gin-example/internal/repository"
	"github.com/EDDYCJY/go-gin-example/pkg/errors"
)

type CategoryService interface {
	CreateCategory(ctx context.Context, category *domain.Category) error
	GetCategory(ctx context.Context, id uint) (*domain.Category, error)
	ListCategories(ctx context.Context, parentID *uint) ([]*domain.Category, error)
	UpdateCategory(ctx context.Context, category *domain.Category) error
	DeleteCategory(ctx context.Context, id uint) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
	articleRepo  repository.ArticleRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository, articleRepo repository.ArticleRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
		articleRepo:  articleRepo,
	}
}

// 实现所有接口方法...