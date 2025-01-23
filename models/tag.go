package models

import (
	"errors"

	"gorm.io/gorm"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// ExistTagByName checks if there is a tag with the same name
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil // 记录不存在，返回 false 且无错误
		}
		return false, err // 其他错误则返回错误
	}
	return true, nil // 找到记录，返回 true
}

// AddTag Add a Tag
func AddTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	return db.Create(&tag).Error
}

// GetTags gets a list of tags based on paging and constraints
func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var tags []Tag
	query := db.Model(&Tag{}).Where(maps)

	if pageSize > 0 && pageNum > 0 {
		query = query.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}

	err := query.Find(&tags).Error
	if err != nil {
		return nil, err
	}

	return tags, nil
}

// GetTagTotal counts the total number of tags based on the constraint
func GetTagTotal(maps interface{}) (int64, error) {
	var count int64
	err := db.Model(&Tag{}).Where(maps).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// ExistTagByID determines whether a Tag exists based on the ID
func ExistTagByID(id uint) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ?", id).First(&tag).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteTag delete a tag
func DeleteTag(id uint) error {
	return db.Where("id = ?", id).Delete(&Tag{}).Error
}

// EditTag modify a single tag
func EditTag(id uint, data interface{}) error {
	return db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error
}

// CleanAllTag clear all soft deleted tags
func CleanAllTag() (bool, error) {
	err := db.Unscoped().Where("deleted_at IS NOT NULL").Delete(&Tag{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
