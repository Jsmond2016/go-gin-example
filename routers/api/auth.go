package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EDDYCJY/go-gin-example/service/auth_service"
)

type auth struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Get Auth
// @Produce json
// @Accept json
// @Param auth body auth true "Auth"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/auth [post]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	var a auth

	httpCode, errCode, errors := app.BindAndValidWithErrors(c, &a)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, errors)
		return
	}

	authService := auth_service.Auth{Username: a.Username, Password: a.Password}
	isExist, err := authService.Check()
	if err != nil {
		logging.Error("验证用户失败:", err) // 添加日志
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(a.Username, a.Password)
	if err != nil {
		logging.Error("生成token失败:", err) // 添加日志
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"token":     token,
		"expire_at": time.Now().Add(24 * time.Hour), // 添加过期时间
	})
}
