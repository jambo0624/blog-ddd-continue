package repository_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/jambo0624/blog/internal/tag/domain/entity"
    tagRepo "github.com/jambo0624/blog/internal/tag/infrastructure/repository"
    "github.com/jambo0624/blog/tests/testutil"
)

func TestGormTagRepository_Save(t *testing.T) {
    cleanup := testutil.SetupTestDB(t)
    defer cleanup()

    repo := tagRepo.NewGormTagRepository(testutil.TestDB)

    tag := &entity.Tag{
        Name:  "Test Tag",
        Color: "#FF0000",
    }

    err := repo.Save(tag)
    assert.NoError(t, err)
    assert.NotZero(t, tag.ID)
}

func TestGormTagRepository_FindByName(t *testing.T) {
    cleanup := testutil.SetupTestDB(t)
    defer cleanup()

    repo := tagRepo.NewGormTagRepository(testutil.TestDB)

    tag := &entity.Tag{
        Name:  "Test Tag",
        Color: "#FF0000",
    }
    err := repo.Save(tag)
    assert.NoError(t, err)

    found, err := repo.FindByName("Test Tag")
    assert.NoError(t, err)
    assert.NotNil(t, found)
    assert.Equal(t, tag.Color, found.Color)
}

// ... 其他测试用例 