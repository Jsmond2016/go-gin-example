package domain

import "time"

type Article struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	Title         string    `json:"title" gorm:"size:100;not null"`
	Content       string    `json:"content" gorm:"type:text;not null"`
	Description   string    `json:"description" gorm:"size:255"`
	CoverImageURL string    `json:"cover_image_url"`
	State        int       `json:"state" gorm:"default:1"`
	ViewCount    int       `json:"view_count" gorm:"default:0"`
	TagID        uint      `json:"tag_id"`
	CategoryID uint   `json:"category_id"`
	Tags      []Tag  `json:"tags" gorm:"many2many:article_tags;"`
	Category  Category `json:"category" gorm:"foreignKey:CategoryID"`
	CreatedBy    string    `json:"created_by"`
	ModifiedBy   string    `json:"modified_by"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}