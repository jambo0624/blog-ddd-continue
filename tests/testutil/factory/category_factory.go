package factory

import (
	"fmt"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	"github.com/jambo0624/blog/internal/category/interfaces/http/dto"
)

type CategoryFactory struct {
	sequence int
}

func NewCategoryFactory() *CategoryFactory {
	return &CategoryFactory{sequence: 2}
}

// BuildEntity creates a Category entity
func (f *CategoryFactory) BuildEntity(opts ...func(*categoryEntity.Category)) *categoryEntity.Category {
	f.sequence++
	category := &categoryEntity.Category{
		ID:   uint(f.sequence),
		Name: fmt.Sprintf("Test Category %d", f.sequence),
		Slug: fmt.Sprintf("test-category-%d", f.sequence),
	}

	for _, opt := range opts {
		opt(category)
	}

	return category
}

// BuildCreateRequest creates a CreateCategoryRequest
func (f *CategoryFactory) BuildCreateRequest(opts ...func(*dto.CreateCategoryRequest)) *dto.CreateCategoryRequest {
	f.sequence++
	req := &dto.CreateCategoryRequest{
		Name: fmt.Sprintf("Test Category %d", f.sequence),
		Slug: fmt.Sprintf("test-category-%d", f.sequence),
	}

	for _, opt := range opts {
		opt(req)
	}

	return req
}

// BuildUpdateRequest creates an UpdateCategoryRequest
func (f *CategoryFactory) BuildUpdateRequest(opts ...func(*dto.UpdateCategoryRequest)) *dto.UpdateCategoryRequest {
	f.sequence++
	req := &dto.UpdateCategoryRequest{
		Name: fmt.Sprintf("Updated Category %d", f.sequence),
		Slug: fmt.Sprintf("updated-category-%d", f.sequence),
	}

	for _, opt := range opts {
		opt(req)
	}

	return req
}

// BuildList creates a list of Category entities
func (f *CategoryFactory) BuildList(count int) []*categoryEntity.Category {
	categories := make([]*categoryEntity.Category, count)
	for i := 0; i < count; i++ {
		categories[i] = f.BuildEntity()
	}
	return categories
}

// WithName sets custom name
func (f *CategoryFactory) WithName(name string) func(*categoryEntity.Category) {
	return func(c *categoryEntity.Category) {
		c.Name = name
	}
}

// WithSlug sets custom slug
func (f *CategoryFactory) WithSlug(slug string) func(*categoryEntity.Category) {
	return func(c *categoryEntity.Category) {
		c.Slug = slug
	}
}
