package util

import "github.com/EDDYCJY/go-gin-example/pkg/setting"

var jwtSecret []byte

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}