package query

import (
	baseQuery "github.com/jambo0624/blog/internal/shared/domain/query"
	"github.com/jambo0624/blog/internal/shared/domain/constants"
)

type ArticleQuery struct {
	baseQuery.BaseQuery
	CategoryID  *uint  // article specific field
	TagIDs      []uint // article specific field
	TitleLike   string // article specific field
	ContentLike string // article specific field
}

func NewArticleQuery() *ArticleQuery {
	return &ArticleQuery{
		BaseQuery: baseQuery.NewBaseQuery(),
	}
}

func (q *ArticleQuery) WithCategoryID(id uint) *ArticleQuery {
	q.CategoryID = &id
	return q
}

func (q *ArticleQuery) WithTagIDs(ids []uint) *ArticleQuery {
	q.TagIDs = ids
	return q
}

func (q *ArticleQuery) WithTitleLike(title string) *ArticleQuery {
	q.TitleLike = title
	return q
}

func (q *ArticleQuery) WithContentLike(content string) *ArticleQuery {
	q.ContentLike = content
	return q
}

func (q *ArticleQuery) Validate() error {
	if err := q.BaseQuery.Validate(); err != nil {
		return err
	}
	if len(q.TitleLike) > constants.MaxTitleLength {
		return baseQuery.ErrTitleTooLong
	}
	if len(q.ContentLike) > constants.MaxContentLength {
		return baseQuery.ErrContentTooLong
	}
	return nil
}
