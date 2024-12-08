package http

import (
	"github.com/gin-gonic/gin"
)

type ArticleRouter struct {
	handler *ArticleHandler
	engine  *gin.Engine
}

func NewArticleRouter(handler *ArticleHandler) *ArticleRouter {
	return &ArticleRouter{handler: handler, engine: gin.New()}
}

func (r *ArticleRouter) Register(api *gin.RouterGroup) {
	articles := api.Group("/articles")
	{
		articles.POST("", r.handler.Create)
		articles.GET("", r.handler.FindAll)
		articles.GET("/:id", r.handler.FindByID)
		articles.PUT("/:id", r.handler.Update)
		articles.DELETE("/:id", r.handler.Delete)
	}
}

func (r *ArticleRouter) Engine() *gin.Engine {
	return r.engine
}
