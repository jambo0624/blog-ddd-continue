package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
	articleService "github.com/jambo0624/blog/internal/article/application/service"
	sharedHttp "github.com/jambo0624/blog/internal/shared/interfaces/http"
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
	"github.com/jambo0624/blog/internal/shared/domain/constants"
	"github.com/jambo0624/blog/internal/shared/domain/query"
	articleQuery "github.com/jambo0624/blog/internal/article/domain/query"
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	"github.com/jambo0624/blog/internal/shared/interfaces/http/response"
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
	builder := sharedHttp.NewBaseQueryBuilder()	

	// Build IDs
	if ids, err := builder.BuildIDs(c); err != nil {
		return nil, err
	} else if ids != nil {
		q.WithIDs(ids)
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
		if len(title) > constants.MaxNameLength {
			return nil, query.ErrTitleTooLong
		}
		q.WithTitleLike(title)
	}

	// Parse content search
	if content := c.Query("content"); content != "" {
		if len(content) > constants.MaxContentLength {
			return nil, query.ErrContentTooLong
		}
		q.WithContentLike(content)
	}

	// Build pagination
	if limit, offset, err := builder.BuildPagination(c, q.Limit, q.Offset); err != nil {
		return nil, err
	} else {
		q.WithPagination(limit, offset)
	}

	// Build order by
	if orderBy, err := builder.BuildOrderBy(c, map[string]bool{
		"title": true,
	}); err != nil {
		return nil, err
	} else if orderBy != "" {
		q.WithOrderBy(orderBy)
	}

	return q, nil
}

func (h *ArticleHandler) FindAll(c *gin.Context) {
	h.BaseHandler.FindAll(c, h.buildQuery)
}

func (h *ArticleHandler) FindByID(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	query := articleQuery.NewArticleQuery()
	preloadAssociations := query.GetPreloadAssociations()
	entity, err := h.Service.FindByID(id, preloadAssociations...)

	if err != nil {
		response.NotFound(c)
		return
	}
	response.Success(c, entity)
}
