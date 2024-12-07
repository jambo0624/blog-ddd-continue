package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	articleHandler "github.com/jambo0624/blog/internal/article/interfaces/http"
	articleService "github.com/jambo0624/blog/internal/article/application/service"
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	mockArticle "github.com/jambo0624/blog/tests/testutil/mock/article"
	mockCategory "github.com/jambo0624/blog/tests/testutil/mock/category"
	mockTag "github.com/jambo0624/blog/tests/testutil/mock/tag"
)

func setupTest(t *testing.T) (*gin.Engine, *mockArticle.MockArticleRepository, *mockCategory.MockCategoryRepository, *mockTag.MockTagRepository) {
	t.Helper()

	gin.SetMode(gin.TestMode)
	r := gin.New()

	mockArticleRepo := new(mockArticle.MockArticleRepository)
	mockCategoryRepo := new(mockCategory.MockCategoryRepository)
	mockTagRepo := new(mockTag.MockTagRepository)

	service := articleService.NewArticleService(mockArticleRepo, mockCategoryRepo, mockTagRepo)
	handler := articleHandler.NewArticleHandler(service)
	router := articleHandler.NewArticleRouter(handler)
	router.Register(r.Group("/api"))

	return r, mockArticleRepo, mockCategoryRepo, mockTagRepo
}

func TestArticleHandler_Create(t *testing.T) {
	r, mockArticleRepo, mockCategoryRepo, mockTagRepo := setupTest(t)

	// prepare test data
	category := &categoryEntity.Category{ID: 1, Name: "Test Category"}
	tag := &tagEntity.Tag{ID: 1, Name: "Test Tag"}

	req := dto.CreateArticleRequest{
		CategoryID: category.ID,
		Title:     "Test Article",
		Content:   "Test Content",
		TagIDs:    []uint{tag.ID},
	}

	// 
	mockCategoryRepo.On("FindByID", category.ID, mock.Anything).Return(category, nil)
	mockTagRepo.On("FindByID", tag.ID, mock.Anything).Return(tag, nil)
	mockArticleRepo.On("Save", mock.AnythingOfType("*entity.Article")).Return(nil)

	body, _ := json.Marshal(req)
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/api/articles", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, request)

	if w.Code != http.StatusCreated {
		t.Logf("Response body: %s", w.Body.String())
	}

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestArticleHandler_GetByID(t *testing.T) {
	r, mockArticleRepo, _, _ := setupTest(t)

	article := &articleEntity.Article{
		ID:      1,
		Title:   "Test Article",
		Content: "Test Content",
	}

	mockArticleRepo.On("FindByID", uint(1), mock.Anything).Return(article, nil)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api/articles/1", nil)
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestArticleHandler_List(t *testing.T) {
	r, mockArticleRepo, _, _ := setupTest(t)

	articles := []*articleEntity.Article{
		{ID: 1, Title: "Article 1"},
		{ID: 2, Title: "Article 2"},
	}

	mockArticleRepo.On("FindAll", mock.AnythingOfType("*query.ArticleQuery")).
		Return(articles, int64(2), nil)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api/articles", nil)
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestArticleHandler_Update(t *testing.T) {
	r, mockArticleRepo, mockCategoryRepo, mockTagRepo := setupTest(t)

	// prepare test data
	category := &categoryEntity.Category{ID: 1, Name: "Test Category"}
	tag := &tagEntity.Tag{ID: 1, Name: "Test Tag"}
	article := &articleEntity.Article{
		ID:         1,
		CategoryID: category.ID,
		Title:      "Test Article",
		Content:    "Test Content",
		Category:   *category,
		Tags:       []tagEntity.Tag{*tag},
	}

	// update request
	req := dto.UpdateArticleRequest{
		CategoryID: category.ID,
		Title:     "Updated Title",
		Content:   "Updated Content",
		TagIDs:    []uint{tag.ID},
	}

	// set mock expectations
	mockArticleRepo.On("FindByID", uint(1), mock.Anything).Return(article, nil)
	mockCategoryRepo.On("FindByID", category.ID, mock.Anything).Return(category, nil)
	mockTagRepo.On("FindByID", tag.ID, mock.Anything).Return(tag, nil)
	mockArticleRepo.On("Update", mock.AnythingOfType("*entity.Article")).Return(nil)

	body, _ := json.Marshal(req)
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPut, "/api/articles/1", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, request)

	if w.Code != http.StatusOK {
		t.Logf("Response body: %s", w.Body.String())
	}

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestArticleHandler_Delete(t *testing.T) {
	r, mockArticleRepo, _, _ := setupTest(t)

	mockArticleRepo.On("Delete", uint(1)).Return(nil)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodDelete, "/api/articles/1", nil)
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusNoContent, w.Code)
}