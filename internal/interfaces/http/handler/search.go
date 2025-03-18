package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/internal/service"
	"github.com/EDDYCJY/go-gin-example/pkg/response"
)

type SearchHandler struct {
	searchService service.SearchService
}

func NewSearchHandler(searchService service.SearchService) *SearchHandler {
	return &SearchHandler{
		searchService: searchService,
	}
}

// SearchArticles godoc
// @Summary Search articles
// @Description Search articles by keyword
// @Tags articles
// @Accept  json
// @Produce  json
// @Param keyword query string true "Search keyword"
// @Param page query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} response.Response
// @Router /articles/search [get]
func (h *SearchHandler) SearchArticles(c *gin.Context) {
	keyword := c.Query("keyword")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	articles, total, err := h.searchService.SearchArticles(c.Request.Context(), keyword, page, pageSize)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerError, err.Error())
		return
	}

	response.Success(c, gin.H{
		"items": articles,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}