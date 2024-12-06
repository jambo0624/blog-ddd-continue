package http

import (
	"github.com/gin-gonic/gin"
	tagService "github.com/jambo0624/blog/internal/tag/application/service"
	"github.com/jambo0624/blog/internal/tag/domain/query"
	"github.com/jambo0624/blog/internal/tag/interfaces/http/dto"
	sharedHttp "github.com/jambo0624/blog/internal/shared/interfaces/http"
	"strconv"
	"strings"
)

type TagHandler struct {
	service *tagService.TagService
}

func NewTagHandler(ts *tagService.TagService) *TagHandler {
	return &TagHandler{
		service: ts,
	}
}

func (h *TagHandler) Create(c *gin.Context) {
	var req dto.CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tag, err := h.service.Create(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, tag)
}

func (h *TagHandler) FindByID(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	tag, err := h.service.FindByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Tag not found"})
		return
	}
	c.JSON(200, tag)
}

func (h *TagHandler) FindAll(c *gin.Context) {
	// Create query object
	q := query.NewTagQuery()
	
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
	
	// Parse name like
	if name := c.Query("name"); name != "" {
		// Add name validation rule
		if len(name) > 100 {
			c.JSON(400, gin.H{"error": "name too long"})
			return
		}
		q.WithNameLike(name)
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
		// Validate order field
		allowedFields := map[string]bool{
			"id":         true,
			"name":       true,
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
	tags, err := h.service.FindAll(q)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Return result
	c.JSON(200, gin.H{
		"data": tags,
		"meta": gin.H{
			"limit":  q.Limit,
			"offset": q.Offset,
			"total":  len(tags), // In actual application, it may need to query total separately
		},
	})
}

func (h *TagHandler) Update(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	var req dto.UpdateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tag, err := h.service.Update(id, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, tag)
}

func (h *TagHandler) Delete(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
} 