package query

type CategoryQuery struct {
    IDs       []uint  // for IN query
    NameLike  string  // for name search
    SlugLike  string  // for slug search
    Limit     int     // for pagination
    Offset    int     // for pagination
    OrderBy   string  // for sorting
}

func NewCategoryQuery() *CategoryQuery {
    return &CategoryQuery{
        Limit:  10,
        Offset: 0,
    }
}

func (q *CategoryQuery) WithIDs(ids []uint) *CategoryQuery {
    q.IDs = ids
    return q
}

func (q *CategoryQuery) WithNameLike(name string) *CategoryQuery {
    q.NameLike = name
    return q
}

func (q *CategoryQuery) WithSlugLike(slug string) *CategoryQuery {
    q.SlugLike = slug
    return q
}

func (q *CategoryQuery) WithPagination(limit, offset int) *CategoryQuery {
    q.Limit = limit
    q.Offset = offset
    return q
}

func (q *CategoryQuery) WithOrderBy(orderBy string) *CategoryQuery {
    q.OrderBy = orderBy
    return q
} 