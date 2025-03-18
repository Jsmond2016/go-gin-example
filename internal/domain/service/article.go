package service

import (
	"context"

	"github.com/EDDYCJY/go-gin-example/internal/domain"
	"github.com/EDDYCJY/go-gin-example/internal/repository"
	"github.com/EDDYCJY/go-gin-example/pkg/errors"
)

type ArticleService interface {
	CreateArticle(ctx context.Context, article *domain.Article) error
	GetArticle(ctx context.Context, id uint) (*domain.Article, error)
	ListArticles(ctx context.Context, page, pageSize int, tagID *uint) ([]*domain.Article, int64, error)
	UpdateArticle(ctx context.Context, article *domain.Article) error
	DeleteArticle(ctx context.Context, id uint) error
	ViewArticle(ctx context.Context, id uint) error
}

type articleService struct {
	articleRepo repository.ArticleRepository
}

func NewArticleService(articleRepo repository.ArticleRepository) ArticleService {
	return &articleService{
		articleRepo: articleRepo,
	}
}

// 实现 ArticleService 接口的所有方法...
func (s *articleService) CreateArticle(ctx context.Context, article *domain.Article) error {
	if err := s.articleRepo.Create(ctx, article); err != nil {
		return err
	}

	// 同步到 Elasticsearch
	_, err := s.esClient.Index().
		Index("articles").
		Id(strconv.FormatUint(uint64(article.ID), 10)).
		BodyJson(article).
		Do(ctx)
	
	if err != nil {
		return errors.Wrap(err, errors.CodeServerError, "同步文章到搜索引擎失败")
	}

	return nil
}