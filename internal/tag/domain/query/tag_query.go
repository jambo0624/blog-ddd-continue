package query

import (
	"github.com/jambo0624/blog/internal/shared/domain/constants"
	baseQuery "github.com/jambo0624/blog/internal/shared/domain/query"
)

type TagQuery struct {
	baseQuery.BaseQuery
	NameLike  string // tag specific field
	ColorLike string // tag specific field
}

func NewTagQuery() *TagQuery {
	return &TagQuery{
		BaseQuery: baseQuery.NewBaseQuery(),
	}
}

func (q *TagQuery) WithNameLike(name string) *TagQuery {
	q.NameLike = name
	return q
}

func (q *TagQuery) WithColorLike(color string) *TagQuery {
	q.ColorLike = color
	return q
}

// override the base class validation method
// add specific validation rules
func (q *TagQuery) Validate() error {
	if err := q.BaseQuery.Validate(); err != nil {
		return err
	}
	if len(q.NameLike) > constants.MaxNameLength {
		return baseQuery.ErrNameTooLong
	}
	return nil
}
