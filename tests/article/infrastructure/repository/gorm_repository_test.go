package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/jambo0624/blog/tests/testutil"
	articleRepo "github.com/jambo0624/blog/internal/article/infrastructure/repository"
	articleQuery "github.com/jambo0624/blog/internal/article/domain/query"
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	factory "github.com/jambo0624/blog/tests/testutil/factory"
)

func TestGormArticleRepository_FindByID(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	repo := articleRepo.NewGormArticleRepository(testDB.DB)
	article := testDB.Data.Articles[0]

	found, err := repo.FindByID(article.ID)
	assert.NoError(t, err)
	assert.Equal(t, article.Title, found.Title)
	assert.Equal(t, article.Content, found.Content)
	assert.Equal(t, article.CategoryID, found.CategoryID)
}

func TestGormArticleRepository_FindAll(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	repo := articleRepo.NewGormArticleRepository(testDB.DB)

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
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	repo := articleRepo.NewGormArticleRepository(testDB.DB)
	category := testDB.Data.Categories[0]
	tag := testDB.Data.Tags[0]
	categoryFactory := factory.NewCategoryFactory()
	tagFactory := factory.NewTagFactory()
	articleFactory := factory.NewArticleFactory(categoryFactory, tagFactory)

	article := articleFactory.BuildEntity(
		func(a *articleEntity.Article) {
			a.CategoryID = category.ID
			a.Tags = []tagEntity.Tag{*tag}
		},
	)

	err := repo.Save(article)
	assert.NoError(t, err)
	assert.NotZero(t, article.ID)

	// Verify saved article
	preloads := []string{"Tags"}
	found, err := repo.FindByID(article.ID, preloads...)
	assert.NoError(t, err)
	assert.Equal(t, article.Title, found.Title)
	assert.Equal(t, article.Content, found.Content)
	assert.Len(t, found.Tags, 1)
	assert.Equal(t, tag.ID, found.Tags[0].ID)
}

func TestGormArticleRepository_Update(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	repo := articleRepo.NewGormArticleRepository(testDB.DB)
	article := testDB.Data.Articles[0]

	article.Title = "Updated Title"
	err := repo.Update(article)
	assert.NoError(t, err)

	found, err := repo.FindByID(article.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", found.Title)
}

func TestGormArticleRepository_Delete(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	repo := articleRepo.NewGormArticleRepository(testDB.DB)
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