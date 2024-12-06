package http

import (
    "github.com/gin-gonic/gin"
    categoryService "github.com/jambo0624/blog/internal/category/application/service"
	sharedHttp "github.com/jambo0624/blog/internal/shared/interfaces/http"
	"github.com/jambo0624/blog/internal/category/interfaces/http/dto"
	"github.com/jambo0624/blog/internal/category/domain/query"
	"strconv"
	"strings"
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
    // Create query object
    q := query.NewCategoryQuery()
    
    // Parse IDs
    if ids := c.QueryArray("ids"); len(ids) > 0 {
        uintIDs := make([]uint, 0, len(ids))
        for _, id := range ids {
            uid, err := strconv.ParseUint(id, 10, 32)
            if err != nil {
                c.JSON(400, gin.H{"error": "invalid id format"})
                return
            }
            uintIDs = append(uintIDs, uint(uid))
        }
        q.WithIDs(uintIDs)
    }
    
    // Parse name search
    if name := c.Query("name"); name != "" {
        if len(name) > 100 {
            c.JSON(400, gin.H{"error": "name too long"})
            return
        }
        q.WithNameLike(name)
    }
    
    // Parse slug search
    if slug := c.Query("slug"); slug != "" {
        if len(slug) > 100 {
            c.JSON(400, gin.H{"error": "slug too long"})
            return
        }
        q.WithSlugLike(slug)
    }
    
    // Parse pagination
    if limitStr := c.Query("limit"); limitStr != "" {
        limit, err := strconv.Atoi(limitStr)
        if err != nil || limit < 0 {
            c.JSON(400, gin.H{"error": "invalid limit"})
            return
        }
        if limit > 100 { // Set max limit
            limit = 100
        }
        q.WithPagination(limit, q.Offset)
    }
    
    if offsetStr := c.Query("offset"); offsetStr != "" {
        offset, err := strconv.Atoi(offsetStr)
        if err != nil || offset < 0 {
            c.JSON(400, gin.H{"error": "invalid offset"})
            return
        }
        q.WithPagination(q.Limit, offset)
    }
    
    // Parse ordering
    if orderBy := c.Query("order_by"); orderBy != "" {
        allowedFields := map[string]bool{
            "id":         true,
            "name":       true,
            "slug":       true,
            "created_at": true,
            "updated_at": true,
        }
        
        field := strings.TrimSuffix(strings.TrimPrefix(orderBy, "-"), " DESC")
        if !allowedFields[field] {
            c.JSON(400, gin.H{"error": "invalid order field"})
            return
        }
        
        q.WithOrderBy(orderBy)
    }

    // Execute query
    categories, err := h.service.FindAll(q)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    // Return result
    c.JSON(200, gin.H{
        "data": categories,
        "meta": gin.H{
            "limit":  q.Limit,
            "offset": q.Offset,
            "total":  len(categories),
        },
    })
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