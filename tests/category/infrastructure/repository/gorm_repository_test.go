package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/jambo0624/blog/tests/testutil"
	categoryPersistence "github.com/jambo0624/blog/internal/category/infrastructure/repository"
	categoryRepository "github.com/jambo0624/blog/internal/category/domain/repository"
	categoryQuery "github.com/jambo0624/blog/internal/category/domain/query"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	factory "github.com/jambo0624/blog/tests/testutil/factory"
)

func setupTest(t *testing.T) (*testutil.TestDB, func(), categoryRepository.CategoryRepository, *factory.CategoryFactory) {
	t.Helper()
	
	testDB, cleanup := testutil.SetupTestDB(t)
	repo := categoryPersistence.NewGormCategoryRepository(testDB.DB)
	factory := factory.NewCategoryFactory()

	return testDB, cleanup, repo, factory
}

func TestGormCategoryRepository_FindByID(t *testing.T) {
	testDB, cleanup, repo, _ := setupTest(t)
	defer cleanup()
	category := testDB.Data.Categories[0]

	found, err := repo.FindByID(category.ID)
	assert.NoError(t, err)
	assert.Equal(t, category.Name, found.Name)
}

func TestGormCategoryRepository_FindAll(t *testing.T) {
	testDB, cleanup, repo, _ := setupTest(t)
	defer cleanup()

	t.Run("with name filter", func(t *testing.T) {
		q := categoryQuery.NewCategoryQuery()
		name := testDB.Data.Categories[0].Name
		q.WithNameLike(name)

		categories, total, err := repo.FindAll(q)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), total)
		assert.Contains(t, categories[0].Name, name)
	})
}

func TestGormCategoryRepository_Save(t *testing.T) {
	_, cleanup, repo, factory := setupTest(t)
	defer cleanup()

	category := factory.BuildEntity()

	err := repo.Save(category)
	assert.NoError(t, err)
	assert.NotZero(t, category.ID)

	found, err := repo.FindByID(category.ID)
	assert.NoError(t, err)
	assert.Equal(t, category.Name, found.Name)
}

func TestGormCategoryRepository_Update(t *testing.T) {
	testDB, cleanup, repo, _ := setupTest(t)
	defer cleanup()

	category := testDB.Data.Categories[0]

	category.Name = "Updated Name"
	err := repo.Update(category)
	assert.NoError(t, err)

	found, err := repo.FindByID(category.ID)
	assert.NoError(t, err)
	assert.Equal(t, category.Name, found.Name)
}

func TestGormCategoryRepository_Delete(t *testing.T) {
	testDB, cleanup, repo, _ := setupTest(t)
	defer cleanup()
	
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

func TestGormCategoryRepository_FindAll_WithFilters(t *testing.T) {
	testDB, cleanup, repo, factory := setupTest(t)
	defer cleanup()

	// Create test data
	category1 := factory.BuildEntity(factory.WithName("Test1"))
	category2 := factory.BuildEntity(factory.WithName("Test2"))
	testDB.DB.Create(category1)
	testDB.DB.Create(category2)

	tests := []struct {
		name          string
		buildQuery    func() *categoryQuery.CategoryQuery
		expectedCount int64
	}{
		{
			name: "filter by name",
			buildQuery: func() *categoryQuery.CategoryQuery {
				q := categoryQuery.NewCategoryQuery()
				q.WithNameLike("Test1")
				return q
			},
			expectedCount: 1,
		},
		{
			name: "no filter",
			buildQuery: func() *categoryQuery.CategoryQuery {
				return categoryQuery.NewCategoryQuery()
			},
			expectedCount: 4, // includes default categories
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			categories, count, err := repo.FindAll(tt.buildQuery())
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCount, count)
			assert.Len(t, categories, int(tt.expectedCount))
		})
	}
}