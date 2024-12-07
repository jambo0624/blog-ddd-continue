package http

import (
	"strconv"
	"strings"
	
	"github.com/gin-gonic/gin"
	"github.com/jambo0624/blog/internal/shared/domain/constants"
	"github.com/jambo0624/blog/internal/shared/domain/validate"
)

// BaseQueryBuilder handles common query parameters
type BaseQueryBuilder struct {
	CommonFields map[string]bool
}

// NewBaseQueryBuilder creates a base query builder
func NewBaseQueryBuilder() *BaseQueryBuilder {
	return &BaseQueryBuilder{
		CommonFields: map[string]bool{
			"id":         true,
			"created_at": true,
			"updated_at": true,
		},
	}
}

// BuildIDs builds ID IN query
func (b *BaseQueryBuilder) BuildIDs(c *gin.Context) ([]uint, error) {
	if ids := c.QueryArray("ids"); len(ids) > 0 {
		uintIDs := make([]uint, 0, len(ids))
		for _, id := range ids {
			uid, err := strconv.ParseUint(id, 10, 32)
			if err != nil {
				return nil, validate.ErrInvalidIDFormat
			}
			uintIDs = append(uintIDs, uint(uid))
		}
		return uintIDs, nil
	}
	return nil, nil
}

// BuildPagination builds pagination parameters
func (b *BaseQueryBuilder) BuildPagination(c *gin.Context, currentLimit, currentOffset int) (int, int, error) {
	limit := currentLimit
	offset := currentOffset

	if limitStr := c.Query("limit"); limitStr != "" {
		l, err := strconv.Atoi(limitStr)
		if err != nil || l < 0 {
			return 0, 0, validate.ErrInvalidLimit
		}
		if l > constants.MaxPageSize {
			l = constants.MaxPageSize
		}
		limit = l
	}

	if offsetStr := c.Query("offset"); offsetStr != "" {
		o, err := strconv.Atoi(offsetStr)
		if err != nil || o < 0 {
			return 0, 0, validate.ErrInvalidOffset
		}
		offset = o
	}

	return limit, offset, nil
}

// BuildOrderBy builds order by parameters
func (b *BaseQueryBuilder) BuildOrderBy(c *gin.Context, additionalFields map[string]bool) (string, error) {
	if orderBy := c.Query("order_by"); orderBy != "" {
		// Merge common fields and additional fields
		allowedFields := make(map[string]bool)
		for k, v := range b.CommonFields {
			allowedFields[k] = v
		}
		for k, v := range additionalFields {
			allowedFields[k] = v
		}

		field := strings.TrimSuffix(strings.TrimPrefix(orderBy, "-"), " DESC")
		if !allowedFields[field] {
			return "", validate.ErrInvalidOrderByField
		}

		return orderBy, nil
	}
	return "", nil
} 