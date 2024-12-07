package query

import (
	"gorm.io/gorm"

	"github.com/jambo0624/blog/internal/shared/domain/constants"
	baseQuery "github.com/jambo0624/blog/internal/shared/domain/query"
	"github.com/jambo0624/blog/internal/shared/domain/validate"
)

// Preload constants for Article queries
const (
	PreloadCategory = "Category"
	PreloadTags     = "Tags"
)
	
type ArticleQuery struct {
	baseQuery.BaseQuery
	CategoryID          *uint    // article specific field
	TagIDs              []uint   // article specific field
	TitleLike           string   // article specific field
	ContentLike         string   // article specific field
	PreloadAssociations []string // store associations to be preloaded
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
	if err := q.BaseQuery.Validate(); err != nil {
		return err
	}
	if len(q.TitleLike) > constants.MaxTitleLength {
		return validate.ErrTitleTooLong
	}
	return nil
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
	
	return db
}


// getDefaultPreloads returns default preload associations for Article queries
func getDefaultPreloads() []string {
	return []string{PreloadCategory, PreloadTags}
}