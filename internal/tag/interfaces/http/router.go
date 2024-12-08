package http

import (
	"github.com/gin-gonic/gin"
)

type TagRouter struct {
	handler *TagHandler
	engine  *gin.Engine
}

func NewTagRouter(handler *TagHandler) *TagRouter {
	return &TagRouter{
		handler: handler,
		engine:  gin.New(),
	}
}

func (r *TagRouter) Register(api *gin.RouterGroup) {
	tags := api.Group("/tags")
	{
		tags.POST("", r.handler.Create)
		tags.GET("", r.handler.FindAll)
		tags.GET("/:id", r.handler.FindByID)
		tags.PUT("/:id", r.handler.Update)
		tags.DELETE("/:id", r.handler.Delete)
	}
}

func (r *TagRouter) Engine() *gin.Engine {
	return r.engine
} 
