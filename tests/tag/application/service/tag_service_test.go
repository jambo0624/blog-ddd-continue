package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	tagService "github.com/jambo0624/blog/internal/tag/application/service"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	tagQuery "github.com/jambo0624/blog/internal/tag/domain/query"
	"github.com/jambo0624/blog/tests/testutil/factory"
	mockTag "github.com/jambo0624/blog/tests/testutil/mock/tag"
)

func setupTest(t *testing.T) (*tagService.TagService, *mockTag.MockTagRepository, *factory.TagFactory) {
	t.Helper()

	mockRepo := new(mockTag.MockTagRepository)
	service := tagService.NewTagService(mockRepo)
	factory := factory.NewTagFactory()

	return service, mockRepo, factory
}

func TestTagService_Create(t *testing.T) {
	service, mockRepo, factory := setupTest(t)

	// prepare data
	req := factory.BuildCreateRequest()
	expectedTag := factory.BuildEntity(
		factory.WithName(req.Name),
		factory.WithColor(req.Color),
	)

	mockRepo.On("Save", mock.MatchedBy(func(t *tagEntity.Tag) bool {
		return t.Name == expectedTag.Name && t.Color == expectedTag.Color
	})).Return(nil)

	tag, err := service.Create(req)
	require.NoError(t, err)
	assert.Equal(t, expectedTag.Name, tag.Name)
	assert.Equal(t, expectedTag.Color, tag.Color)
}

func TestTagService_Update(t *testing.T) {
	service, mockRepo, factory := setupTest(t)

	existingTag := factory.BuildEntity()
	req := factory.BuildUpdateRequest()

	mockRepo.On("FindByID", existingTag.ID, mock.Anything).Return(existingTag, nil)
	mockRepo.On("Update", mock.MatchedBy(func(t *tagEntity.Tag) bool {
		return t.ID == existingTag.ID && t.Name == req.Name && t.Color == req.Color
	})).Return(nil)

	tag, err := service.Update(existingTag.ID, req)
	require.NoError(t, err)
	assert.Equal(t, req.Name, tag.Name)
	assert.Equal(t, req.Color, tag.Color)
}

func TestTagService_FindByID(t *testing.T) {
	service, mockRepo, factory := setupTest(t)

	expectedTag := factory.BuildEntity()
	mockRepo.On("FindByID", expectedTag.ID, mock.Anything).Return(expectedTag, nil)

	tag, err := service.FindByID(expectedTag.ID)
	require.NoError(t, err)
	assert.Equal(t, expectedTag.Name, tag.Name)
	assert.Equal(t, expectedTag.Color, tag.Color)
}

func TestTagService_FindAll(t *testing.T) {
	service, mockRepo, factory := setupTest(t)

	expectedTags := factory.BuildList(2)
	mockRepo.On("FindAll", mock.AnythingOfType("*query.TagQuery")).
		Return(expectedTags, int64(len(expectedTags)), nil)

	tags, total, err := service.FindAll(tagQuery.NewTagQuery())
	require.NoError(t, err)
	assert.Equal(t, int64(len(expectedTags)), total)
	assert.Len(t, tags, len(expectedTags))
}

func TestTagService_Create_ValidationError(t *testing.T) {
	service, mockRepo, factory := setupTest(t)

	req := factory.BuildCreateRequest()
	req.Name = "" // invalid name

	tag, err := service.Create(req)
	require.Error(t, err)
	assert.Nil(t, tag)
	mockRepo.AssertNotCalled(t, "Save")
}

func TestTagService_Update_NotFound(t *testing.T) {
	service, mockRepo, factory := setupTest(t)

	mockRepo.On("FindByID", uint(999), mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	req := factory.BuildUpdateRequest()
	tag, err := service.Update(999, req)

	require.Error(t, err)
	assert.Nil(t, tag)
	mockRepo.AssertNotCalled(t, "Update")
}
