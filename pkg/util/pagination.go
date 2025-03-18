package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

// Pagination 分页参数
type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Offset   int `json:"offset"`
}

// GetPagination 获取分页参数
func GetPagination(c *gin.Context) Pagination {
	page := com.StrTo(c.Query("page")).MustInt()
	pageSize := com.StrTo(c.Query("page_size")).MustInt()
	
	if pageSize <= 0 {
		pageSize = setting.AppSetting.PageSize
	}
	
	if page <= 0 {
		page = 1
	}
	
	offset := (page - 1) * pageSize
	
	return Pagination{
		Page:     page,
		PageSize: pageSize,
		Offset:   offset,
	}
}
