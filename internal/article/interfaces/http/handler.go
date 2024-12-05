package http

import (
	"github.com/gin-gonic/gin"
	articleService "github.com/jambo0624/blog/internal/article/application/service"
	sharedHttp "github.com/jambo0624/blog/internal/shared/interfaces/http"
)

type ArticleHandler struct {
	articleService *articleService.ArticleService
}

func NewArticleHandler(as *articleService.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		articleService: as,
	}
}

func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var req struct {
		CategoryID uint     `json:"category_id" binding:"required"`
		Title     string   `json:"title" binding:"required"`
		Content   string   `json:"content" binding:"required"`
		TagIDs    []uint   `json:"tag_ids"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	article, err := h.articleService.CreateArticle(
		req.CategoryID,
		req.Title,
		req.Content,
		req.TagIDs,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, article)
}

func (h *ArticleHandler) GetArticle(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	article, err := h.articleService.GetArticleByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Article not found"})
		return
	}
	c.JSON(200, article)
}

func (h *ArticleHandler) GetAllArticles(c *gin.Context) {
	articles, err := h.articleService.GetAllArticles()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, articles)
}

func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	
	var req struct {
		CategoryID uint     `json:"category_id" binding:"required"`
		Title     string   `json:"title" binding:"required"`
		Content   string   `json:"content" binding:"required"`
		TagIDs    []uint   `json:"tag_ids"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	article, err := h.articleService.UpdateArticle(
		id,
		req.CategoryID,
		req.Title,
		req.Content,
		req.TagIDs,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, article)
}

func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	if err := h.articleService.DeleteArticle(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
} 