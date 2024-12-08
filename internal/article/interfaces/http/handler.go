package http

import (
	"strconv"

	"github.com/gin-gonic/gin"

	articleService "github.com/jambo0624/blog/internal/article/application/service"
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	articleQuery "github.com/jambo0624/blog/internal/article/domain/query"
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
	"github.com/jambo0624/blog/internal/shared/domain/constants"
	"github.com/jambo0624/blog/internal/shared/domain/errors"
	sharedHttp "github.com/jambo0624/blog/internal/shared/interfaces/http"
	"github.com/jambo0624/blog/internal/shared/interfaces/http/response"
)

type ArticleHandler struct {
	*sharedHttp.BaseHandler[
		articleEntity.Article,
		*articleQuery.ArticleQuery,
		dto.CreateArticleRequest,
		dto.UpdateArticleRequest,
	]
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

	if err := h.applyIDFilters(c, q, builder); err != nil {
		return nil, err
	}

	if err := h.applyCategoryFilter(c, q); err != nil {
		return nil, err
	}

	if err := h.applyTagFilters(c, q); err != nil {
		return nil, err
	}

	if err := h.applyTextFilters(c, q); err != nil {
		return nil, err
	}

	if err := h.applyPaginationAndOrder(c, q, builder); err != nil {
		return nil, err
	}

	return q, nil
}

func (h *ArticleHandler) applyIDFilters(c *gin.Context, q *articleQuery.ArticleQuery, builder *sharedHttp.BaseQueryBuilder) error {
	ids, err := builder.BuildIDs(c)
	if err != nil {
		return err
	}
	if ids != nil {
		q.WithIDs(ids)
	}
	return nil
}

func (h *ArticleHandler) applyCategoryFilter(c *gin.Context, q *articleQuery.ArticleQuery) error {
	if categoryID := c.Query("category_id"); categoryID != "" {
		uid, err := strconv.ParseUint(categoryID, 10, 32)
		if err != nil {
			return errors.ErrInvalidIDFormat
		}
		q.WithCategoryID(uint(uid))
	}
	return nil
}

func (h *ArticleHandler) applyTagFilters(c *gin.Context, q *articleQuery.ArticleQuery) error {
	if tagIDs := c.QueryArray("tag_ids"); len(tagIDs) > 0 {
		uintIDs := make([]uint, 0, len(tagIDs))
		for _, id := range tagIDs {
			uid, err := strconv.ParseUint(id, 10, 32)
			if err != nil {
				return errors.ErrInvalidIDFormat
			}
			uintIDs = append(uintIDs, uint(uid))
		}
		q.WithTagIDs(uintIDs)
	}
	return nil
}

func (h *ArticleHandler) applyTextFilters(c *gin.Context, q *articleQuery.ArticleQuery) error {
	if title := c.Query("title"); title != "" {
		if len(title) > constants.MaxNameLength {
			return errors.ErrTitleTooLong
		}
		q.WithTitleLike(title)
	}

	if content := c.Query("content"); content != "" {
		if len(content) > constants.MaxContentLength {
			return errors.ErrContentTooLong
		}
		q.WithContentLike(content)
	}
	return nil
}

func (h *ArticleHandler) applyPaginationAndOrder(c *gin.Context, q *articleQuery.ArticleQuery, builder *sharedHttp.BaseQueryBuilder) error {
	limit, offset, err := builder.BuildPagination(c, q.Limit, q.Offset)
	if err != nil {
		return err
	}
	q.WithPagination(limit, offset)

	orderBy, err := builder.BuildOrderBy(c, map[string]bool{
		"title": true,
	})
	if err != nil {
		return err
	}
	if orderBy != "" {
		q.WithOrderBy(orderBy)
	}
	return nil
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
