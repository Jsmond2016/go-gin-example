package util

import (
	"time"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

// Setup 初始化工具包
func Setup() {
	config = Config{
		JWT: JWTConfig{
			Secret:         []byte(setting.AppSetting.JwtSecret),
			ExpireDuration: 3 * time.Hour,
			Issuer:        "gin-blog",
		},
	}
}