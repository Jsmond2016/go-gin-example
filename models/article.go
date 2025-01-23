package models

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

// ExistArticleByID checks if an article exists based on ID
func ExistArticleByID(id uint) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetArticleTotal gets the total number of articles based on the constraints
func GetArticleTotal(maps interface{}) (int64, error) {
	var count int64
	err := db.Model(&Article{}).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetArticles gets a list of articles based on paging constraints
func GetArticles(pageNum int, pageSize int, maps interface{}) ([]*Article, error) {
	var articles []*Article
	query := db.Model(&Article{}).Preload("Tag").Where(maps)

	if pageSize > 0 && pageNum > 0 {
		query = query.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}

	err := query.Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

// GetArticle Get a single article based on ID
func GetArticle(id uint) (*Article, error) {
	var article Article
	err := db.Where("id = ?", id).Preload("Tag").First(&article).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// EditArticle modify a single article
func EditArticle(id uint, data interface{}) error {
	return db.Model(&Article{}).Where("id = ?", id).Updates(data).Error
}

// AddArticle add a single article
func AddArticle(data map[string]interface{}) error {
	article := Article{
		TagID:         data["tag_id"].(uint),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	}
	return db.Create(&article).Error
}

// DeleteArticle delete a single article
func DeleteArticle(id uint) error {
	return db.Where("id = ?", id).Delete(&Article{}).Error
}

// CleanAllArticle clear all soft deleted articles
func CleanAllArticle() error {
	return db.Unscoped().Where("deleted_at IS NOT NULL").Delete(&Article{}).Error
}

// AddArticle2 添加文章（新方法，直接接收结构体）
func AddArticle2(article *Article) error {
	return db.Create(article).Error
}
