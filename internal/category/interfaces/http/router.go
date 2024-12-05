package http

import (
	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {
	handler *CategoryHandler
}

func NewCategoryRouter(handler *CategoryHandler) *CategoryRouter {
	return &CategoryRouter{handler: handler}
}

func (r *CategoryRouter) Register(api *gin.RouterGroup) {
	categories := api.Group("/categories")
	{
		categories.POST("/", r.handler.CreateCategory)
		categories.GET("/:id", r.handler.GetCategory)
		categories.PUT("/:id", r.handler.UpdateCategory)
		categories.DELETE("/:id", r.handler.DeleteCategory)
	}
} 