package query_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/jambo0624/blog/internal/article/domain/entity"
	"github.com/jambo0624/blog/internal/article/domain/query"
	"github.com/jambo0624/blog/tests/testutil"
)

func TestArticleQuery_Validate(t *testing.T) {
	tests := []struct {
		name        string
		query       func() *query.ArticleQuery
		wantErr     bool
		errContains string
	}{
		{
			name: "valid query",
			query: func() *query.ArticleQuery {
				q := query.NewArticleQuery()
				q.WithTitleLike("test")
				q.WithCategoryID(1)
				q.WithTagIDs([]uint{1, 2})
				return q
			},
			wantErr: false,
		},
		{
			name: "invalid title length",
			query: func() *query.ArticleQuery {
				q := query.NewArticleQuery()
				q.WithTitleLike(strings.Repeat("a", 256))
				return q
			},
			wantErr: true,
		},
		{
			name: "invalid content length",
			query: func() *query.ArticleQuery {
				q := query.NewArticleQuery()
				q.WithContentLike(strings.Repeat("a", 256))
				return q
			},
			wantErr: true,
		},
		{
			name: "invalid limit",
			query: func() *query.ArticleQuery {
				q := query.NewArticleQuery()
				q.Limit = -1
				return q
			},
			wantErr: true,
		},
		{
			name: "invalid offset",
			query: func() *query.ArticleQuery {
				q := query.NewArticleQuery()
				q.Offset = -1
				return q
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.query().Validate()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestArticleQuery_ApplyFilters(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	tests := []struct {
		name            string
		setupQuery      func() *query.ArticleQuery
		expectedClauses []string
	}{
		{
			name: "with category filter",
			setupQuery: func() *query.ArticleQuery {
				q := query.NewArticleQuery()
				q.WithCategoryID(1)
				return q
			},
			expectedClauses: []string{
				"category_id = 1",
			},
		},
		{
			name: "with tag filter",
			setupQuery: func() *query.ArticleQuery {
				q := query.NewArticleQuery()
				q.WithTagIDs([]uint{1, 2})
				return q
			},
			expectedClauses: []string{
				"LEFT JOIN article_tags ON articles.id = article_tags.article_id",
				"article_tags.tag_id IN (1,2)",
			},
		},
		{
			name: "with title filter",
			setupQuery: func() *query.ArticleQuery {
				q := query.NewArticleQuery()
				q.WithTitleLike("test")
				return q
			},
			expectedClauses: []string{
				"title LIKE '%test%'",
			},
		},
		{
			name: "with content filter",
			setupQuery: func() *query.ArticleQuery {
				q := query.NewArticleQuery()
				q.WithContentLike("test")
				return q
			},
			expectedClauses: []string{
				"content LIKE '%test%'",
			},
		},
		{
			name: "with multiple filters",
			setupQuery: func() *query.ArticleQuery {
				q := query.NewArticleQuery()
				q.WithCategoryID(1)
				q.WithTagIDs([]uint{1})
				q.WithTitleLike("test")
				return q
			},
			expectedClauses: []string{
				"category_id = 1",
				"LEFT JOIN article_tags ON articles.id = article_tags.article_id",
				"article_tags.tag_id IN (1)",
				"title LIKE '%test%'",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.setupQuery()
			db := testDB.DB.Model(&entity.Article{})
			db = q.ApplyFilters(db)

			// use ToSQL to get the query string
			sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
				var articles []*entity.Article
				return tx.Find(&articles)
			})

			t.Logf("Generated SQL: %s", sql)
			for _, clause := range tt.expectedClauses {
				assert.Contains(t, sql, clause)
			}
		})
	}
}
