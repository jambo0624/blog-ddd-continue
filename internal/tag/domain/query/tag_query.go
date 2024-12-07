package query

import (
	"gorm.io/gorm"

	"github.com/jambo0624/blog/internal/shared/domain/constants"
	"github.com/jambo0624/blog/internal/shared/domain/validate"
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
		return validate.ErrNameTooLong
	}
	return nil
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
	return db
}