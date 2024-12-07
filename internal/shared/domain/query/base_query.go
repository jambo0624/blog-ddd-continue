package query

import (
	"github.com/go-playground/validator/v10"
	"github.com/jambo0624/blog/internal/shared/domain/constants"
)

var ValidateQuery = validator.New()

// BaseQuery base query struct
type BaseQuery struct {
	IDs                 []uint   `json:"ids" binding:"omitempty" validate:"omitempty,dive,gt=0"`
	Limit               int      `json:"limit" binding:"omitempty, min=1" validate:"omitempty,min=1"`
	Offset              int      `json:"offset" binding:"omitempty, min=0" validate:"omitempty,min=0"`
	OrderBy             string   `json:"order_by" binding:"omitempty" validate:"omitempty"`
	PreloadAssociations []string `json:"preload_associations" binding:"omitempty"`
}

// NewBaseQuery create a new base query
func NewBaseQuery() BaseQuery {
	return BaseQuery{
		Limit:               constants.DefaultPageSize,
		Offset:              constants.DefaultPageOffset,
		PreloadAssociations: []string{},
	}
}

// WithIDs add ID filter
func (q *BaseQuery) WithIDs(ids []uint) *BaseQuery {
	q.IDs = ids
	return q
}

// WithPagination add pagination filter
func (q *BaseQuery) WithPagination(limit, offset int) *BaseQuery {
	q.Limit = limit
	q.Offset = offset
	return q
}

// WithOrderBy add order by filter
func (q *BaseQuery) WithOrderBy(orderBy string) *BaseQuery {
	q.OrderBy = orderBy
	return q
}

// ValidateQuery validate the query parameters
func (q *BaseQuery) ValidateQuery(v any) error {
	return ValidateQuery.Struct(v)
}

// GetPreloadAssociations get the preload associations
func (q *BaseQuery) GetPreloadAssociations() []string {
	return q.PreloadAssociations
}
