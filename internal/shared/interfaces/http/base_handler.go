package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jambo0624/blog/internal/shared/application/service"
	"github.com/jambo0624/blog/internal/shared/domain/repository"
	"github.com/jambo0624/blog/internal/shared/interfaces/http/dto"
	"github.com/jambo0624/blog/internal/shared/interfaces/http/response"
)

// EntityService interface for create/update operations
type EntityService[T repository.Entity, Q repository.Query, C dto.RequestDTO, U dto.RequestDTO] interface {
	Create(req *C) (*T, error)
	Update(id uint, req *U) (*T, error)
}

type BaseHandler[T repository.Entity, Q repository.Query, C dto.RequestDTO, U dto.RequestDTO] struct {
	Service       *service.BaseService[T, Q]
	EntityService EntityService[T, Q, C, U]
}

func NewBaseHandler[T repository.Entity, Q repository.Query, C dto.RequestDTO, U dto.RequestDTO](
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
		response.BadRequest(c, err)
		return
	}

	entity, err := h.EntityService.Create(&req)
	if err != nil {
		response.InternalError(c, err)
		return
	}

	response.Created(c, entity)
}

// Update handles PUT /:id requests
func (h *BaseHandler[T, Q, C, U]) Update(c *gin.Context) {
	id := ParseUintParam(c, "id")

	var req U
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err)
		return
	}

	entity, err := h.EntityService.Update(id, &req)
	if err != nil {
		response.InternalError(c, err)
		return
	}

	response.Success(c, entity)
}

// FindByID handles GET /:id requests
func (h *BaseHandler[T, Q, C, U]) FindByID(c *gin.Context) {
	id := ParseUintParam(c, "id")
	entity, err := h.Service.FindByID(id)
	if err != nil {
		response.NotFound(c)
		return
	}
	response.Success(c, entity)
}

// FindAll handles GET / requests with query parameters
func (h *BaseHandler[T, Q, C, U]) FindAll(c *gin.Context, buildQuery func(*gin.Context) (Q, error)) {
	query, err := buildQuery(c)
	if err != nil {
		response.BadRequest(c, err)
		return
	}

	entities, total, err := h.Service.FindAll(query)
	if err != nil {
		response.InternalError(c, err)
		return
	}

	meta := response.NewMetaFromQuery(total, query.GetBaseQuery())
	response.SuccessWithMeta(c, entities, *meta)
}

// Delete handles DELETE /:id requests
func (h *BaseHandler[T, Q, C, U]) Delete(c *gin.Context) {
	id := ParseUintParam(c, "id")
	if err := h.Service.Delete(id); err != nil {
		response.InternalError(c, err)
		return
	}
	response.NoContent(c)
}
