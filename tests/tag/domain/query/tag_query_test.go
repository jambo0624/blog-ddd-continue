package query_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/jambo0624/blog/internal/tag/domain/entity"
	"github.com/jambo0624/blog/internal/tag/domain/query"
	"github.com/jambo0624/blog/tests/testutil"
)

func TestTagQuery_Validate(t *testing.T) {
	tests := []struct {
		name    string
		query   func() *query.TagQuery
		wantErr bool
	}{
		{
			name: "valid query",
			query: func() *query.TagQuery {
				q := query.NewTagQuery()
				q.WithNameLike("test")
				return q
			},
			wantErr: false,
		},
		{
			name: "invalid name length",
			query: func() *query.TagQuery {
				q := query.NewTagQuery()
				q.WithNameLike(strings.Repeat("a", 101))
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

func TestTagQuery_ApplyFilters(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	tests := []struct {
		name            string
		setupQuery      func() *query.TagQuery
		expectedClauses []string
	}{
		{
			name: "with name filter",
			setupQuery: func() *query.TagQuery {
				q := query.NewTagQuery()
				q.WithNameLike("test")
				return q
			},
			expectedClauses: []string{
				"name LIKE '%test%'",
			},
		},
		{
			name: "with color filter",
			setupQuery: func() *query.TagQuery {
				q := query.NewTagQuery()
				q.WithColorLike("#FF0000")
				return q
			},
			expectedClauses: []string{
				"color LIKE '%#FF0000%'",
			},
		},
		{
			name: "with multiple filters",
			setupQuery: func() *query.TagQuery {
				q := query.NewTagQuery()
				q.WithNameLike("test")
				q.WithColorLike("#FF0000")
				return q
			},
			expectedClauses: []string{
				"name LIKE '%test%'",
				"color LIKE '%#FF0000%'",
			},
		},
		{
			name: "with IDs filter",
			setupQuery: func() *query.TagQuery {
				q := query.NewTagQuery()
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
			db := testDB.DB.Model(&entity.Tag{})
			db = q.ApplyFilters(db)

			// use ToSQL to get the query string
			sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
				var tags []*entity.Tag
				return tx.Find(&tags)
			})

			t.Logf("Generated SQL: %s", sql)
			for _, clause := range tt.expectedClauses {
				assert.Contains(t, sql, clause)
			}
		})
	}
}
