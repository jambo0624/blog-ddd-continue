package service_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/jambo0624/blog/internal/article/application/service"
	"github.com/jambo0624/blog/internal/article/domain/entity"
	"github.com/jambo0624/blog/internal/article/domain/query"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
	mockArticle "github.com/jambo0624/blog/tests/testutil/mock/article"
	mockCategory "github.com/jambo0624/blog/tests/testutil/mock/category"
	mockTag "github.com/jambo0624/blog/tests/testutil/mock/tag"
)

func TestArticleService_Create(t *testing.T) {
	mockArticleRepo := new(mockArticle.MockArticleRepository)
	mockCategoryRepo := new(mockCategory.MockCategoryRepository)
	mockTagRepo := new(mockTag.MockTagRepository)

	articleService := service.NewArticleService(mockArticleRepo, mockCategoryRepo, mockTagRepo)

	category := &categoryEntity.Category{ID: 1, Name: "Test Category"}
	tag := &tagEntity.Tag{ID: 1, Name: "Test Tag"}

	// Setup expectations
	mockCategoryRepo.On("FindByID", uint(1), mock.Anything).Return(category, nil)
	mockTagRepo.On("FindByID", uint(1), mock.Anything).Return(tag, nil)
	mockArticleRepo.On("Save", mock.AnythingOfType("*entity.Article")).Return(nil)

	req := &dto.CreateArticleRequest{
		CategoryID: 1,
		Title:     "Test Article",
		Content:   "Test Content",
		TagIDs:    []uint{1},
	}

	article, err := articleService.Create(req)

	assert.NoError(t, err)
	assert.NotNil(t, article)
	assert.Equal(t, req.Title, article.Title)
	assert.Equal(t, req.Content, article.Content)
	assert.Equal(t, category.ID, article.CategoryID)
	assert.Len(t, article.Tags, 1)
	assert.Equal(t, tag.ID, article.Tags[0].ID)
}

func TestArticleService_FindAll(t *testing.T) {
	mockArticleRepo := new(mockArticle.MockArticleRepository)
	mockCategoryRepo := new(mockCategory.MockCategoryRepository)
	mockTagRepo := new(mockTag.MockTagRepository)

	articleService := service.NewArticleService(mockArticleRepo, mockCategoryRepo, mockTagRepo)

	articles := []*entity.Article{
		{ID: 1, Title: "Test Article 1"},
		{ID: 2, Title: "Test Article 2"},
	}

	q := query.NewArticleQuery()
	mockArticleRepo.On("FindAll", q).Return(articles, int64(2), nil)

	found, total, err := articleService.FindAll(q)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Len(t, found, 2)
	assert.Equal(t, "Test Article 1", found[0].Title)
}

func TestArticleService_Update(t *testing.T) {
	mockArticleRepo := new(mockArticle.MockArticleRepository)
	mockCategoryRepo := new(mockCategory.MockCategoryRepository)
	mockTagRepo := new(mockTag.MockTagRepository)

	articleService := service.NewArticleService(mockArticleRepo, mockCategoryRepo, mockTagRepo)

	article := &entity.Article{ID: 1, Title: "Old Title"}
	category := &categoryEntity.Category{ID: 2, Name: "New Category"}
	tag := &tagEntity.Tag{ID: 2, Name: "New Tag"}

	mockArticleRepo.On("FindByID", uint(1), mock.Anything).Return(article, nil)
	mockCategoryRepo.On("FindByID", uint(2), mock.Anything).Return(category, nil)
	mockTagRepo.On("FindByID", uint(2), mock.Anything).Return(tag, nil)
	mockArticleRepo.On("Update", mock.AnythingOfType("*entity.Article")).Return(nil)

	req := &dto.UpdateArticleRequest{
		CategoryID: 2,
		Title:     "New Title",
		Content:   "New Content",
		TagIDs:    []uint{2},
	}

	updated, err := articleService.Update(1, req)

	assert.NoError(t, err)
	assert.Equal(t, "New Title", updated.Title)
	assert.Equal(t, "New Content", updated.Content)
	assert.Equal(t, category.ID, updated.CategoryID)
	assert.Len(t, updated.Tags, 1)
	assert.Equal(t, tag.ID, updated.Tags[0].ID)
}

func TestArticleService_Delete(t *testing.T) {
	mockArticleRepo := new(mockArticle.MockArticleRepository)
	mockCategoryRepo := new(mockCategory.MockCategoryRepository)
	mockTagRepo := new(mockTag.MockTagRepository)

	articleService := service.NewArticleService(mockArticleRepo, mockCategoryRepo, mockTagRepo)

	mockArticleRepo.On("Delete", uint(1)).Return(nil)

	err := articleService.Delete(1)

	assert.NoError(t, err)
	mockArticleRepo.AssertExpectations(t)
}