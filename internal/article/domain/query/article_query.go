package query

type ArticleQuery struct {
    IDs         []uint  // for IN query
    CategoryID  *uint   // for category filter
    TagIDs      []uint  // for tag filter
    TitleLike   string  // for title search
    ContentLike string  // for content search
    Limit       int     // for pagination
    Offset      int     // for pagination
    OrderBy     string  // for sorting
}

func NewArticleQuery() *ArticleQuery {
    return &ArticleQuery{
        Limit:  10,
        Offset: 0,
    }
}

func (q *ArticleQuery) WithIDs(ids []uint) *ArticleQuery {
    q.IDs = ids
    return q
}

func (q *ArticleQuery) WithCategoryID(id uint) *ArticleQuery {
    q.CategoryID = &id
    return q
}

func (q *ArticleQuery) WithTagIDs(ids []uint) *ArticleQuery {
    q.TagIDs = ids
    return q
}

func (q *ArticleQuery) WithTitleLike(title string) *ArticleQuery {
    q.TitleLike = title
    return q
}

func (q *ArticleQuery) WithContentLike(content string) *ArticleQuery {
    q.ContentLike = content
    return q
}

func (q *ArticleQuery) WithPagination(limit, offset int) *ArticleQuery {
    q.Limit = limit
    q.Offset = offset
    return q
}

func (q *ArticleQuery) WithOrderBy(orderBy string) *ArticleQuery {
    q.OrderBy = orderBy
    return q
} 