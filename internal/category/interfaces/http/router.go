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
		categories.POST("", r.handler.Create)
		categories.GET("", r.handler.FindAll)
		categories.GET("/:id", r.handler.FindByID)
		categories.PUT("/:id", r.handler.Update)
		categories.DELETE("/:id", r.handler.Delete)
	}
}
