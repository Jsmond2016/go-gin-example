package domain

import "time"

type Tag struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"size:50;not null;unique"`
	Description string    `json:"description" gorm:"size:200"`
	CreatedBy   string    `json:"created_by"`
	ModifiedBy  string    `json:"modified_by"`
	State       int       `json:"state" gorm:"default:1"`
	Count       int       `json:"count" gorm:"default:0"` // 使用次数
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}