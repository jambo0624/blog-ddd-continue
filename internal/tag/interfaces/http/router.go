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
		tags.POST("/", r.handler.CreateTag)
		tags.GET("/:id", r.handler.GetTag)
		tags.PUT("/:id", r.handler.UpdateTag)
		tags.DELETE("/:id", r.handler.DeleteTag)
	}
} 