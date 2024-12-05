package repository_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/jambo0624/blog/internal/category/domain/entity"
    categoryRepo "github.com/jambo0624/blog/internal/category/infrastructure/repository"
    "github.com/jambo0624/blog/tests/testutil"
)

func TestGormCategoryRepository_Save(t *testing.T) {
    cleanup := testutil.SetupTestDB(t)
    defer cleanup()

    repo := categoryRepo.NewGormCategoryRepository(testutil.TestDB)

    category := &entity.Category{
        Name: "Test Category",
        Slug: "test-category",
    }

    err := repo.Save(category)
    assert.NoError(t, err)
    assert.NotZero(t, category.ID)
}

func TestGormCategoryRepository_FindBySlug(t *testing.T) {
    cleanup := testutil.SetupTestDB(t)
    defer cleanup()

    repo := categoryRepo.NewGormCategoryRepository(testutil.TestDB)

    category := &entity.Category{
        Name: "Test Category",
        Slug: "test-category",
    }
    err := repo.Save(category)
    assert.NoError(t, err)

    found, err := repo.FindBySlug("test-category")
    assert.NoError(t, err)
    assert.NotNil(t, found)
    assert.Equal(t, category.Name, found.Name)
}

// ... 其他测试用例 