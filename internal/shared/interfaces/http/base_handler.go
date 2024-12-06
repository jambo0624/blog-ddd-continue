package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jambo0624/blog/internal/shared/application/service"
	"github.com/jambo0624/blog/internal/shared/domain/repository"
)

// RequestDTO interface for create/update requests
type RequestDTO interface {
	Validate() error
}

// EntityService interface for create/update operations
type EntityService[T repository.Entity, Q repository.Query, C RequestDTO, U RequestDTO] interface {
	Create(req *C) (*T, error)
	Update(id uint, req *U) (*T, error)
}

type BaseHandler[T repository.Entity, Q repository.Query, C RequestDTO, U RequestDTO] struct {
	Service       *service.BaseService[T, Q]
	EntityService EntityService[T, Q, C, U]
}

func NewBaseHandler[T repository.Entity, Q repository.Query, C RequestDTO, U RequestDTO](
	service *service.BaseService[T, Q],
	entityService EntityService[T, Q, C, U],
) *BaseHandler[T, Q, C, U] {
	return &BaseHandler[T, Q, C, U]{
		Service:       service,
		EntityService: entityService,
	}
}

// Create handles POST / requests
func (h *BaseHandler[T, Q, C, U]) Create(c *gin.Context) {
	var req C
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	entity, err := h.EntityService.Create(&req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, entity)
}

// Update handles PUT /:id requests
func (h *BaseHandler[T, Q, C, U]) Update(c *gin.Context) {
	id := ParseUintParam(c, "id")

	var req U
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	entity, err := h.EntityService.Update(id, &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, entity)
}

// FindByID handles GET /:id requests
func (h *BaseHandler[T, Q, C, U]) FindByID(c *gin.Context) {
	id := ParseUintParam(c, "id")
	entity, err := h.Service.FindByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}
	c.JSON(200, entity)
}

// FindAll handles GET / requests with query parameters
func (h *BaseHandler[T, Q, C, U]) FindAll(c *gin.Context, buildQuery func(*gin.Context) (Q, error)) {
	query, err := buildQuery(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	entities, err := h.Service.FindAll(query)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": entities,
		"meta": gin.H{
			"total": len(entities),
		},
	})
}

// Delete handles DELETE /:id requests
func (h *BaseHandler[T, Q, C, U]) Delete(c *gin.Context) {
	id := ParseUintParam(c, "id")
	if err := h.Service.Delete(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil)
}
