package factory

import (
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	"github.com/jambo0624/blog/internal/category/interfaces/http/dto"
)

type CategoryFactory struct {
	BaseFactory
}

func NewCategoryFactory() *CategoryFactory {
	return &CategoryFactory{
		BaseFactory: NewBaseFactory(),
	}
}

// BuildEntity creates a Category entity.
func (f *CategoryFactory) BuildEntity(opts ...func(*categoryEntity.Category)) *categoryEntity.Category {
	seq := f.NextSequence()
	entity := &categoryEntity.Category{
		ID:   seq,
		Name: f.FormatTestName("Category"),
		Slug: f.FormatTestSlug("category"),
	}
	return ApplyOptions(entity, opts)
}

func (f *CategoryFactory) buildRequest(isUpdate bool) interface{} {
	name := f.FormatTestName("Category")
	slug := f.FormatTestSlug("category")
	if isUpdate {
		name = f.FormatUpdatedName("Category")
		slug = f.FormatUpdatedSlug("category")
	}

	if isUpdate {
		return &dto.UpdateCategoryRequest{Name: name, Slug: slug}
	}
	return &dto.CreateCategoryRequest{Name: name, Slug: slug}
}

// BuildCreateRequest creates a CreateCategoryRequest.
func (f *CategoryFactory) BuildCreateRequest(opts ...func(*dto.CreateCategoryRequest)) *dto.CreateCategoryRequest {
	req := BuildRequest[*dto.CreateCategoryRequest](false, f.buildRequest)
	return ApplyOptions(req, opts)
}

// BuildUpdateRequest creates an UpdateCategoryRequest.
func (f *CategoryFactory) BuildUpdateRequest(opts ...func(*dto.UpdateCategoryRequest)) *dto.UpdateCategoryRequest {
	req := BuildRequest[*dto.UpdateCategoryRequest](true, f.buildRequest)
	return ApplyOptions(req, opts)
}

// BuildList creates a list of Category entities.
func (f *CategoryFactory) BuildList(count int) []*categoryEntity.Category {
	categories := make([]*categoryEntity.Category, count)
	for i := range categories {
		categories[i] = f.BuildEntity()
	}
	return categories
}

// WithName sets custom name.
func (f *CategoryFactory) WithName(name string) func(*categoryEntity.Category) {
	return func(c *categoryEntity.Category) {
		c.Name = name
	}
}

// WithSlug sets custom slug.
func (f *CategoryFactory) WithSlug(slug string) func(*categoryEntity.Category) {
	return func(c *categoryEntity.Category) {
		c.Slug = slug
	}
}
