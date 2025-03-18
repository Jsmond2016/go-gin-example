package router

import (
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/internal/handler"
	"github.com/EDDYCJY/go-gin-example/internal/middleware/jwt"
)

func RegisterRoutes(r *gin.Engine, handlers *handler.Handlers) {
	api := r.Group("/api/v1")
	{
		// 公开接口
		api.POST("/register", handlers.User.Register)
		api.POST("/login", handlers.User.Login)

		// 需要认证的接口
		auth := api.Group("/", jwt.JWT())
		{
			auth.GET("/users/:id", handlers.User.GetUserInfo)
		}

		// 文章相关路由
		articles := api.Group("/articles")
		{
			articles.POST("", handlers.Article.CreateArticle)
			articles.GET("", handlers.Article.ListArticles)
			articles.GET("/:id", handlers.Article.GetArticle)
			articles.PUT("/:id", handlers.Article.UpdateArticle)
			articles.GET("/search", handlers.Search.SearchArticles)
			articles.DELETE("/:id", handlers.Article.DeleteArticle)
		}

		// 评论相关路由
		comments := api.Group("/comments")
		{
			comments.POST("", handlers.Comment.CreateComment)
			comments.GET("/article/:id", handlers.Comment.GetArticleComments)
			comments.DELETE("/:id", handlers.Comment.DeleteComment)
		}

		// 分类相关路由
		categories := api.Group("/categories")
		{
			categories.POST("", handlers.Category.CreateCategory)
			categories.GET("", handlers.Category.ListCategories)
			categories.GET("/:id", handlers.Category.GetCategory)
			categories.PUT("/:id", handlers.Category.UpdateCategory)
			categories.DELETE("/:id", handlers.Category.DeleteCategory)
		}

		// 标签相关路由
		tags := api.Group("/tags")
		{
			tags.POST("", handlers.Tag.CreateTag)
			tags.GET("", handlers.Tag.ListTags)
			tags.GET("/:id", handlers.Tag.GetTag)
			tags.PUT("/:id", handlers.Tag.UpdateTag)
			tags.DELETE("/:id", handlers.Tag.DeleteTag)
		}
	}
}