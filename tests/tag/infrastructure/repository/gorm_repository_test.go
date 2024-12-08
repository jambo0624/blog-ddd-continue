package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	tagQuery "github.com/jambo0624/blog/internal/tag/domain/query"
	tagRepository "github.com/jambo0624/blog/internal/tag/domain/repository"
	tagPersistence "github.com/jambo0624/blog/internal/tag/infrastructure/repository"
	"github.com/jambo0624/blog/tests/testutil"
	factory "github.com/jambo0624/blog/tests/testutil/factory"
)

func setupTest(t *testing.T) (*testutil.TestDB, func(), tagRepository.TagRepository, *factory.TagFactory) {
	t.Helper()

	testDB, cleanup := testutil.SetupTestDB(t)
	repo := tagPersistence.NewGormTagRepository(testDB.DB)
	factory := factory.NewTagFactory()

	return testDB, cleanup, repo, factory
}

func TestGormTagRepository_FindByID(t *testing.T) {
	testDB, cleanup, repo, _ := setupTest(t)
	defer cleanup()

	tag := testDB.Data.Tags[0]

	found, err := repo.FindByID(tag.ID)
	require.NoError(t, err)
	assert.Equal(t, tag.Name, found.Name)
	assert.Equal(t, tag.Color, found.Color)
}

func TestGormTagRepository_FindAll(t *testing.T) {
	testDB, cleanup, repo, _ := setupTest(t)
	defer cleanup()

	t.Run("with name filter", func(t *testing.T) {
		q := tagQuery.NewTagQuery()
		name := testDB.Data.Tags[0].Name
		q.WithNameLike(name)

		tags, total, err := repo.FindAll(q)
		require.NoError(t, err)
		assert.Equal(t, int64(1), total)
		assert.Contains(t, tags[0].Name, name)
	})
}

func TestGormTagRepository_Save(t *testing.T) {
	_, cleanup, repo, factory := setupTest(t)
	defer cleanup()

	tag := factory.BuildEntity()

	err := repo.Save(tag)
	require.NoError(t, err)
	assert.NotZero(t, tag.ID)

	found, err := repo.FindByID(tag.ID)
	require.NoError(t, err)
	assert.Equal(t, tag.Name, found.Name)
	assert.Equal(t, tag.Color, found.Color)
}

func TestGormTagRepository_Update(t *testing.T) {
	testDB, cleanup, repo, _ := setupTest(t)
	defer cleanup()

	tag := testDB.Data.Tags[0]
	tag.Name = "Updated Name"

	err := repo.Update(tag)
	require.NoError(t, err)

	found, err := repo.FindByID(tag.ID)
	require.NoError(t, err)
	assert.Equal(t, tag.Name, found.Name)
}

func TestGormTagRepository_Delete(t *testing.T) {
	testDB, cleanup, repo, _ := setupTest(t)
	defer cleanup()

	tag := testDB.Data.Tags[0]

	err := repo.Delete(tag.ID)
	require.NoError(t, err)

	var found tagEntity.Tag
	err = testDB.DB.Unscoped().First(&found, tag.ID).Error
	require.NoError(t, err)
	assert.NotNil(t, found.DeletedAt)

	_, err = repo.FindByID(tag.ID)
	require.NoError(t, err)
}

func TestGormTagRepository_FindAll_WithFilters(t *testing.T) {
	testDB, cleanup, repo, factory := setupTest(t)
	defer cleanup()

	// Create test data
	tag1 := factory.BuildEntity(factory.WithName("Test1"))
	tag2 := factory.BuildEntity(factory.WithName("Test2"))
	testDB.DB.Create(tag1)
	testDB.DB.Create(tag2)

	tests := []struct {
		name          string
		buildQuery    func() *tagQuery.TagQuery
		expectedCount int64
	}{
		{
			name: "filter by name",
			buildQuery: func() *tagQuery.TagQuery {
				q := tagQuery.NewTagQuery()
				q.WithNameLike("Test1")
				return q
			},
			expectedCount: 1,
		},
		{
			name:          "no filter",
			buildQuery:    tagQuery.NewTagQuery,
			expectedCount: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tags, count, err := repo.FindAll(tt.buildQuery())
			require.NoError(t, err)
			assert.Equal(t, tt.expectedCount, count)
			assert.Len(t, tags, int(tt.expectedCount))
		})
	}
}
