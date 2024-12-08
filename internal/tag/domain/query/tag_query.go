package query

import (
	"gorm.io/gorm"

	baseQuery "github.com/jambo0624/blog/internal/shared/domain/query"
)

type TagQuery struct {
	baseQuery.BaseQuery
	NameLike  string `binding:"omitempty,max=100"  validate:"omitempty,max=100"`
	ColorLike string `binding:"omitempty,hexcolor" validate:"omitempty,hexcolor"`
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

func (q *TagQuery) Validate() error {
	return q.BaseQuery.ValidateQuery(q)
}

func (q *TagQuery) GetBaseQuery() baseQuery.BaseQuery {
	return q.BaseQuery
}

func (q *TagQuery) ApplyFilters(db *gorm.DB) *gorm.DB {
	if len(q.IDs) > 0 {
		db = db.Where("id IN ?", q.IDs)
	}
	if q.NameLike != "" {
		db = db.Where("name LIKE ?", "%"+q.NameLike+"%")
	}
	if q.ColorLike != "" {
		db = db.Where("color LIKE ?", "%"+q.ColorLike+"%")
	}

	return db
}
