package models

import (
	"errors"
	"gorm.io/gorm"
)

// Article 文章模型
type Article struct {
	Model

	TagID uint `json:"tag_id" gorm:"index"`
	Tag   Tag  `json:"tag"`

	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
}

// ArticleFilter 文章查询过滤器
type ArticleFilter struct {
	TagID     *uint
	Title     string
	State     *int
	CreatedBy string
}

// ExistArticleByID 检查文章是否存在
func ExistArticleByID(id uint) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// GetArticleTotal 获取文章总数
func GetArticleTotal(filter ArticleFilter) (int64, error) {
	var count int64
	query := db.Model(&Article{})
	
	query = applyFilter(query, filter)
	
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetArticles 获取文章列表
func GetArticles(pageNum int, pageSize int, filter ArticleFilter) ([]*Article, error) {
	var articles []*Article
	query := db.Model(&Article{}).Preload("Tag")
	
	query = applyFilter(query, filter)

	if pageSize > 0 && pageNum > 0 {
		query = query.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}

	err := query.Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

// GetArticle 获取单个文章
func GetArticle(id uint) (*Article, error) {
	var article Article
	err := db.Where("id = ?", id).Preload("Tag").First(&article).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("article not found")
		}
		return nil, err
	}
	return &article, nil
}

// CreateArticle 创建文章
func CreateArticle(article *Article) error {
	return db.Create(article).Error
}

// CreateArticleTx 在事务中创建文章
func CreateArticleTx(tx *gorm.DB, article *Article) error {
	if tx == nil {
		tx = db
	}
	return tx.Create(article).Error
}

// UpdateArticle 更新文章
func UpdateArticle(id uint, updates map[string]interface{}) error {
	result := db.Model(&Article{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("article not found")
	}
	return nil
}

// DeleteArticle 删除文章
func DeleteArticle(id uint) error {
	result := db.Where("id = ?", id).Delete(&Article{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("article not found")
	}
	return nil
}

// CleanAllArticle 清理所有软删除的文章
func CleanAllArticle() error {
	return db.Unscoped().Where("deleted_at IS NOT NULL").Delete(&Article{}).Error
}

// applyFilter 应用查询过滤器
func applyFilter(query *gorm.DB, filter ArticleFilter) *gorm.DB {
	if filter.TagID != nil {
		query = query.Where("tag_id = ?", *filter.TagID)
	}
	
	if filter.Title != "" {
		query = query.Where("title LIKE ?", "%"+filter.Title+"%")
	}
	
	if filter.State != nil {
		query = query.Where("state = ?", *filter.State)
	}
	
	if filter.CreatedBy != "" {
		query = query.Where("created_by = ?", filter.CreatedBy)
	}
	
	return query
}
