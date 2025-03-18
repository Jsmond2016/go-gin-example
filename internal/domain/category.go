package domain

import "time"

type Category struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"size:50;not null;unique"`
	Description string    `json:"description" gorm:"size:200"`
	ParentID    *uint     `json:"parent_id" gorm:"default:null"`
	Level       int       `json:"level" gorm:"default:1"`
	Sort        int       `json:"sort" gorm:"default:0"`
	CreatedBy   string    `json:"created_by"`
	ModifiedBy  string    `json:"modified_by"`
	State       int       `json:"state" gorm:"default:1"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}