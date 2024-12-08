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

	categoryHandler "github.com/jambo0624/blog/internal/category/interfaces/http"
	categoryService "github.com/jambo0624/blog/internal/category/application/service"
	"github.com/jambo0624/blog/internal/category/domain/entity"
	mockCategory "github.com/jambo0624/blog/tests/testutil/mock/category"
	factory "github.com/jambo0624/blog/tests/testutil/factory"
)

func setupTest(t *testing.T) (*gin.Engine, *mockCategory.MockCategoryRepository) {
	t.Helper()

	gin.SetMode(gin.TestMode)
	r := gin.New()

	mockRepo := new(mockCategory.MockCategoryRepository)
	service := categoryService.NewCategoryService(mockRepo)
	handler := categoryHandler.NewCategoryHandler(service)
	router := categoryHandler.NewCategoryRouter(handler)
	router.Register(r.Group("/api"))

	return r, mockRepo
}

func TestCategoryHandler_Create(t *testing.T) {
	r, mockRepo := setupTest(t)
	factory := factory.NewCategoryFactory()
	req := factory.BuildCreateRequest()
	expectedCategory := factory.BuildEntity(
		factory.WithName(req.Name),
		factory.WithSlug(req.Slug),
	)

	mockRepo.On("Save", mock.MatchedBy(func(c *entity.Category) bool {
		return c.Name == expectedCategory.Name && c.Slug == expectedCategory.Slug
	})).Return(nil)

	body, _ := json.Marshal(req)
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/api/categories", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCategoryHandler_GetByID(t *testing.T) {
	r, mockRepo := setupTest(t)
	factory := factory.NewCategoryFactory()
	category := factory.BuildEntity()

	mockRepo.On("FindByID", uint(1), mock.Anything).Return(category, nil)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api/categories/1", nil)
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCategoryHandler_List(t *testing.T) {
	r, mockRepo := setupTest(t)
	factory := factory.NewCategoryFactory()
	categories := factory.BuildList(2)

	mockRepo.On("FindAll", mock.AnythingOfType("*query.CategoryQuery")).
		Return(categories, int64(2), nil)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api/categories", nil)
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCategoryHandler_Update(t *testing.T) {
	r, mockRepo := setupTest(t)
	factory := factory.NewCategoryFactory()
	req := factory.BuildUpdateRequest()
	expectedCategory := factory.BuildEntity(
		factory.WithName(req.Name),
		factory.WithSlug(req.Slug),
	)

	mockRepo.On("FindByID", uint(1), mock.Anything).Return(expectedCategory, nil)
	mockRepo.On("Update", mock.MatchedBy(func(c *entity.Category) bool {
		return c.Name == expectedCategory.Name && c.Slug == expectedCategory.Slug
	})).Return(nil)

	body, _ := json.Marshal(req)
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPut, "/api/categories/1", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCategoryHandler_Delete(t *testing.T) {
	r, mockRepo := setupTest(t)

	mockRepo.On("Delete", uint(1)).Return(nil)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodDelete, "/api/categories/1", nil)
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusNoContent, w.Code)
} 