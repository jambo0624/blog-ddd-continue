package service_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/jambo0624/blog/internal/article/application/service"
    "github.com/jambo0624/blog/internal/article/domain/entity"
    articleMock "github.com/jambo0624/blog/tests/testutil/mock/article"
    categoryMock "github.com/jambo0624/blog/tests/testutil/mock/category"
    tagMock "github.com/jambo0624/blog/tests/testutil/mock/tag"
		categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
)

func TestArticleService_CreateArticle(t *testing.T) {
    mockArticleRepo := new(articleMock.MockArticleRepository)
    mockCategoryRepo := new(categoryMock.MockCategoryRepository)
    mockTagRepo := new(tagMock.MockTagRepository)

    articleService := service.NewArticleService(mockArticleRepo, mockCategoryRepo, mockTagRepo)

    // 设置预期行为
    mockCategoryRepo.On("FindByID", uint(1)).Return(&categoryEntity.Category{ID: 1}, nil)
    mockArticleRepo.On("Save", mock.AnythingOfType("*entity.Article")).Return(nil)

    article, err := articleService.CreateArticle(1, "Test Title", "Test Content", []uint{})

    assert.NoError(t, err)
    assert.NotNil(t, article)
    assert.Equal(t, "Test Title", article.Title)
    mockArticleRepo.AssertExpectations(t)
    mockCategoryRepo.AssertExpectations(t)
}

func TestArticleService_GetArticleByID(t *testing.T) {
    mockArticleRepo := new(articleMock.MockArticleRepository)
    mockCategoryRepo := new(categoryMock.MockCategoryRepository)
    mockTagRepo := new(tagMock.MockTagRepository)

    articleService := service.NewArticleService(mockArticleRepo, mockCategoryRepo, mockTagRepo)

    expectedArticle := &entity.Article{ID: 1, Title: "Test"}
    mockArticleRepo.On("FindByID", uint(1)).Return(expectedArticle, nil)

    article, err := articleService.GetArticleByID(1)

    assert.NoError(t, err)
    assert.Equal(t, expectedArticle, article)
    mockArticleRepo.AssertExpectations(t)
} 