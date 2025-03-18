package tag_service

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/export"
	"github.com/EDDYCJY/go-gin-example/pkg/file"
	"github.com/EDDYCJY/go-gin-example/pkg/gredis"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/service"
	"github.com/EDDYCJY/go-gin-example/service/cache_service"
)

// Tag 服务结构体
type Tag struct {
	service.BaseService
	ID         uint
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int
	PageNum    int
	PageSize   int
}

// NewTagService 创建标签服务
func NewTagService() *Tag {
	return &Tag{
		BaseService: service.NewBaseService(),
		State:       -1, // 默认状态为-1，表示全部
	}
}

// Create 创建标签
func (t *Tag) Create() (uint, error) {
	id, err := models.AddTag(t.Name, t.State, t.CreatedBy)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update 更新标签
func (t *Tag) Update() error {
	updates := map[string]interface{}{
		"modified_by": t.ModifiedBy,
		"name":        t.Name,
	}
	if t.State >= 0 {
		updates["state"] = t.State
	}

	if err := models.EditTag(t.ID, updates); err != nil {
		return err
	}

	// 删除缓存
	cache := cache_service.Tag{
		State:    t.State,
		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := cache.GetTagsKey()
	if gredis.Exists(key) {
		if _, err := gredis.Delete(key); err != nil {
			t.BaseService.HandleCacheError(err, "删除标签缓存失败")
		}
	}

	return nil
}

// GetAll 获取标签列表
func (t *Tag) GetAll() ([]models.Tag, error) {
	var (
		tags, cacheTags []models.Tag
	)

	cache := cache_service.Tag{
		State:    t.State,
		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := cache.GetTagsKey()

	// 尝试从缓存获取
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err == nil {
			if err := json.Unmarshal(data, &cacheTags); err == nil {
				return cacheTags, nil
			}
			t.BaseService.HandleCacheError(err, "解析标签缓存失败")
		}
	}

	// 从数据库获取
	filter := t.getFilter()
	tags, err := models.GetTags(t.PageNum, t.PageSize, filter)
	if err != nil {
		return nil, err
	}

	// 设置缓存，使用 BaseService 的过期时间
	if err := gredis.Set(key, tags, 3600); err != nil {
		t.BaseService.HandleCacheError(err, "设置标签缓存失败")
	}

	return tags, nil
}

// Delete 删除标签
func (t *Tag) Delete() error {
	if err := models.DeleteTag(t.ID); err != nil {
		return err
	}

	// 删除缓存
	cache := cache_service.Tag{
		State:    t.State,
		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := cache.GetTagsKey()
	if gredis.Exists(key) {
		if _, err := gredis.Delete(key); err != nil {
			t.BaseService.HandleCacheError(err, "删除标签缓存失败")
		}
	}

	return nil
}

// ExistByName 检查标签名是否存在
func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

// ExistByID 检查标签ID是否存在
func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

// Count 获取标签总数
func (t *Tag) Count() (int64, error) {
	filter := t.getFilter()
	return models.GetTagTotal(filter)
}

// getFilter 返回查询过滤条件
func (t *Tag) getFilter() models.TagFilter {
	filter := models.TagFilter{
		Name: t.Name,
	}

	if t.State >= 0 {
		state := t.State
		filter.State = &state
	}

	if t.ID > 0 {
		filter.IDs = []uint{t.ID}
	}

	return filter
}

// Export 导出标签
func (t *Tag) Export() (string, error) {
	tags, err := t.GetAll()
	if err != nil {
		return "", fmt.Errorf("获取标签列表失败: %w", err)
	}

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			logging.Error(err)
		}
	}()

	sheet := "标签信息"
	f.NewSheet(sheet)

	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	for i, title := range titles {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, title)
	}

	for i, v := range tags {
		row := i + 2
		values := []interface{}{
			v.ID,
			v.Name,
			v.CreatedBy,
			v.CreatedAt.Unix(),
			v.ModifiedBy,
			v.UpdatedAt.Unix(),
		}

		for j, value := range values {
			cell, _ := excelize.CoordinatesToCellName(j+1, row)
			f.SetCellValue(sheet, cell, value)
		}
	}

	timeStr := strconv.FormatInt(time.Now().Unix(), 10)
	filename := "tags-" + timeStr + export.EXT

	dirFullPath := export.GetExcelFullPath()
	if err = file.IsNotExistMkDir(dirFullPath); err != nil {
		return "", err
	}

	if err = f.SaveAs(dirFullPath + filename); err != nil {
		return "", err
	}

	return filename, nil
}

// Import 导入标签
func (t *Tag) Import(r io.Reader) error {
	if r == nil {
		return fmt.Errorf("无效的输入源")
	}

	f, err := excelize.OpenReader(r)
	if err != nil {
		return fmt.Errorf("打开Excel文件失败: %w", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			logging.Error("关闭Excel文件失败:", err)
		}
	}()

	rows, err := f.GetRows("标签信息")
	if err != nil {
		return err
	}

	for i, row := range rows {
		if i > 0 && len(row) >= 3 {
			t.Name = row[1]
			t.State = 1
			t.CreatedBy = row[2]
			if _, err := t.Create(); err != nil {
				logging.Error("导入标签失败:", err)
				continue
			}
		}
	}

	return nil
}

func (t *Tag) BatchDelete(ids []uint) error {
	if len(ids) == 0 {
		return nil
	}

	// 批量删除标签
	if err := models.BatchEditTags(ids, map[string]interface{}{"deleted_at": time.Now()}); err != nil {
		return err
	}

	// 删除缓存
	cache := cache_service.Tag{
		State:    t.State,
		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := cache.GetTagsKey()
	if gredis.Exists(key) {
		if _, err := gredis.Delete(key); err != nil {
			t.BaseService.HandleCacheError(err, "删除标签缓存失败")
		}
	}

	return nil
}
