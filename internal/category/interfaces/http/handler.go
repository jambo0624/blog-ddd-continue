package http

import (
	"github.com/gin-gonic/gin"
	categoryService "github.com/jambo0624/blog/internal/category/application/service"
	"github.com/jambo0624/blog/internal/shared/domain/constants"
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
			return nil, query.ErrNameTooLong
		}
		q.WithNameLike(name)
	}
    
	// Parse slug
	if slug := c.Query("slug"); slug != "" {
		if len(slug) > constants.MaxSlugLength {
			return nil, query.ErrSlugTooLong
		}
		q.WithSlugLike(slug)
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
		"slug": true,
	}); err != nil {
		return nil, err
	} else if orderBy != "" {
		q.WithOrderBy(orderBy)
	}

	return q, nil
}

func (h *CategoryHandler) FindAll(c *gin.Context) {
	h.BaseHandler.FindAll(c, h.buildQuery)
}
