package http_test

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"encoding/json"
	"bytes"
	tagHttp "github.com/jambo0624/blog/internal/tag/interfaces/http"
	tagService "github.com/jambo0624/blog/internal/tag/application/service"
	tagMock "github.com/jambo0624/blog/tests/testutil/mock/tag"
)

func TestCreateTag(t *testing.T) {
	mockRepo := new(tagMock.MockTagRepository)
	mockService := tagService.NewTagService(mockRepo)
	handler := tagHttp.NewTagHandler(mockService)

	body := `{"name": "Test Tag"}`
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	jsonData, _ := json.Marshal(body)
	c.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.CreateTag(c)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetTag(t *testing.T) {
	mockRepo := new(tagMock.MockTagRepository)
	mockService := tagService.NewTagService(mockRepo)
	handler := tagHttp.NewTagHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	handler.GetTag(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTag(t *testing.T) {
	mockRepo := new(tagMock.MockTagRepository)
	mockService := tagService.NewTagService(mockRepo)
	handler := tagHttp.NewTagHandler(mockService)

	body := `{"name": "Updated Tag"}`
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}
	jsonData, _ := json.Marshal(body)
	c.Request = httptest.NewRequest("PUT", "/", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.UpdateTag(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteTag(t *testing.T) {
	mockRepo := new(tagMock.MockTagRepository)
	mockService := tagService.NewTagService(mockRepo)
	handler := tagHttp.NewTagHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	handler.DeleteTag(c)

	assert.Equal(t, http.StatusOK, w.Code)
}
