package query

import (
	"github.com/jambo0624/blog/internal/shared/domain/constants"
	"github.com/jambo0624/blog/internal/shared/domain/validate"
)

// BaseQuery base query struct
type BaseQuery struct {
	IDs                 []uint   // for IN query
	Limit               int      // for pagination
	Offset              int      // for pagination
	OrderBy             string   // for sorting
	PreloadAssociations []string // store associations to be preloaded
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

// Validate validate the query parameters
func (q *BaseQuery) Validate() error {
	if q.Limit < 0 {
		return validate.ErrInvalidLimit
	}
	if q.Offset < 0 {
		return validate.ErrInvalidOffset
	}
	return nil
}

// GetPreloadAssociations get the preload associations
func (q *BaseQuery) GetPreloadAssociations() []string {
	return q.PreloadAssociations
}
