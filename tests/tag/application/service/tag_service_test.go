package service_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	tagService "github.com/jambo0624/blog/internal/tag/application/service"
	tagMock "github.com/jambo0624/blog/tests/testutil/mock/tag"
)

func TestCreateTag(t *testing.T) {
	mockRepo := new(tagMock.MockTagRepository)
	mockService := tagService.NewTagService(mockRepo)

	mockRepo.On("Save", mock.AnythingOfType("*entity.Tag")).Return(nil)

	tag, err := mockService.CreateTag("Test Tag", "test")

	assert.NoError(t, err)
	assert.NotNil(t, tag)
	assert.Equal(t, "Test Tag", tag.Name)
	mockRepo.AssertExpectations(t)
}