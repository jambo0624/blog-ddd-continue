package http

import (
	"github.com/gin-gonic/gin"
	tagService "github.com/jambo0624/blog/internal/tag/application/service"
	"github.com/jambo0624/blog/internal/tag/interfaces/http/dto"
	sharedHttp "github.com/jambo0624/blog/internal/shared/interfaces/http"
)

type TagHandler struct {
	service *tagService.TagService
}

func NewTagHandler(ts *tagService.TagService) *TagHandler {
	return &TagHandler{
		service: ts,
	}
}

func (h *TagHandler) Create(c *gin.Context) {
	var req dto.CreateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tag, err := h.service.Create(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, tag)
}

func (h *TagHandler) FindByID(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	tag, err := h.service.FindByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Tag not found"})
		return
	}
	c.JSON(200, tag)
}

func (h *TagHandler) FindAll(c *gin.Context) {
	tags, err := h.service.FindAll()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tags)
}

func (h *TagHandler) Update(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	var req dto.UpdateTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	tag, err := h.service.Update(id, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, tag)
}

func (h *TagHandler) Delete(c *gin.Context) {
	id := sharedHttp.ParseUintParam(c, "id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
} 