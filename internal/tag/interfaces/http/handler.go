package http

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jambo0624/blog/internal/shared/interfaces/http"
	"github.com/jambo0624/blog/internal/tag/domain/entity"
	tagQuery "github.com/jambo0624/blog/internal/tag/domain/query"
	"github.com/jambo0624/blog/internal/shared/domain/query"
	"github.com/jambo0624/blog/internal/tag/application/service"
	"github.com/jambo0624/blog/internal/tag/interfaces/http/dto"
)

type TagHandler struct {
	*http.BaseHandler[entity.Tag, *tagQuery.TagQuery, dto.CreateTagRequest, dto.UpdateTagRequest]
}

func NewTagHandler(s *service.TagService) *TagHandler {
	baseHandler := http.NewBaseHandler(s.BaseService, s)
	return &TagHandler{
		BaseHandler: baseHandler,
	}
}

// Only need to implement buildQuery method
func (h *TagHandler) buildQuery(c *gin.Context) (*tagQuery.TagQuery, error) {
	// create new query
	q := tagQuery.NewTagQuery()

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
	
	// Parse name like
	if name := c.Query("name"); name != "" {
		// Add name validation rule
		if len(name) > 100 {
			return nil, query.ErrNameTooLong
		}
		q.WithNameLike(name)
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
		// Validate order field
		allowedFields := map[string]bool{
			"id":         true,
			"name":       true,
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

// FindAll overrides BaseHandler.FindAll to use buildQuery
func (h *TagHandler) FindAll(c *gin.Context) {
	h.BaseHandler.FindAll(c, h.buildQuery)
}
