package bootstrap

import (
	"github.com/gin-gonic/gin"

	articleHttp "github.com/jambo0624/blog/internal/article/interfaces/http"
	categoryHttp "github.com/jambo0624/blog/internal/category/interfaces/http"
	tagHttp "github.com/jambo0624/blog/internal/tag/interfaces/http"
)

type Router interface {
	Register(r *gin.RouterGroup)
}

func SetupRouter(handlers *Handlers) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	// register all router groups
	routers := []Router{
		articleHttp.NewArticleRouter(handlers.Article),
		categoryHttp.NewCategoryRouter(handlers.Category),
		tagHttp.NewTagRouter(handlers.Tag),
	}

	// register all router groups
	for _, r := range routers {
		r.Register(api)
	}

	return r
}
