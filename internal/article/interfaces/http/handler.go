package http

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	articleService "github.com/jambo0624/blog/internal/article/application/service"
	sharedHttp "github.com/jambo0624/blog/internal/shared/interfaces/http"
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
	"github.com/jambo0624/blog/internal/shared/domain/query"
	articleQuery "github.com/jambo0624/blog/internal/article/domain/query"
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
)

type ArticleHandler struct {
	*sharedHttp.BaseHandler[articleEntity.Article, *articleQuery.ArticleQuery, dto.CreateArticleRequest, dto.UpdateArticleRequest]
}

func NewArticleHandler(as *articleService.ArticleService) *ArticleHandler {
	baseHandler := sharedHttp.NewBaseHandler(as.BaseService, as)
	return &ArticleHandler{
		BaseHandler: baseHandler,
	}
}

func (h *ArticleHandler) buildQuery(c *gin.Context) (*articleQuery.ArticleQuery, error) {
	q := articleQuery.NewArticleQuery()
	
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

	// Parse category ID
	if categoryID := c.Query("category_id"); categoryID != "" {
		uid, err := strconv.ParseUint(categoryID, 10, 32)
		if err != nil {
			return nil, query.ErrInvalidIDFormat
		}
		q.WithCategoryID(uint(uid))
	}

	// Parse tag IDs
	if tagIDs := c.QueryArray("tag_ids"); len(tagIDs) > 0 {
		uintIDs := make([]uint, 0, len(tagIDs))
		for _, id := range tagIDs {
			uid, err := strconv.ParseUint(id, 10, 32)
			if err != nil {
				return nil, query.ErrInvalidIDFormat
			}
			uintIDs = append(uintIDs, uint(uid))
		}
		q.WithTagIDs(uintIDs)
	}

	// Parse title search
	if title := c.Query("title"); title != "" {
		if len(title) > 255 {
			return nil, query.ErrTitleTooLong
		}
		q.WithTitleLike(title)
	}

	// Parse content search
	if content := c.Query("content"); content != "" {
		if len(content) > 1000 {
			return nil, query.ErrContentTooLong
		}
		q.WithContentLike(content)
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
			"title":      true,
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

func (h *ArticleHandler) FindAll(c *gin.Context) {
	h.BaseHandler.FindAll(c, h.buildQuery)
}
