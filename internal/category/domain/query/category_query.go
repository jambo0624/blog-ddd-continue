package query

import (
	"gorm.io/gorm"

	baseQuery "github.com/jambo0624/blog/internal/shared/domain/query"
)

type CategoryQuery struct {
	baseQuery.BaseQuery
	NameLike string `json:"name_like" binding:"omitempty, max=100" validate:"omitempty,max=100"`
	SlugLike string `json:"slug_like" binding:"omitempty, max=100" validate:"omitempty,max=100"`
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
	return q.BaseQuery.ValidateQuery(q)
}

func (q *CategoryQuery) GetBaseQuery() baseQuery.BaseQuery {
	return q.BaseQuery
}

func (q *CategoryQuery) ApplyFilters(db *gorm.DB) *gorm.DB {
	if len(q.IDs) > 0 {
		db = db.Where("id IN ?", q.IDs)
	}
	if q.NameLike != "" {
		db = db.Where("name LIKE ?", "%"+q.NameLike+"%")
	}
	if q.SlugLike != "" {
		db = db.Where("slug LIKE ?", "%"+q.SlugLike+"%")
	}
	return db
}
