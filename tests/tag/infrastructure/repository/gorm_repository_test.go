package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/jambo0624/blog/tests/testutil"
	tagRepo "github.com/jambo0624/blog/internal/tag/infrastructure/repository"
	tagQuery "github.com/jambo0624/blog/internal/tag/domain/query"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
)

func TestGormTagRepository_FindByID(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	repo := tagRepo.NewGormTagRepository(testDB.DB)
	tag := testDB.Data.Tags[0]

	found, err := repo.FindByID(tag.ID)
	assert.NoError(t, err)
	assert.Equal(t, tag.Name, found.Name)
	assert.Equal(t, tag.Color, found.Color)
}

func TestGormTagRepository_FindAll(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	repo := tagRepo.NewGormTagRepository(testDB.DB)

	t.Run("with name filter", func(t *testing.T) {
		q := tagQuery.NewTagQuery()
		q.WithNameLike("Go")

		tags, total, err := repo.FindAll(q)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), total)
		assert.Contains(t, tags[0].Name, "Go")
	})
}

func TestGormTagRepository_Save(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	repo := tagRepo.NewGormTagRepository(testDB.DB)

	tag := &tagEntity.Tag{
		Name:  "New Tag",
		Color: "#000000",
	}

	err := repo.Save(tag)
	assert.NoError(t, err)
	assert.NotZero(t, tag.ID)

	found, err := repo.FindByID(tag.ID)
	assert.NoError(t, err)
	assert.Equal(t, tag.Name, found.Name)
}

func TestGormTagRepository_Update(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	repo := tagRepo.NewGormTagRepository(testDB.DB)
	tag := testDB.Data.Tags[0]

	tag.Name = "Updated Name"
	err := repo.Update(tag)
	assert.NoError(t, err)

	found, err := repo.FindByID(tag.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", found.Name)
}

func TestGormTagRepository_Delete(t *testing.T) {
	testDB, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	repo := tagRepo.NewGormTagRepository(testDB.DB)
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