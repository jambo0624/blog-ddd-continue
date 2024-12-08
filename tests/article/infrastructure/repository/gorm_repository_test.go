package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	articleQuery "github.com/jambo0624/blog/internal/article/domain/query"
	articleRepository "github.com/jambo0624/blog/internal/article/domain/repository"
	articlePersistence "github.com/jambo0624/blog/internal/article/infrastructure/repository"
	"github.com/jambo0624/blog/tests/testutil"
	factory "github.com/jambo0624/blog/tests/testutil/factory"
)

func setupTest(t *testing.T) (*testutil.TestDB, func(), articleRepository.ArticleRepository, *factory.ArticleFactory) {
	t.Helper()

	testDB, cleanup := testutil.SetupTestDB(t)
	repo := articlePersistence.NewGormArticleRepository(testDB.DB)
	categoryFactory := factory.NewCategoryFactory()
	tagFactory := factory.NewTagFactory()
	factory := factory.NewArticleFactory(categoryFactory, tagFactory)

	return testDB, cleanup, repo, factory
}

func TestGormArticleRepository_FindByID(t *testing.T) {
	testDB, cleanup, repo, _ := setupTest(t)
	defer cleanup()

	article := testDB.Data.Articles[0]

	found, err := repo.FindByID(article.ID)
	require.NoError(t, err)
	assert.Equal(t, article.Title, found.Title)
	assert.Equal(t, article.Content, found.Content)
	assert.Equal(t, article.CategoryID, found.CategoryID)
}

func TestGormArticleRepository_FindAll(t *testing.T) {
	testDB, cleanup, repo, _ := setupTest(t)
	defer cleanup()

	t.Run("with category filter", func(t *testing.T) {
		q := articleQuery.NewArticleQuery()
		q.WithCategoryID(testDB.Data.Categories[0].ID)

		articles, total, err := repo.FindAll(q)
		require.NoError(t, err)
		assert.Equal(t, int64(1), total)
		assert.Equal(t, testDB.Data.Articles[0].Title, articles[0].Title)
	})

	t.Run("with tag filter", func(t *testing.T) {
		q := articleQuery.NewArticleQuery()
		q.WithTagIDs([]uint{testDB.Data.Tags[0].ID})
		q.PreloadAssociations = []string{"Tags"}

		articles, total, err := repo.FindAll(q)
		require.NoError(t, err)
		assert.Equal(t, int64(1), total)
		assert.Len(t, articles, 1)
		assert.Equal(t, testDB.Data.Articles[0].Title, articles[0].Title)
		assert.Len(t, articles[0].Tags, 1)
		assert.Equal(t, testDB.Data.Tags[0].ID, articles[0].Tags[0].ID)
	})
}

func TestGormArticleRepository_Save(t *testing.T) {
	testDB, cleanup, repo, factory := setupTest(t)
	defer cleanup()

	// Build Entity, contains 2 tags
	article, category, tag := factory.BuildEntity()

	// Create Category
	err := testDB.DB.Create(category).Error
	require.NoError(t, err)

	// Create Tag
	err = testDB.DB.Create(tag).Error
	require.NoError(t, err)

	// Save Article
	err = repo.Save(article)
	require.NoError(t, err)
	require.NotZero(t, article.ID)

	// Verify saved article
	preloads := []string{"Tags", "Category"}
	found, err := repo.FindByID(article.ID, preloads...)
	require.NoError(t, err)
	assert.Equal(t, article.Title, found.Title)
	assert.Equal(t, article.Content, found.Content)
	assert.Equal(t, category.ID, found.CategoryID)
	assert.Len(t, found.Tags, 2)
	assert.Equal(t, tag.ID, found.Tags[0].ID)
}

func TestGormArticleRepository_Update(t *testing.T) {
	testDB, cleanup, repo, _ := setupTest(t)
	defer cleanup()

	article := testDB.Data.Articles[0]

	article.Title = "Updated Title"
	err := repo.Update(article)
	require.NoError(t, err)

	found, err := repo.FindByID(article.ID)
	require.NoError(t, err)
	assert.Equal(t, "Updated Title", found.Title)
}

func TestGormArticleRepository_Delete(t *testing.T) {
	testDB, cleanup, repo, _ := setupTest(t)
	defer cleanup()

	article := testDB.Data.Articles[0]

	err := repo.Delete(article.ID)
	require.NoError(t, err)

	var found *articleEntity.Article
	err = testDB.DB.Unscoped().First(&found, article.ID).Error
	require.NoError(t, err)
	require.NotNil(t, found.DeletedAt)

	_, err = repo.FindByID(article.ID)
	require.NoError(t, err)
}

// buildTestCase is a helper function to create test cases.
func buildTestCase(
	t *testing.T,
	name string,
	buildQueryFn func() *articleQuery.ArticleQuery,
	expectedCount int64,
	validate func(t *testing.T, articles []*articleEntity.Article),
) struct {
	name          string
	buildQuery    func() *articleQuery.ArticleQuery
	expectedCount int64
	validate      func(t *testing.T, articles []*articleEntity.Article)
} {
	t.Helper()
	return struct {
		name          string
		buildQuery    func() *articleQuery.ArticleQuery
		expectedCount int64
		validate      func(t *testing.T, articles []*articleEntity.Article)
	}{
		name:          name,
		buildQuery:    buildQueryFn,
		expectedCount: expectedCount,
		validate:      validate,
	}
}

