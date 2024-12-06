package http

import (
	"github.com/gin-gonic/gin"
	articleService "github.com/jambo0624/blog/internal/article/application/service"
	sharedHttp "github.com/jambo0624/blog/internal/shared/interfaces/http"
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
	"github.com/jambo0624/blog/internal/article/domain/query"
	"strconv"
	"strings"
)

type ArticleHandler struct {
	service *articleService.ArticleService
}

func NewArticleHandler(as *articleService.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		service: as,
	}
}

func (h *ArticleHandler) Create(c *gin.Context) {
	var req dto.CreateArticleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	article, err := h.service.Create(&req)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, article)
}

func (h *ArticleHandler) FindByID(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	article, err := h.service.FindByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(200, article)
}

func (h *ArticleHandler) FindAll(c *gin.Context) {
	// Create query object
	q := query.NewArticleQuery()
	
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

	// Parse category ID
	if categoryID := c.Query("category_id"); categoryID != "" {
		uid, err := strconv.ParseUint(categoryID, 10, 32)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid category_id format"})
			return
		}
		q.WithCategoryID(uint(uid))
	}

	// Parse tag IDs
	if tagIDs := c.QueryArray("tag_ids"); len(tagIDs) > 0 {
		uintIDs := make([]uint, 0, len(tagIDs))
		for _, id := range tagIDs {
			uid, err := strconv.ParseUint(id, 10, 32)
			if err != nil {
				c.JSON(400, gin.H{"error": "invalid tag_id format"})
				return
			}
			uintIDs = append(uintIDs, uint(uid))
		}
		q.WithTagIDs(uintIDs)
	}

	// Parse title search
	if title := c.Query("title"); title != "" {
		if len(title) > 255 {
			c.JSON(400, gin.H{"error": "title too long"})
			return
		}
		q.WithTitleLike(title)
	}

	// Parse content search
	if content := c.Query("content"); content != "" {
		if len(content) > 1000 {
			c.JSON(400, gin.H{"error": "content search term too long"})
			return
		}
		q.WithContentLike(content)
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
			"title":      true,
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
	articles, err := h.service.FindAll(q)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Return result
	c.JSON(200, gin.H{
		"data": articles,
		"meta": gin.H{
			"limit":  q.Limit,
			"offset": q.Offset,
			"total":  len(articles),
		},
	})
}

func (h *ArticleHandler) Update(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")

	var req dto.UpdateArticleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	article, err := h.service.Update(id, &req)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, article)
}

func (h *ArticleHandler) Delete(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
} 