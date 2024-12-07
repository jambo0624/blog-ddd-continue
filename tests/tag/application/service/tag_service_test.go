package service_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jambo0624/blog/internal/tag/application/service"
	"github.com/jambo0624/blog/internal/tag/domain/entity"
	"github.com/jambo0624/blog/internal/tag/domain/query"
	"github.com/jambo0624/blog/internal/tag/interfaces/http/dto"
	mockTag "github.com/jambo0624/blog/tests/testutil/mock/tag"
)

func TestTagService_Create(t *testing.T) {
	mockRepo := new(mockTag.MockTagRepository)
	tagService := service.NewTagService(mockRepo)

	req := &dto.CreateTagRequest{
		Name:  "Test Tag",
		Color: "#FF0000",
	}

	mockRepo.On("Save", mock.AnythingOfType("*entity.Tag")).Return(nil)

	tag, err := tagService.Create(req)

	assert.NoError(t, err)
	assert.NotNil(t, tag)
	assert.Equal(t, req.Name, tag.Name)
	assert.Equal(t, req.Color, tag.Color)
}

func TestTagService_FindAll(t *testing.T) {
	mockRepo := new(mockTag.MockTagRepository)
	tagService := service.NewTagService(mockRepo)

	tags := []*entity.Tag{
		{ID: 1, Name: "Tag 1", Color: "#FF0000"},
		{ID: 2, Name: "Tag 2", Color: "#00FF00"},
	}

	q := query.NewTagQuery()
	mockRepo.On("FindAll", q).Return(tags, int64(2), nil)

	found, total, err := tagService.FindAll(q)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, found, 2)
	assert.Equal(t, "Tag 1", found[0].Name)
}

func TestTagService_Update(t *testing.T) {
	mockRepo := new(mockTag.MockTagRepository)
	tagService := service.NewTagService(mockRepo)

	tag := &entity.Tag{ID: 1, Name: "Old Name", Color: "#000000"}

	mockRepo.On("FindByID", uint(1), mock.Anything).Return(tag, nil)
	mockRepo.On("Update", mock.AnythingOfType("*entity.Tag")).Return(nil)

	req := &dto.UpdateTagRequest{
		Name:  "New Name",
		Color: "#FFFFFF",
	}

	updated, err := tagService.Update(1, req)

	assert.NoError(t, err)
	assert.Equal(t, "New Name", updated.Name)
	assert.Equal(t, "#FFFFFF", updated.Color)
}

func TestTagService_Delete(t *testing.T) {
	mockRepo := new(mockTag.MockTagRepository)
	tagService := service.NewTagService(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := tagService.Delete(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
} 