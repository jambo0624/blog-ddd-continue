package http

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	categoryService "github.com/jambo0624/blog/internal/category/application/service"
	"github.com/jambo0624/blog/internal/shared/interfaces/http"
	"github.com/jambo0624/blog/internal/category/interfaces/http/dto"
	"github.com/jambo0624/blog/internal/shared/domain/query"
	categoryQuery "github.com/jambo0624/blog/internal/category/domain/query"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
)

type CategoryHandler struct {
	*http.BaseHandler[categoryEntity.Category, *categoryQuery.CategoryQuery, dto.CreateCategoryRequest, dto.UpdateCategoryRequest]
}

func NewCategoryHandler(cs *categoryService.CategoryService) *CategoryHandler {
	baseHandler := http.NewBaseHandler(cs.BaseService, cs	)
	return &CategoryHandler{
		BaseHandler: baseHandler,
  }
}


func (h *CategoryHandler) buildQuery(c *gin.Context) (*categoryQuery.CategoryQuery, error) {
	q := categoryQuery.NewCategoryQuery()
	// Parse IDs
	if ids := c.QueryArray("ids"); len(ids) > 0 {
		uintIDs := make([]uint, 0, len(ids))
		for _, id := range ids {
			uid, err := strconv.ParseUint(id, 10, 32)
			if err != nil {
				return nil, query.ErrInvalidIDFormat
			}
			uintIDs = append(uintIDs, uint(uid))
		}
		q.WithIDs(uintIDs)
	}
    
	// Parse name search
	if name := c.Query("name"); name != "" {
		if len(name) > 100 {
			return nil, query.ErrNameTooLong
		}
		q.WithNameLike(name)
	}
    
	// Parse slug search
	if slug := c.Query("slug"); slug != "" {
		if len(slug) > 100 {
			return nil, query.ErrSlugTooLong
		}
		q.WithSlugLike(slug)
	}
    
	// Parse pagination
	if limitStr := c.Query("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 0 {
			return nil, query.ErrInvalidLimit
		}
		if limit > 100 { // Set max limit
			limit = 100
		}
		q.WithPagination(limit, q.Offset)
	}
    
	if offsetStr := c.Query("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			return nil, query.ErrInvalidOffset
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
			return nil, query.ErrInvalidOrderByField
		}
        
		q.WithOrderBy(orderBy)
	}

	return q, nil
}

func (h *CategoryHandler) FindAll(c *gin.Context) {
	h.BaseHandler.FindAll(c, h.buildQuery)
}
