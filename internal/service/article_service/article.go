package article_service

import (
	"encoding/json"
	"fmt"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/gredis"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/service/cache_service"
)

type Article struct {
	ID            uint
	TagID         uint
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

func (a *Article) Add() error {
	article := models.Article{
		TagID:         a.TagID,
		Title:         a.Title,
		Desc:          a.Desc,
		Content:       a.Content,
		CreatedBy:     a.CreatedBy,
		State:         a.State,
		CoverImageUrl: a.CoverImageUrl,
	}
	return models.AddArticle2(&article)
}

func (a *Article) Edit() error {
	return models.EditArticle(a.ID, map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
		"modified_by":     a.ModifiedBy,
	})
}

func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}

	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, article, 3600)
	return article, nil
}

func (a *Article) GetAll() ([]*models.Article, error) {
	var (
		articles, cacheArticles []*models.Article
	)

	cache := cache_service.Article{
		TagID: a.TagID,
		State: a.State,

		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetArticlesKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticles)
			return cacheArticles, nil
		}
	}

	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, articles, 3600)
	return articles, nil
}

func (a *Article) Delete() error {
	return models.DeleteArticle(a.ID)
}

func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}

func (a *Article) Count() (int64, error) {
	return models.GetArticleTotal(a.getMaps())
}

func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if a.State != -1 {
		maps["state"] = a.State
	}
	if a.TagID != 0 {
		maps["tag_id"] = a.TagID
	}

	return maps
}

// GetQrCodeUrl returns the QR code URL for the article
func (a *Article) GetQrCodeUrl() string {
	return fmt.Sprintf("%s/api/v1/articles/%d", setting.AppSetting.PrefixUrl, a.ID)
}
