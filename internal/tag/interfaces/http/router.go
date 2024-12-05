package http

import (
	"github.com/gin-gonic/gin"
)

type TagRouter struct {
	handler *TagHandler
}

func NewTagRouter(handler *TagHandler) *TagRouter {
	return &TagRouter{handler: handler}
}

func (r *TagRouter) Register(api *gin.RouterGroup) {
	tags := api.Group("/tags")
	{
		tags.POST("/", r.handler.Create)
		tags.GET("/:id", r.handler.FindByID)
		tags.GET("/", r.handler.FindAll)
		tags.PUT("/:id", r.handler.Update)
		tags.DELETE("/:id", r.handler.Delete)
	}
} 