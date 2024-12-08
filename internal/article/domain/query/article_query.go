package query

import (
	"gorm.io/gorm"

	baseQuery "github.com/jambo0624/blog/internal/shared/domain/query"
)

// Preload constants for Article queries.
const (
	PreloadCategory = "Category"
	PreloadTags     = "Tags"
)

type ArticleQuery struct {
	baseQuery.BaseQuery
	CategoryID          *uint    `binding:"omitempty"          json:"categoryId"          validate:"omitempty,gt=0"`
	TagIDs              []uint   `binding:"omitempty"          json:"tagIds"              validate:"omitempty,dive,gt=0"`
	TitleLike           string   `binding:"omitempty"          json:"titleLike"           validate:"omitempty,max=255"`
	ContentLike         string   `binding:"omitempty, max=255" json:"contentLike"         validate:"omitempty,max=255"`
	PreloadAssociations []string `binding:"omitempty"          json:"preloadAssociations"`
}

func NewArticleQuery() *ArticleQuery {
	return &ArticleQuery{
		BaseQuery:           baseQuery.NewBaseQuery(),
		PreloadAssociations: getDefaultPreloads(),
	}
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

func (q *ArticleQuery) Validate() error {
	return q.BaseQuery.ValidateQuery(q)
}

func (q *ArticleQuery) GetBaseQuery() baseQuery.BaseQuery {
	return q.BaseQuery
}

func (q *ArticleQuery) GetPreloadAssociations() []string {
	return q.PreloadAssociations
}

func (q *ArticleQuery) ApplyFilters(db *gorm.DB) *gorm.DB {
	if q.CategoryID != nil {
		db = db.Where("category_id = ?", q.CategoryID)
	}

	if len(q.TagIDs) > 0 {
		db = db.Joins("LEFT JOIN article_tags ON articles.id = article_tags.article_id").
			Where("article_tags.tag_id IN ?", q.TagIDs).
			Group("articles.id")
	}

	if q.TitleLike != "" {
		db = db.Where("title LIKE ?", "%"+q.TitleLike+"%")
	}

	if q.ContentLike != "" {
		db = db.Where("content LIKE ?", "%"+q.ContentLike+"%")
	}

	return db
}

// getDefaultPreloads returns default preload associations for Article queries.
func getDefaultPreloads() []string {
	return []string{PreloadCategory, PreloadTags}
}
