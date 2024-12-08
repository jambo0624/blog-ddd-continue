package response

import (
	"strings"

	"github.com/jambo0624/blog/internal/shared/domain/query"
)

// Meta standard metadata structure.
type Meta struct {
	Total       int         `json:"total"`                 // Total number of records
	Limit       int         `json:"limit,omitempty"`       // Page size
	Offset      int         `json:"offset,omitempty"`      // Page offset
	Page        int         `json:"page,omitempty"`        // Current page number
	TotalPages  int         `json:"totalPages,omitempty"`  // Total number of pages
	Sort        string      `json:"sort,omitempty"`        // Sort field
	Order       string      `json:"order,omitempty"`       // Sort order (asc/desc)
	Filter      interface{} `json:"filter,omitempty"`      // Applied filters
	Aggregation interface{} `json:"aggregation,omitempty"` // Aggregation results
}

// NewMeta creates a new Meta instance with pagination info.
func NewMeta(total, limit, offset int) *Meta {
	meta := &Meta{
		Total:  total,
		Limit:  limit,
		Offset: offset,
	}

	if limit > 0 {
		meta.Page = (offset / limit) + 1
		meta.TotalPages = (total + limit - 1) / limit
	}

	return meta
}

// WithSort adds sorting information.
func (m *Meta) WithSort(field, order string) *Meta {
	m.Sort = field
	m.Order = order
	return m
}

// WithFilter adds filter information.
func (m *Meta) WithFilter(filter interface{}) *Meta {
	m.Filter = filter
	return m
}

// WithAggregation adds aggregation results.
func (m *Meta) WithAggregation(aggregation interface{}) *Meta {
	m.Aggregation = aggregation
	return m
}

// NewMetaFromQuery creates a new Meta instance from query parameters.
func NewMetaFromQuery(total int64, baseQuery query.BaseQuery) *Meta {
	meta := NewMeta(int(total), baseQuery.Limit, baseQuery.Offset)

	// Add sort info
	if baseQuery.OrderBy != "" {
		field := strings.TrimSuffix(strings.TrimPrefix(baseQuery.OrderBy, "-"), " DESC")
		order := "asc"
		if strings.HasSuffix(baseQuery.OrderBy, " DESC") || strings.HasPrefix(baseQuery.OrderBy, "-") {
			order = "desc"
		}
		meta.WithSort(field, order)
	}

	return meta
}
