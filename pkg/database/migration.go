package database

import (
	"gorm.io/gorm"

	"github.com/EDDYCJY/go-gin-example/internal/domain"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&domain.User{},
	)
}