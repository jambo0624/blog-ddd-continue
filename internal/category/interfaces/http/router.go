package http

import (
	"github.com/gin-gonic/gin"
)

type CategoryRouter struct {
	handler *CategoryHandler
	engine  *gin.Engine
}

func NewCategoryRouter(handler *CategoryHandler) *CategoryRouter {
	return &CategoryRouter{handler: handler, engine: gin.New()}
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

func (r *CategoryRouter) Engine() *gin.Engine {
	return r.engine
}
