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
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
	mockArticle "github.com/jambo0624/blog/tests/testutil/mock/article"
	mockCategory "github.com/jambo0624/blog/tests/testutil/mock/category"
	mockTag "github.com/jambo0624/blog/tests/testutil/mock/tag"
	factory "github.com/jambo0624/blog/tests/testutil/factory"
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

	categoryFactory := factory.NewCategoryFactory()
	tagFactory := factory.NewTagFactory()
	articleFactory := factory.NewArticleFactory(categoryFactory, tagFactory)

	category := categoryFactory.BuildEntity()
	tag := tagFactory.BuildEntity()

	req := articleFactory.BuildCreateRequest(func(r *dto.CreateArticleRequest) {
		r.CategoryID = category.ID
		r.TagIDs = []uint{tag.ID}
	})

	mockCategoryRepo.On("FindByID", category.ID, mock.Anything).Return(category, nil)
	mockTagRepo.On("FindByID", tag.ID, mock.Anything).Return(tag, nil)
	mockArticleRepo.On("Save", mock.AnythingOfType("*entity.Article")).Return(nil)

	body, _ := json.Marshal(req)
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/api/articles", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestArticleHandler_GetByID(t *testing.T) {
	r, mockArticleRepo, _, _ := setupTest(t)

	articleFactory := factory.NewArticleFactory(factory.NewCategoryFactory(), factory.NewTagFactory())
	article := articleFactory.BuildEntity()

	mockArticleRepo.On("FindByID", uint(1), mock.Anything).Return(article, nil)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api/articles/1", nil)
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestArticleHandler_List(t *testing.T) {
	r, mockArticleRepo, _, _ := setupTest(t)

	articleFactory := factory.NewArticleFactory(factory.NewCategoryFactory(), factory.NewTagFactory())
	articles := articleFactory.BuildList(2)

	mockArticleRepo.On("FindAll", mock.AnythingOfType("*query.ArticleQuery")).
		Return(articles, int64(2), nil)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api/articles", nil)
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestArticleHandler_Update(t *testing.T) {
	r, mockArticleRepo, mockCategoryRepo, mockTagRepo := setupTest(t)

	categoryFactory := factory.NewCategoryFactory()
	tagFactory := factory.NewTagFactory()
	articleFactory := factory.NewArticleFactory(categoryFactory, tagFactory)

	category := categoryFactory.BuildEntity()
	tag := tagFactory.BuildEntity()
	article := articleFactory.BuildEntity()

	// update request
	req := articleFactory.BuildUpdateRequest(func(r *dto.UpdateArticleRequest) {
		r.CategoryID = category.ID
		r.TagIDs = []uint{tag.ID}
	})

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