package service_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/jambo0624/blog/internal/category/application/service"
    categoryMock "github.com/jambo0624/blog/tests/testutil/mock/category"
)

func TestCategoryService_CreateCategory(t *testing.T) {
    mockRepo := new(categoryMock.MockCategoryRepository)
    categoryService := service.NewCategoryService(mockRepo)

    mockRepo.On("Save", mock.AnythingOfType("*entity.Category")).Return(nil)

    category, err := categoryService.CreateCategory("Test Category")

    assert.NoError(t, err)
    assert.NotNil(t, category)
    assert.Equal(t, "Test Category", category.Name)
    mockRepo.AssertExpectations(t)
} 