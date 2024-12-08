package query_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/jambo0624/blog/internal/category/domain/entity"
	"github.com/jambo0624/blog/internal/category/domain/query"
	"github.com/jambo0624/blog/tests/testutil"
)

func TestCategoryQuery_Validate(t *testing.T) {
	tests := []struct {
		name    string
		query   func() *query.CategoryQuery
		wantErr bool
	}{
		{
			name: "valid query",
			query: func() *query.CategoryQuery {
				q := query.NewCategoryQuery()
				q.WithNameLike("test")
				return q
			},
			wantErr: false,
		},
		{
			name: "invalid name length",
			query: func() *query.CategoryQuery {
				q := query.NewCategoryQuery()
				q.WithNameLike(strings.Repeat("a", 101))
				return q
			},
			wantErr: true,
		},
		{
			name: "invalid slug length",
			query: func() *query.CategoryQuery {
				q := query.NewCategoryQuery()
				q.WithSlugLike(strings.Repeat("a", 101))
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

func TestCategoryQuery_ApplyFilters(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	tests := []struct {
		name            string
		setupQuery      func() *query.CategoryQuery
		expectedClauses []string
	}{
		{
			name: "with name filter",
			setupQuery: func() *query.CategoryQuery {
				q := query.NewCategoryQuery()
				q.WithNameLike("test")
				return q
			},
			expectedClauses: []string{
				"name LIKE '%test%'",
			},
		},
		{
			name: "with slug filter",
			setupQuery: func() *query.CategoryQuery {
				q := query.NewCategoryQuery()
				q.WithSlugLike("test-slug")
				return q
			},
			expectedClauses: []string{
				"slug LIKE '%test-slug%'",
			},
		},
		{
			name: "with multiple filters",
			setupQuery: func() *query.CategoryQuery {
				q := query.NewCategoryQuery()
				q.WithNameLike("test")
				q.WithSlugLike("test-slug")
				return q
			},
			expectedClauses: []string{
				"name LIKE '%test%'",
				"slug LIKE '%test-slug%'",
			},
		},
		{
			name: "with IDs filter",
			setupQuery: func() *query.CategoryQuery {
				q := query.NewCategoryQuery()
				q.WithIDs([]uint{1, 2})
				return q
			},
			expectedClauses: []string{
				"id IN (1,2)",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := tt.setupQuery()
			db := testDB.DB.Model(&entity.Category{})
			db = q.ApplyFilters(db)

			// use ToSQL to get the query string
			sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
				var categories []*entity.Category
				return tx.Find(&categories)
			})

			t.Logf("Generated SQL: %s", sql)
			for _, clause := range tt.expectedClauses {
				assert.Contains(t, sql, clause)
			}
		})
	}
}
