package cache_service

import (
	"strconv"
	"strings"

	"github.com/EDDYCJY/go-gin-example/pkg/e"
)

type Article struct {
	ID    uint
	TagID uint
	State int

	PageNum  int
	PageSize int
}

func (a *Article) GetArticleKey() string {
	return e.CACHE_ARTICLE + "_" + strconv.FormatUint(uint64(a.ID), 10)
}

func (a *Article) GetArticlesKey() string {
	keys := []string{
		e.CACHE_ARTICLE,
		"LIST",
	}

	if a.ID > 0 {
		keys = append(keys, strconv.FormatUint(uint64(a.ID), 10))
	}
	if a.TagID > 0 {
		keys = append(keys, strconv.FormatUint(uint64(a.TagID), 10))
	}
	if a.State >= 0 {
		keys = append(keys, strconv.Itoa(a.State))
	}
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}
