package http

import (
    "github.com/gin-gonic/gin"
    categoryService "github.com/jambo0624/blog/internal/category/application/service"
	sharedHttp "github.com/jambo0624/blog/internal/shared/interfaces/http"
	"github.com/jambo0624/blog/internal/category/interfaces/http/dto"
)

type CategoryHandler struct {
	service *categoryService.CategoryService
}

func NewCategoryHandler(cs *categoryService.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: cs,
    }
}

func (h *CategoryHandler) Create(c *gin.Context) {
    var req dto.CreateCategoryRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    category, err := h.service.Create(&req)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(201, category)
}

func (h *CategoryHandler) FindByID(c *gin.Context) {
    id := sharedHttp.ParseUintParam(c, "id")
    category, err := h.service.FindByID(id)
    if err != nil {
        c.JSON(404, gin.H{"error": "Category not found"})
        return
    }
    c.JSON(200, category)
}

func (h *CategoryHandler) FindAll(c *gin.Context) {
	categories, err := h.service.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, categories)
}

func (h *CategoryHandler) Update(c *gin.Context) {
    id := sharedHttp.ParseUintParam(c, "id")
    var req dto.UpdateCategoryRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    category, err := h.service.Update(id, &req)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, category)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
    id := sharedHttp.ParseUintParam(c, "id")
    if err := h.service.Delete(id); err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(204, nil)
} 