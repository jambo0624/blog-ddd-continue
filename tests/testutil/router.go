package testutil

import "github.com/gin-gonic/gin"

// Router interface for test routers
type Router interface {
    Register(*gin.RouterGroup)
    Engine() *gin.Engine
} 