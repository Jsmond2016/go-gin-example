package article_service

import (
	"encoding/json"
	"fmt"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/gredis"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/service"
	"github.com/EDDYCJY/go-gin-example/service/cache_service"
)

// Article 服务结构体
type Article struct {
	service.BaseService
	ID            uint
	TagID         uint
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string
	PageNum       int
	PageSize      int
}

// NewArticleService 创建文章服务
func NewArticleService() *Article {
	return &Article{
		BaseService: service.NewBaseService(),
	}
}

// Get 获取单个文章
func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()

	// 尝试从缓存获取
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err == nil {
			if err := json.Unmarshal(data, &cacheArticle); err == nil {
				return cacheArticle, nil
			}
			a.BaseService.HandleCacheError(err, "解析文章缓存失败")
		}
	}

	// 从数据库获取
	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}

	// 设置缓存
	if err := gredis.Set(key, article, a.CacheExpire); err != nil {
		a.BaseService.HandleCacheError(err, "设置文章缓存失败")
	}

	return article, nil
}

// GetAll 获取文章列表
func (a *Article) GetAll() ([]*models.Article, error) {
	var (
		articles, cacheArticles []*models.Article
	)

	cache := cache_service.Article{
		TagID:    a.TagID,
		State:    a.State,
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := cache.GetArticlesKey()

	// 尝试从缓存获取
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err == nil {
			if err := json.Unmarshal(data, &cacheArticles); err == nil {
				return cacheArticles, nil
			}
			a.BaseService.HandleCacheError(err, "解析文章列表缓存失败")
		}
	}

	// 从数据库获取
	filter := a.getFilter()
	articles, err := models.GetArticles(a.PageNum, a.PageSize, filter)
	if err != nil {
		return nil, err
	}

	// 设置缓存
	if err := gredis.Set(key, articles, a.CacheExpire); err != nil {
		a.BaseService.HandleCacheError(err, "设置文章列表缓存失败")
	}

	return articles, nil
}

// Create 创建文章
func (a *Article) Create() (uint, error) {
	article := &models.Article{
		TagID:         a.TagID,
		Title:         a.Title,
		Desc:          a.Desc,
		Content:       a.Content,
		CreatedBy:     a.CreatedBy,
		State:         a.State,
		CoverImageUrl: a.CoverImageUrl,
	}

	articleID, err := models.CreateArticle(article)
	if err != nil {
		return 0, err
	}

	// 删除列表缓存
	listCache := cache_service.Article{
		TagID:    a.TagID,
		State:    a.State,
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	key := listCache.GetArticlesKey()
	if gredis.Exists(key) {
		if _, err := gredis.Delete(key); err != nil {
			a.BaseService.HandleCacheError(err, "删除文章列表缓存失败")
		}
	}

	return articleID, nil
}

// Update 更新文章
func (a *Article) Update() error {
	updates := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
		"modified_by":     a.ModifiedBy,
	}

	if err := models.UpdateArticle(a.ID, updates); err != nil {
		return err
	}

	// 删除缓存
	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		if _, err := gredis.Delete(key); err != nil {
			a.BaseService.HandleCacheError(err, "删除文章缓存失败")
		}
	}

	return nil
}

// Delete 删除文章
func (a *Article) Delete() error {
	if err := models.DeleteArticle(a.ID); err != nil {
		return err
	}

	// 删除缓存
	cache := cache_service.Article{ID: a.ID}
	key := cache.GetArticleKey()
	if gredis.Exists(key) {
		if _, err := gredis.Delete(key); err != nil {
			a.BaseService.HandleCacheError(err, "删除文章缓存失败")
		}
	}

	// 删除列表缓存
	listCache := cache_service.Article{
		TagID:    a.TagID,
		State:    a.State,
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	listKey := listCache.GetArticlesKey()
	if gredis.Exists(listKey) {
		if _, err := gredis.Delete(listKey); err != nil {
			a.BaseService.HandleCacheError(err, "删除文章列表缓存失败")
		}
	}

	return nil
}

// ExistByID 检查文章是否存在
func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}

// Count 获取文章总数
func (a *Article) Count() (int64, error) {
	return models.GetArticleTotal(a.getFilter())
}

// getFilter 获取查询过滤器
func (a *Article) getFilter() models.ArticleFilter {
	filter := models.ArticleFilter{
		CreatedBy: a.CreatedBy,
	}

	if a.State != -1 {
		state := a.State
		filter.State = &state
	}

	if a.TagID != 0 {
		filter.TagID = &a.TagID
	}

	if a.Title != "" {
		filter.Title = a.Title
	}

	return filter
}

// GetQrCodeUrl 获取文章二维码URL
func (a *Article) GetQrCodeUrl() string {
	return fmt.Sprintf("%s/api/v1/articles/%d", setting.AppSetting.PrefixUrl, a.ID)
}
