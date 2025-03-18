package domain

import "time"

type Comment struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	ArticleID uint      `json:"article_id"`
	UserID    uint      `json:"user_id"`
	Content   string    `json:"content" gorm:"type:text;not null"`
	ParentID  *uint     `json:"parent_id"`
	State     int       `json:"state" gorm:"default:1"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}