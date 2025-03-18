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

// TagFilter 定义查询过滤条件
type TagFilter struct {
	Name  string
	State *int
	IDs   []uint
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

// ExistTagByID determines whether a Tag exists based on the ID
func ExistTagByID(id uint) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ?", id).First(&tag).Error
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

// AddTagTx 在事务中添加标签
func AddTagTx(tx *gorm.DB, name string, state int, createdBy string) error {
	if tx == nil {
		tx = db
	}

	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	return tx.Create(&tag).Error
}

// GetTags gets a list of tags based on paging and constraints
func GetTags(pageNum int, pageSize int, filter TagFilter) ([]Tag, error) {
	var tags []Tag
	query := db.Model(&Tag{})

	// 应用过滤条件
	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}

	if filter.State != nil {
		query = query.Where("state = ?", *filter.State)
	}

	if len(filter.IDs) > 0 {
		query = query.Where("id IN ?", filter.IDs)
	}

	// 分页
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
func GetTagTotal(filter TagFilter) (int64, error) {
	var count int64
	query := db.Model(&Tag{})

	// 应用过滤条件
	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}

	if filter.State != nil {
		query = query.Where("state = ?", *filter.State)
	}

	if len(filter.IDs) > 0 {
		query = query.Where("id IN ?", filter.IDs)
	}

	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

// DeleteTag delete a tag
func DeleteTag(id uint) error {
	return db.Where("id = ?", id).Delete(&Tag{}).Error
}

// DeleteTagTx 在事务中删除标签
func DeleteTagTx(tx *gorm.DB, id uint) error {
	if tx == nil {
		tx = db
	}
	return tx.Where("id = ?", id).Delete(&Tag{}).Error
}

// EditTag modify a single tag
func EditTag(id uint, data interface{}) error {
	return db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error
}

// BatchEditTags 批量更新标签
func BatchEditTags(ids []uint, data interface{}) error {
	return db.Model(&Tag{}).Where("id IN ?", ids).Updates(data).Error
}

// CleanAllTag clear all soft deleted tags
func CleanAllTag() (bool, error) {
	err := db.Unscoped().Where("deleted_at IS NOT NULL").Delete(&Tag{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
