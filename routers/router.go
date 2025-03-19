package routers

import (
	"net/http"

	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/EDDYCJY/go-gin-example/docs"
	"github.com/EDDYCJY/go-gin-example/middleware/cors"
	"github.com/EDDYCJY/go-gin-example/middleware/jwt"
	"github.com/EDDYCJY/go-gin-example/middleware/logger"
	"github.com/EDDYCJY/go-gin-example/middleware/permission"
	"github.com/EDDYCJY/go-gin-example/middleware/ratelimit"
	"github.com/EDDYCJY/go-gin-example/middleware/recovery"
	"github.com/EDDYCJY/go-gin-example/pkg/export"
	"github.com/EDDYCJY/go-gin-example/pkg/qrcode"
	"github.com/EDDYCJY/go-gin-example/pkg/upload"
	"github.com/EDDYCJY/go-gin-example/routers/api"
	v1 "github.com/EDDYCJY/go-gin-example/routers/api/v1"
)

// InitRouter initialize routing information
// 将 CORS 配置抽取为独立函数
// func corsConfig() cors.Config {
// 	return cors.Config{
// 		AllowOrigins:     []string{"http://localhost:3000"},
// 		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
// 		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
// 		ExposeHeaders:    []string{"Content-Length"},
// 		AllowCredentials: true,
// 		MaxAge:           12 * time.Hour,
// 	}
// }

func InitRouter() *gin.Engine {
	r := gin.New()

	// 基础中间件
	r.Use(recovery.Recovery()) // panic 恢复
	r.Use(logger.Logger())     // 请求日志
	r.Use(gin.Logger())        // Request time logging
	r.Use(cors.Cors())         // CORS 跨域

	// 限流中间件（每分钟允许100个请求，突发200个）
	rateLimiter := ratelimit.NewRateLimiter(100, 200)
	r.Use(rateLimiter.RateLimit())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)
	r.POST("/api/v1/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())                    // JWT 认证
	apiv1.Use(permission.CheckPermission()) // 权限检查
	{
		// Tags
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		apiv1.POST("/tags/export", v1.ExportTag)
		apiv1.POST("/tags/import", v1.ImportTag)

		// Articles
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
