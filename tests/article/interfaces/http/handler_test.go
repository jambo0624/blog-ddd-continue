package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	articleHttp "github.com/jambo0624/blog/internal/article/interfaces/http"
	articleService "github.com/jambo0624/blog/internal/article/application/service"
	articleMock "github.com/jambo0624/blog/tests/testutil/mock/article"
	categoryMock "github.com/jambo0624/blog/tests/testutil/mock/category"
	tagMock "github.com/jambo0624/blog/tests/testutil/mock/tag"
)

func TestArticleHandler_CreateArticle(t *testing.T) {
	mockRepo := new(articleMock.MockArticleRepository)
	mockCategoryRepo := new(categoryMock.MockCategoryRepository)
	mockTagRepo := new(tagMock.MockTagRepository)
	mockService := articleService.NewArticleService(mockRepo, mockCategoryRepo, mockTagRepo)
	handler := articleHttp.NewArticleHandler(mockService)

	body := map[string]interface{}{
		"category_id": 1,
		"title":      "Test Article",
		"content":    "Test Content",
		"tag_ids":    []uint{1, 2},
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	jsonData, _ := json.Marshal(body)
	c.Request = httptest.NewRequest("POST", "/", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.CreateArticle(c)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestArticleHandler_GetArticle(t *testing.T) {
	mockRepo := new(articleMock.MockArticleRepository)
	mockCategoryRepo := new(categoryMock.MockCategoryRepository)
	mockTagRepo := new(tagMock.MockTagRepository)
	mockService := articleService.NewArticleService(mockRepo, mockCategoryRepo, mockTagRepo)
	handler := articleHttp.NewArticleHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	handler.GetArticle(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestArticleHandler_UpdateArticle(t *testing.T) {
	mockRepo := new(articleMock.MockArticleRepository)
	mockCategoryRepo := new(categoryMock.MockCategoryRepository)
	mockTagRepo := new(tagMock.MockTagRepository)
	mockService := articleService.NewArticleService(mockRepo, mockCategoryRepo, mockTagRepo)
	handler := articleHttp.NewArticleHandler(mockService)

	body := map[string]interface{}{
		"category_id": 1,
		"title":      "Updated Article",
		"content":    "Updated Content",
		"tag_ids":    []uint{1, 2},
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	jsonData, _ := json.Marshal(body)
	c.Request = httptest.NewRequest("PUT", "/", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	handler.UpdateArticle(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestArticleHandler_DeleteArticle(t *testing.T) {
	mockRepo := new(articleMock.MockArticleRepository)
	mockCategoryRepo := new(categoryMock.MockCategoryRepository)
	mockTagRepo := new(tagMock.MockTagRepository)
	mockService := articleService.NewArticleService(mockRepo, mockCategoryRepo, mockTagRepo)
	handler := articleHttp.NewArticleHandler(mockService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "id", Value: "1"}}

	handler.DeleteArticle(c)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
