package util

import "time"

// JWTConfig JWT配置
type JWTConfig struct {
	Secret        []byte
	ExpireDuration time.Duration
	Issuer        string
}

// Config 工具包配置
type Config struct {
	JWT JWTConfig
}

var config Config