package service

import (
	"context"

	"github.com/EDDYCJY/go-gin-example/internal/domain"
	"github.com/EDDYCJY/go-gin-example/internal/repository"
	"github.com/EDDYCJY/go-gin-example/pkg/errors"
)

type CommentService interface {
	CreateComment(ctx context.Context, comment *domain.Comment) error
	GetArticleComments(ctx context.Context, articleID uint) ([]*domain.Comment, error)
	DeleteComment(ctx context.Context, id uint, userID uint) error
}

type commentService struct {
	commentRepo repository.CommentRepository
}

func NewCommentService(commentRepo repository.CommentRepository) CommentService {
	return &commentService{
		commentRepo: commentRepo,
	}
}

// 实现 CommentService 接口的所有方法...