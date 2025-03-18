package repository

import (
	"context"
	"gorm.io/gorm"

	"github.com/EDDYCJY/go-gin-example/internal/domain"
)

type ArticleRepository interface {
	Create(ctx context.Context, article *domain.Article) error
	GetByID(ctx context.Context, id uint) (*domain.Article, error)
	List(ctx context.Context, offset, limit int, tagID *uint) ([]*domain.Article, int64, error)
	Update(ctx context.Context, article *domain.Article) error
	Delete(ctx context.Context, id uint) error
	IncrementViewCount(ctx context.Context, id uint) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

// 实现 ArticleRepository 接口的所有方法...