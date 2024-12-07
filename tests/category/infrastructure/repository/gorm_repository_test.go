package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/jambo0624/blog/tests/testutil"
	categoryRepo "github.com/jambo0624/blog/internal/category/infrastructure/repository"
	categoryQuery "github.com/jambo0624/blog/internal/category/domain/query"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
)

func TestGormCategoryRepository_FindByID(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	repo := categoryRepo.NewGormCategoryRepository(testDB.DB)
	category := testDB.Data.Categories[0]

	found, err := repo.FindByID(category.ID)
	assert.NoError(t, err)
	assert.Equal(t, category.Name, found.Name)
}

func TestGormCategoryRepository_FindAll(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	repo := categoryRepo.NewGormCategoryRepository(testDB.DB)

	t.Run("with name filter", func(t *testing.T) {
		q := categoryQuery.NewCategoryQuery()
		q.WithNameLike("Technology")

		categories, total, err := repo.FindAll(q)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), total)
		assert.Contains(t, categories[0].Name, "Technology")
	})
}

func TestGormCategoryRepository_Save(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	repo := categoryRepo.NewGormCategoryRepository(testDB.DB)

	category := &categoryEntity.Category{
		Name:  "New Category",
	}

	err := repo.Save(category)
	assert.NoError(t, err)
	assert.NotZero(t, category.ID)

	found, err := repo.FindByID(category.ID)
	assert.NoError(t, err)
	assert.Equal(t, category.Name, found.Name)
}

func TestGormCategoryRepository_Update(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	repo := categoryRepo.NewGormCategoryRepository(testDB.DB)
	category := testDB.Data.Categories[0]

	category.Name = "Updated Name"
	err := repo.Update(category)
	assert.NoError(t, err)

	found, err := repo.FindByID(category.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", found.Name)
}

func TestGormCategoryRepository_Delete(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	repo := categoryRepo.NewGormCategoryRepository(testDB.DB)
	category := testDB.Data.Categories[0]

	err := repo.Delete(category.ID)
	assert.NoError(t, err)

	var found categoryEntity.Category
	err = testDB.DB.Unscoped().First(&found, category.ID).Error
	assert.NoError(t, err)
	assert.NotNil(t, found.DeletedAt)

	_, err = repo.FindByID(category.ID)
	assert.Nil(t, err)
}