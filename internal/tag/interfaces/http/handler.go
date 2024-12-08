package http

import (
	"github.com/gin-gonic/gin"

	"github.com/jambo0624/blog/internal/shared/domain/constants"
	"github.com/jambo0624/blog/internal/shared/domain/errors"
	"github.com/jambo0624/blog/internal/shared/interfaces/http"
	"github.com/jambo0624/blog/internal/tag/application/service"
	"github.com/jambo0624/blog/internal/tag/domain/entity"
	tagQuery "github.com/jambo0624/blog/internal/tag/domain/query"
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

// Only need to implement buildQuery method.
func (h *TagHandler) buildQuery(c *gin.Context) (*tagQuery.TagQuery, error) {
	q := tagQuery.NewTagQuery()
	builder := http.NewBaseQueryBuilder()

	// Build IDs
	if ids, err := builder.BuildIDs(c); err != nil {
		return nil, err
	} else if ids != nil {
		q.WithIDs(ids)
	}

	// Parse name
	if name := c.Query("name"); name != "" {
		if len(name) > constants.MaxNameLength {
			return nil, errors.ErrNameTooLong
		}
		q.WithNameLike(name)
	}

	// Build pagination
	if limit, offset, err := builder.BuildPagination(c, q.Limit, q.Offset); err != nil {
		return nil, err
	} else {
		q.WithPagination(limit, offset)
	}

	// Build order by
	if orderBy, err := builder.BuildOrderBy(c, map[string]bool{
		"name": true,
		// Add other Tag specific fields
	}); err != nil {
		return nil, err
	} else if orderBy != "" {
		q.WithOrderBy(orderBy)
	}

	return q, nil
}

// FindAll overrides BaseHandler.FindAll to use buildQuery.
func (h *TagHandler) FindAll(c *gin.Context) {
	h.BaseHandler.FindAll(c, h.buildQuery)
}
