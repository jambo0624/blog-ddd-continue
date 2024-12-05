package http

import (
	"github.com/gin-gonic/gin"
	tagService "github.com/jambo0624/blog/internal/tag/application/service"
	sharedHttp "github.com/jambo0624/blog/internal/shared/interfaces/http"
)

type TagHandler struct {
	tagService *tagService.TagService
}

func NewTagHandler(ts *tagService.TagService) *TagHandler {
	return &TagHandler{
		tagService: ts,
	}
}

func (h *TagHandler) CreateTag(c *gin.Context) {
	var req struct {
		Name  string `json:"name" binding:"required"`
		Color string `json:"color" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tag, err := h.tagService.CreateTag(req.Name, req.Color)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, tag)
}

func (h *TagHandler) GetTag(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	tag, err := h.tagService.GetTagByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Tag not found"})
		return
	}
	c.JSON(200, tag)
}

func (h *TagHandler) UpdateTag(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	var req struct {
		Name  string `json:"name" binding:"required"`
		Color string `json:"color" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tag, err := h.tagService.UpdateTag(id, req.Name, req.Color)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, tag)
}

func (h *TagHandler) DeleteTag(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	if err := h.tagService.DeleteTag(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
} 