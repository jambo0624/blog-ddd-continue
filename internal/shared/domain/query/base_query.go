package query

// BaseQuery base query struct
type BaseQuery struct {
    IDs     []uint  // for IN query
    Limit   int     // for pagination
    Offset  int     // for pagination
    OrderBy string  // for sorting
}

// NewBaseQuery create a new base query
func NewBaseQuery() BaseQuery {
    return BaseQuery{
        Limit:  10,
        Offset: 0,
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
        return ErrInvalidLimit
    }
    if q.Offset < 0 {
        return ErrInvalidOffset
    }
    return nil
}