package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/internal/domain"
	"github.com/EDDYCJY/go-gin-example/internal/service"
	"github.com/EDDYCJY/go-gin-example/pkg/response"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with the input payload
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body domain.Category true "Create category"
// @Success 200 {object} response.Response
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var category domain.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.categoryService.CreateCategory(c.Request.Context(), &category); err != nil {
		response.Error(c, http.StatusInternalServerError, response.CodeServerError, err.Error())
		return
	}

	response.Success(c, category)
}

// 实现其他处理方法...