package service

import (
	"context"
	"encoding/json"

	"github.com/olivere/elastic/v7"

	"github.com/EDDYCJY/go-gin-example/internal/domain"
	"github.com/EDDYCJY/go-gin-example/pkg/errors"
)

type SearchService interface {
	SearchArticles(ctx context.Context, keyword string, page, pageSize int) ([]*domain.Article, int64, error)
}

type searchService struct {
	esClient    *elastic.Client
	articleRepo repository.ArticleRepository
}

func NewSearchService(esClient *elastic.Client, articleRepo repository.ArticleRepository) SearchService {
	return &searchService{
		esClient:    esClient,
		articleRepo: articleRepo,
	}
}

func (s *searchService) SearchArticles(ctx context.Context, keyword string, page, pageSize int) ([]*domain.Article, int64, error) {
	query := elastic.NewMultiMatchQuery(keyword, "title", "content", "description")
	searchResult, err := s.esClient.Search().
		Index("articles").
		Query(query).
		From((page - 1) * pageSize).
		Size(pageSize).
		Do(ctx)
	
	if err != nil {
		return nil, 0, errors.Wrap(err, errors.CodeServerError, "搜索文章失败")
	}

	var articles []*domain.Article
	for _, hit := range searchResult.Hits.Hits {
		var article domain.Article
		if err := json.Unmarshal(hit.Source, &article); err != nil {
			return nil, 0, errors.Wrap(err, errors.CodeServerError, "解析搜索结果失败")
		}
		articles = append(articles, &article)
	}

	return articles, searchResult.TotalHits(), nil
}