package repository_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/jambo0624/blog/internal/article/domain/entity"
    articleRepo "github.com/jambo0624/blog/internal/article/infrastructure/repository"
    "github.com/jambo0624/blog/tests/testutil"
)

func TestGormArticleRepository_Save(t *testing.T) {
    cleanup := testutil.SetupTestDB(t)
    defer cleanup()

    repo := articleRepo.NewGormArticleRepository(testutil.TestDB)

    article := &entity.Article{
        CategoryID: 1,
        Title:     "Test Article",
        Content:   "Test Content",
    }

    err := repo.Save(article)
    assert.NoError(t, err)
    assert.NotZero(t, article.ID)
}

func TestGormArticleRepository_FindByID(t *testing.T) {
    cleanup := testutil.SetupTestDB(t)
    defer cleanup()

    repo := articleRepo.NewGormArticleRepository(testutil.TestDB)

    article := &entity.Article{
        CategoryID: 1,
        Title:     "Test Article",
        Content:   "Test Content",
    }
    err := repo.Save(article)
    assert.NoError(t, err)

    found, err := repo.FindByID(article.ID)
    assert.NoError(t, err)
    assert.NotNil(t, found)
    assert.Equal(t, article.Title, found.Title)
}

func TestGormArticleRepository_Update(t *testing.T) {
    cleanup := testutil.SetupTestDB(t)
    defer cleanup()

    repo := articleRepo.NewGormArticleRepository(testutil.TestDB)

    article := &entity.Article{
        CategoryID: 1,
        Title:     "Test Article",
        Content:   "Test Content",
    }
    err := repo.Save(article)
    assert.NoError(t, err)

    article.Title = "Updated Title"
    err = repo.Update(article)
    assert.NoError(t, err)

    found, err := repo.FindByID(article.ID)
    assert.NoError(t, err)
    assert.Equal(t, "Updated Title", found.Title)
}

func TestGormArticleRepository_Delete(t *testing.T) {
    cleanup := testutil.SetupTestDB(t)
    defer cleanup()

    repo := articleRepo.NewGormArticleRepository(testutil.TestDB)

    article := &entity.Article{
        CategoryID: 1,
        Title:     "Test Article",
        Content:   "Test Content",
    }
    err := repo.Save(article)
    assert.NoError(t, err)

    err = repo.Delete(article.ID)
    assert.NoError(t, err)

    found, err := repo.FindByID(article.ID)
    assert.Error(t, err)
    assert.Nil(t, found)
} 