package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/internal/domain"
	"github.com/EDDYCJY/go-gin-example/internal/service"
	"github.com/EDDYCJY/go-gin-example/pkg/response"
)

type ArticleHandler struct {
	articleService service.ArticleService
}

func NewArticleHandler(articleService service.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		articleService: articleService,
	}
}

// CreateArticle godoc
// @Summary Create a new article
// @Description Create a new article with the input payload
// @Tags articles
// @Accept  json
// @Produce  json
// @Param article body domain.Article true "Create article"
// @Success 200 {object} response.Response
// @Router /articles [post]
func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var article domain.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.articleService.CreateArticle(c.Request.Context(), &article); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerError, err.Error())
		return
	}

	response.Success(c, article)
}

// 实现其他处理方法...