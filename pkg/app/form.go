package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
)

var validate *validator.Validate

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		v.RegisterValidation("is-valid-state", validateState)
	}
}

// validateState 验证状态值是否有效
func validateState(fl validator.FieldLevel) bool {
	state := fl.Field().Int()
	return state == 0 || state == 1
}

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.ShouldBind(form)
	if err != nil {
		logging.Error("ShouldBind error: %v", err)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	err = validate.Struct(form)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			logging.Error("InvalidValidationError: %v", err)
			return http.StatusInternalServerError, e.ERROR
		}

		for _, err := range err.(validator.ValidationErrors) {
			logging.Error("ValidationError: %v", err)
		}
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
