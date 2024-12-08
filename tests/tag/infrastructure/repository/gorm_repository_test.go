package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/jambo0624/blog/tests/testutil"
	tagQuery "github.com/jambo0624/blog/internal/tag/domain/query"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	factory "github.com/jambo0624/blog/tests/testutil/factory"
	tagRepository "github.com/jambo0624/blog/internal/tag/domain/repository"
	tagPersistence "github.com/jambo0624/blog/internal/tag/infrastructure/repository"
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
	assert.NoError(t, err)
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
		assert.NoError(t, err)
		assert.Equal(t, int64(1), total)
		assert.Contains(t, tags[0].Name, name)
	})
}

func TestGormTagRepository_Save(t *testing.T) {
	_, cleanup, repo, factory := setupTest(t)
	defer cleanup()

	tag := factory.BuildEntity()

	err := repo.Save(tag)
	assert.NoError(t, err)
	assert.NotZero(t, tag.ID)

	found, err := repo.FindByID(tag.ID)
	assert.NoError(t, err)
	assert.Equal(t, tag.Name, found.Name)
	assert.Equal(t, tag.Color, found.Color)
}

func TestGormTagRepository_Update(t *testing.T) {
	testDB, cleanup, repo, _ := setupTest(t)
	defer cleanup()

	tag := testDB.Data.Tags[0]
	tag.Name = "Updated Name"

	err := repo.Update(tag)
	assert.NoError(t, err)

	found, err := repo.FindByID(tag.ID)
	assert.NoError(t, err)
	assert.Equal(t, tag.Name, found.Name)
}

func TestGormTagRepository_Delete(t *testing.T) {
	testDB, cleanup, repo, _ := setupTest(t)
	defer cleanup()

	tag := testDB.Data.Tags[0]

	err := repo.Delete(tag.ID)
	assert.NoError(t, err)

	var found tagEntity.Tag
	err = testDB.DB.Unscoped().First(&found, tag.ID).Error
	assert.NoError(t, err)
	assert.NotNil(t, found.DeletedAt)

	_, err = repo.FindByID(tag.ID)
	assert.Nil(t, err)
}