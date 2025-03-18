package tag_service

import (
	"encoding/json"
	"io"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"

	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/export"
	"github.com/EDDYCJY/go-gin-example/pkg/file"
	"github.com/EDDYCJY/go-gin-example/pkg/gredis"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/service/cache_service"
)

type Tag struct {
	ID         uint
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

func (t *Tag) ExistByID() (bool, error) {
	return models.ExistTagByID(t.ID)
}

func (t *Tag) Add() error {
	return models.AddTag(t.Name, t.State, t.CreatedBy)
}

func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}

	return models.EditTag(t.ID, data)
}

func (t *Tag) Delete() error {
	return models.DeleteTag(t.ID)
}

func (t *Tag) Count() (int64, error) {
	filter := t.getFilter()
	return models.GetTagTotal(filter)
}

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
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheTags)
			return cacheTags, nil
		}
	}

	filter := t.getFilter()
	tags, err := models.GetTags(t.PageNum, t.PageSize, filter)
	if err != nil {
		return nil, err
	}

	gredis.Set(key, tags, 3600)
	return tags, nil
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

func (t *Tag) Export() (string, error) {
	tags, err := t.GetAll()
	if err != nil {
		return "", err
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

	time := strconv.FormatInt(time.Now().Unix(), 10)
	filename := "tags-" + time + export.EXT

	dirFullPath := export.GetExcelFullPath()
	err = file.IsNotExistMkDir(dirFullPath)
	if err != nil {
		return "", err
	}

	err = f.SaveAs(dirFullPath + filename)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func (t *Tag) Import(r io.Reader) error {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			logging.Error(err)
		}
	}()

	rows, err := f.GetRows("标签信息")
	if err != nil {
		return err
	}

	for i, row := range rows {
		if i > 0 && len(row) >= 3 {
			models.AddTag(row[1], 1, row[2])
		}
	}

	return nil
}
