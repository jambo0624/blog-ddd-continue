package query

import (
    baseQuery "github.com/jambo0624/blog/internal/shared/domain/query"
)

type CategoryQuery struct {
    baseQuery.BaseQuery
    NameLike string  // category specific field
    SlugLike string  // category specific field
}

func NewCategoryQuery() *CategoryQuery {
    return &CategoryQuery{
        BaseQuery: baseQuery.NewBaseQuery(),
    }
}

func (q *CategoryQuery) WithNameLike(name string) *CategoryQuery {
    q.NameLike = name
    return q
}

func (q *CategoryQuery) WithSlugLike(slug string) *CategoryQuery {
    q.SlugLike = slug
    return q
}

func (q *CategoryQuery) Validate() error {
    if err := q.BaseQuery.Validate(); err != nil {
        return err
    }
    if len(q.NameLike) > 100 {
        return baseQuery.ErrNameTooLong
    }
    return nil
} 