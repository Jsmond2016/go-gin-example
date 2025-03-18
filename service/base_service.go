package service

import "github.com/EDDYCJY/go-gin-example/pkg/logging"

// BaseService 提供基础服务功能
type BaseService struct {
	CacheExpire int // 缓存过期时间(秒)
}

// NewBaseService 创建基础服务
func NewBaseService() BaseService {
	return BaseService{
		CacheExpire: 3600, // 默认缓存1小时
	}
}

// handleCacheError 处理缓存错误
func (b *BaseService) HandleCacheError(err error, msg string) {
	if err != nil {
		logging.Error(msg, err)
	}
}
