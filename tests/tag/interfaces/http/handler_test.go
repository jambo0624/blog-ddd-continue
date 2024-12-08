package http_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	tagHandler "github.com/jambo0624/blog/internal/tag/interfaces/http"
	tagService "github.com/jambo0624/blog/internal/tag/application/service"
	"github.com/jambo0624/blog/internal/tag/domain/entity"
	mockTag "github.com/jambo0624/blog/tests/testutil/mock/tag"
	"github.com/jambo0624/blog/tests/testutil/factory"
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
	factory := factory.NewTagFactory()

	req := factory.BuildCreateRequest()
	expectedTag := factory.BuildEntity(
		factory.WithName(req.Name),
		factory.WithColor(req.Color),
	)

	mockRepo.On("Save", mock.MatchedBy(func(t *entity.Tag) bool {
		return t.Name == expectedTag.Name && t.Color == expectedTag.Color
	})).Return(nil)

	body, _ := json.Marshal(req)
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPost, "/api/tags", bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestTagHandler_GetByID(t *testing.T) {
	r, mockRepo := setupTest(t)
	factory := factory.NewTagFactory()

	tag := factory.BuildEntity()
	mockRepo.On("FindByID", tag.ID, mock.Anything).Return(tag, nil)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/tags/%d", tag.ID), nil)
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTagHandler_List(t *testing.T) {
	r, mockRepo := setupTest(t)
	factory := factory.NewTagFactory()

	tags := factory.BuildList(2)
	mockRepo.On("FindAll", mock.AnythingOfType("*query.TagQuery")).
		Return(tags, int64(len(tags)), nil)

	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/api/tags", nil)
	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTagHandler_Update(t *testing.T) {
	r, mockRepo := setupTest(t)
	factory := factory.NewTagFactory()

	existingTag := factory.BuildEntity()
	req := factory.BuildUpdateRequest()

	mockRepo.On("FindByID", existingTag.ID, mock.Anything).Return(existingTag, nil)
	mockRepo.On("Update", mock.MatchedBy(func(t *entity.Tag) bool {
		return t.ID == existingTag.ID && t.Name == req.Name && t.Color == req.Color
	})).Return(nil)

	body, _ := json.Marshal(req)
	w := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/tags/%d", existingTag.ID), bytes.NewBuffer(body))
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