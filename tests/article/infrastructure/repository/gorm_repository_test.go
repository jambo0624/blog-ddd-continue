package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/jambo0624/blog/tests/testutil"
	articlePersistence "github.com/jambo0624/blog/internal/article/infrastructure/repository"
	articleRepository "github.com/jambo0624/blog/internal/article/domain/repository"
	articleQuery "github.com/jambo0624/blog/internal/article/domain/query"
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
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
	assert.NoError(t, err)
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
		assert.NoError(t, err)
		assert.Equal(t, int64(1), total)
		assert.Equal(t, testDB.Data.Articles[0].Title, articles[0].Title)
	})

	t.Run("with tag filter", func(t *testing.T) {
		q := articleQuery.NewArticleQuery()
		q.WithTagIDs([]uint{testDB.Data.Tags[0].ID})
		q.PreloadAssociations = []string{"Tags"}

		articles, total, err := repo.FindAll(q)
		assert.NoError(t, err)
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
	assert.NoError(t, err)

	// Create Tag
	err = testDB.DB.Create(tag).Error
	assert.NoError(t, err)

	// Save Article
	err = repo.Save(article)
	assert.NoError(t, err)
	assert.NotZero(t, article.ID)

	// Verify saved article
	preloads := []string{"Tags", "Category"}
	found, err := repo.FindByID(article.ID, preloads...)
	assert.NoError(t, err)
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
	assert.NoError(t, err)

	found, err := repo.FindByID(article.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", found.Title)
}

func TestGormArticleRepository_Delete(t *testing.T) {
	testDB, cleanup, repo, _ := setupTest(t)
	defer cleanup()

	article := testDB.Data.Articles[0]

	err := repo.Delete(article.ID)
	assert.NoError(t, err)

	var found *articleEntity.Article
	err = testDB.DB.Unscoped().First(&found, article.ID).Error
	assert.NoError(t, err)
	assert.NotNil(t, found.DeletedAt)

	_, err = repo.FindByID(article.ID)
	assert.Nil(t, err)
} 