package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"github.com/EDDYCJY/go-gin-example/pkg/e"
)

// 修改验证器的初始化方式
func init() {
	// 使用 gin 的验证器实例
	validate := binding.Validator.Engine().(*validator.Validate)

	// 注册自定义验证器
	validate.RegisterValidation("is-valid-state", validateState)
}

// validateState 验证状态值是否有效
func validateState(fl validator.FieldLevel) bool {
	state := fl.Field().Int()
	return state == 0 || state == 1
}

// BindAndValid binds and validates data
// func BindAndValid(c *gin.Context, form interface{}) (int, int) {
// 	// 进行 JSON 绑定和验证
// 	err := c.ShouldBindJSON(form)
// 	if err != nil {
// 		logging.Error("ShouldBind error: %v", err)
// 		return http.StatusBadRequest, e.INVALID_PARAMS
// 	}

// 	return http.StatusOK, e.SUCCESS
// }

// BindAndValidWithErrors binds and validates data with detailed error messages
func BindAndValidWithErrors(c *gin.Context, form interface{}) (int, int, map[string]string) {
	err := c.ShouldBind(form)
	if err != nil {
		validationErrors := make(map[string]string)

		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, err := range errs {
				validationErrors[err.Field()] = err.Error()
			}
		} else {
			validationErrors["error"] = err.Error()
		}

		return http.StatusBadRequest, e.INVALID_PARAMS, validationErrors
	}

	return http.StatusOK, e.SUCCESS, nil
}
