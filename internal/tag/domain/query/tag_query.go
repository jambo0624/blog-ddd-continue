package query

type TagQuery struct {
    IDs        []uint  // for IN query
    NameLike   string  // for LIKE query
    Limit      int     // for pagination
    Offset     int     // for pagination
    OrderBy    string  // for sorting
}

// NewTagQuery creates a new query with default values
func NewTagQuery() *TagQuery {
    return &TagQuery{
        Limit:  10,    // default limit
        Offset: 0,
    }
}

// WithIDs adds IN condition
func (q *TagQuery) WithIDs(ids []uint) *TagQuery {
    q.IDs = ids
    return q
}

// WithNameLike adds LIKE condition
func (q *TagQuery) WithNameLike(name string) *TagQuery {
    q.NameLike = name
    return q
}

// WithPagination adds pagination
func (q *TagQuery) WithPagination(limit, offset int) *TagQuery {
    q.Limit = limit
    q.Offset = offset
    return q
}

// WithOrderBy adds sorting
func (q *TagQuery) WithOrderBy(orderBy string) *TagQuery {
    q.OrderBy = orderBy
    return q
} 