package http

import (
    "github.com/gin-gonic/gin"
    categoryService "github.com/jambo0624/blog/internal/category/application/service"
	sharedHttp "github.com/jambo0624/blog/internal/shared/interfaces/http"
)

type CategoryHandler struct {
	categoryService *categoryService.CategoryService
}

func NewCategoryHandler(cs *categoryService.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: cs,
    }
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
    var req struct {
        Name string `json:"name" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    category, err := h.categoryService.CreateCategory(req.Name)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(201, category)
}

func (h *CategoryHandler) GetCategory(c *gin.Context) {
    id := sharedHttp.ParseUintParam(c, "id")
    category, err := h.categoryService.GetCategoryByID(id)
    if err != nil {
        c.JSON(404, gin.H{"error": "Category not found"})
        return
    }
    c.JSON(200, category)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
    id := sharedHttp.ParseUintParam(c, "id")
    var req struct {
        Name string `json:"name" binding:"required"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    category, err := h.categoryService.UpdateCategory(id, req.Name)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, category)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
    id := sharedHttp.ParseUintParam(c, "id")
    if err := h.categoryService.DeleteCategory(id); err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(204, nil)
} 