// buildFilterTestCase creates a test case for simple filter tests.
func buildFilterTestCase[T any](
	t *testing.T,
	name string,
	filterValue T,
	buildFilterFn func(T) func(*articleQuery.ArticleQuery),
	validate func(t *testing.T, articles []*articleEntity.Article),
	expectedCount int64,
) struct {
	name          string
	buildQuery    func() *articleQuery.ArticleQuery
	expectedCount int64
	validate      func(t *testing.T, articles []*articleEntity.Article)
} {
	t.Helper()
	return buildTestCase(
		t,
		name,
		func() *articleQuery.ArticleQuery {
			q := articleQuery.NewArticleQuery()
			buildFilterFn(filterValue)(q)
			return q
		},
		expectedCount,
		validate,
	)
}

func TestGormArticleRepository_FindAll_WithFilters(t *testing.T) {
	testDB, cleanup, repo, factory := setupTest(t)
	defer cleanup()

	// Create test data
	article1, category1, tag1 := factory.BuildEntity(
		factory.WithTitle("Test1"),
		factory.WithContent("Content1"),
	)
	article2, category2, tag2 := factory.BuildEntity(
		factory.WithTitle("Test2"),
		factory.WithContent("Content2"),
	)

	// Create categories first
	err := testDB.DB.Create(category1).Error
	require.NoError(t, err)
	err = testDB.DB.Create(category2).Error
	require.NoError(t, err)

	// Create tags
	err = testDB.DB.Create(tag1).Error
	require.NoError(t, err)
	err = testDB.DB.Create(tag2).Error
	require.NoError(t, err)

	// Create articles
	err = testDB.DB.Create(article1).Error
	require.NoError(t, err)
	err = testDB.DB.Create(article2).Error
	require.NoError(t, err)

	tests := []struct {
		name          string
		buildQuery    func() *articleQuery.ArticleQuery
		expectedCount int64
		validate      func(t *testing.T, articles []*articleEntity.Article)
	}{
		buildFilterTestCase[string](
			t,
			"filter by title",
			"Test1",
			func(value string) func(*articleQuery.ArticleQuery) {
				return func(q *articleQuery.ArticleQuery) {
					q.WithTitleLike(value)
				}
			},
			func(t *testing.T, articles []*articleEntity.Article) {
				t.Helper()
				assert.Equal(t, "Test1", articles[0].Title)
			},
			1,
		),
		buildFilterTestCase[string](
			t,
			"filter by content",
			"Content2",
			func(value string) func(*articleQuery.ArticleQuery) {
				return func(q *articleQuery.ArticleQuery) {
					q.WithContentLike(value)
				}
			},
			func(t *testing.T, articles []*articleEntity.Article) {
				t.Helper()
				assert.Equal(t, "Test2", articles[0].Title)
			},
			1,
		),
		buildFilterTestCase[uint](
			t,
			"filter by category",
			article1.CategoryID,
			func(value uint) func(*articleQuery.ArticleQuery) {
				return func(q *articleQuery.ArticleQuery) {
					q.WithCategoryID(value)
				}
			},
			func(t *testing.T, articles []*articleEntity.Article) {
				t.Helper()
				assert.Equal(t, article1.CategoryID, articles[0].CategoryID)
			},
			1,
		),
		buildFilterTestCase[uint](
			t,
			"filter by tags",
			article1.Tags[0].ID,
			func(value uint) func(*articleQuery.ArticleQuery) {
				return func(q *articleQuery.ArticleQuery) {
					q.WithTagIDs([]uint{value})
				}
			},
			func(t *testing.T, articles []*articleEntity.Article) {
				t.Helper()
				assert.Equal(t, article1.Tags[0].ID, articles[0].Tags[0].ID)
			},
			1,
		),
		buildFilterTestCase[string](
			t,
			"with multiple filters",
			"Test",
			func(value string) func(*articleQuery.ArticleQuery) {
				return func(q *articleQuery.ArticleQuery) {
					q.WithTitleLike(value)
					q.WithCategoryID(article1.CategoryID)
				}
			},
			func(t *testing.T, articles []*articleEntity.Article) {
				t.Helper()
				assert.Equal(t, article1.Title, articles[0].Title)
				assert.Equal(t, article1.CategoryID, articles[0].CategoryID)
			},
			1,
		),
		buildTestCase(
			t,
			"no filter",
			articleQuery.NewArticleQuery,
			4, // includes default articles
			func(t *testing.T, articles []*articleEntity.Article) {
				t.Helper()
				assert.Len(t, articles, 4)
			},
		),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			articles, count, err := repo.FindAll(tt.buildQuery())
			require.NoError(t, err)
			assert.Equal(t, tt.expectedCount, count)
			if tt.validate != nil {
				tt.validate(t, articles)
			}
		})
	}
}
