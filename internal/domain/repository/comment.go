package repository

import (
	"context"
	"gorm.io/gorm"

	"github.com/EDDYCJY/go-gin-example/internal/domain"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *domain.Comment) error
	GetByArticleID(ctx context.Context, articleID uint) ([]*domain.Comment, error)
	Delete(ctx context.Context, id uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

// 实现 CommentRepository 接口的所有方法...