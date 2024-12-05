package http

import (
	"github.com/gin-gonic/gin"
	articleService "github.com/jambo0624/blog/internal/article/application/service"
	sharedHttp "github.com/jambo0624/blog/internal/shared/interfaces/http"
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
)

type ArticleHandler struct {
	service *articleService.ArticleService
}

func NewArticleHandler(as *articleService.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		service: as,
	}
}

func (h *ArticleHandler) Create(c *gin.Context) {
	var req dto.CreateArticleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	article, err := h.service.Create(&req)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, article)
}

func (h *ArticleHandler) FindByID(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	article, err := h.service.FindByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(200, article)
}

func (h *ArticleHandler) FindAll(c *gin.Context) {
	articles, err := h.service.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, articles)
}

func (h *ArticleHandler) Update(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")

	var req dto.UpdateArticleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	article, err := h.service.Update(id, &req)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, article)
}

func (h *ArticleHandler) Delete(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
} 