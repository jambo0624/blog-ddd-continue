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

	tagHandler "github.com/jambo0624/blog/internal/tag/interfaces/http"
	tagService "github.com/jambo0624/blog/internal/tag/application/service"
	"github.com/jambo0624/blog/internal/tag/interfaces/http/dto"
	"github.com/jambo0624/blog/internal/tag/domain/entity"
	mockTag "github.com/jambo0624/blog/tests/testutil/mock/tag"
)

func setupTest(t *testing.T) (*gin.Engine, *mockTag.MockTagRepository) {
	t.Helper()

	gin.SetMode(gin.TestMode)
	r := gin.New()

	mockRepo := new(mockTag.MockTagRepository)
	service := tagService.NewTagService(mockRepo)
	handler := tagHandler.NewTagHandler(service)
	router := tagHandler.NewTagRouter(handler)
	router.Register(r.Group("/api"))

	return r, mockRepo
}

func TestTagHandler_Create(t *testing.T) {
	r, mockRepo := setupTest(t)

	req := dto.CreateTagRequest{
		Name:  "Test Tag",
		Color: "#FF0000",
	}

	mockRepo.On("Save", mock.AnythingOfType("*entity.Tag")).Return(nil)
	
	body, _ := json.Marshal(req)
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/api/tags", bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, request)

	if w.Code != http.StatusCreated {
		t.Logf("Response body: %s", w.Body.String())
	}

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestTagHandler_GetByID(t *testing.T) {
	r, mockRepo := setupTest(t)

	tag := &entity.Tag{
		ID:    1,
		Name:  "Test Tag",
		Color: "#FF0000",
	}

	mockRepo.On("FindByID", uint(1), mock.Anything).Return(tag, nil)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api/tags/1", nil)
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTagHandler_List(t *testing.T) {
	r, mockRepo := setupTest(t)

	tags := []*entity.Tag{
		{ID: 1, Name: "Tag 1", Color: "#FF0000"},
		{ID: 2, Name: "Tag 2", Color: "#00FF00"},
	}

	mockRepo.On("FindAll", mock.AnythingOfType("*query.TagQuery")).
		Return(tags, int64(2), nil)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api/tags", nil)
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTagHandler_Update(t *testing.T) {
	r, mockRepo := setupTest(t)

	req := dto.UpdateTagRequest{
		Name:  "Updated Tag",
		Color: "#0000FF",
	}

	mockRepo.On("FindByID", uint(1), mock.Anything).Return(&entity.Tag{ID: 1}, nil)
	mockRepo.On("Update", mock.AnythingOfType("*entity.Tag")).Return(nil)

	body, _ := json.Marshal(req)
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPut, "/api/tags/1", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTagHandler_Delete(t *testing.T) {
	r, mockRepo := setupTest(t)

	mockRepo.On("Delete", uint(1)).Return(nil)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodDelete, "/api/tags/1", nil)
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusNoContent, w.Code)
} 