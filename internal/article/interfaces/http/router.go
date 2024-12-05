package http

import (
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
	handler *ArticleHandler
}

func NewArticleRouter(handler *ArticleHandler) *ArticleRouter {
	return &ArticleRouter{handler: handler}
}

func (r *ArticleRouter) Register(api *gin.RouterGroup) {
	articles := api.Group("/articles")
	{
		articles.POST("/", r.handler.CreateArticle)
		articles.GET("/", r.handler.GetAllArticles)
		articles.GET("/:id", r.handler.GetArticle)
		articles.PUT("/:id", r.handler.UpdateArticle)
		articles.DELETE("/:id", r.handler.DeleteArticle)
	}
} 