package factory

import (
	"fmt"

	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	"github.com/jambo0624/blog/internal/tag/interfaces/http/dto"
)

type TagFactory struct {
	sequence int
}

func NewTagFactory() *TagFactory {
	return &TagFactory{sequence: 2}
}

// BuildEntity creates a Tag entity with default or custom values
func (f *TagFactory) BuildEntity(opts ...func(*tagEntity.Tag)) *tagEntity.Tag {
	f.sequence++
	tag := &tagEntity.Tag{
		ID:   uint(f.sequence),
		Name: fmt.Sprintf("Test Tag %d", f.sequence),
		Color: fmt.Sprintf("#%06X", f.sequence),
	}

	// Apply custom options
	for _, opt := range opts {
		opt(tag)
	}

	return tag
}

// BuildCreateRequest creates a CreateTagRequest with default or custom values
func (f *TagFactory) BuildCreateRequest(opts ...func(*dto.CreateTagRequest)) *dto.CreateTagRequest {
	f.sequence++
	req := &dto.CreateTagRequest{
		Name:  fmt.Sprintf("Test Tag %d", f.sequence),
		Color: fmt.Sprintf("#%06X", f.sequence),
	}

	for _, opt := range opts {
		opt(req)
	}

	return req
}

// WithName sets custom name
func (f *TagFactory) WithName(name string) func(*tagEntity.Tag) {
	return func(t *tagEntity.Tag) {
		t.Name = name
	}
}

// WithColor sets custom color
func (f *TagFactory) WithColor(color string) func(*tagEntity.Tag) {
	return func(t *tagEntity.Tag) {
		t.Color = color
	}
}

// BuildUpdateRequest creates an UpdateTagRequest
func (f *TagFactory) BuildUpdateRequest(opts ...func(*dto.UpdateTagRequest)) *dto.UpdateTagRequest {
	f.sequence++
	req := &dto.UpdateTagRequest{
		Name:  fmt.Sprintf("Updated Tag %d", f.sequence),
		Color: fmt.Sprintf("#%06X", f.sequence),
	}

	for _, opt := range opts {
		opt(req)
	}

	return req
}

// BuildList creates a list of Tag entities
func (f *TagFactory) BuildList(count int) []*tagEntity.Tag {
	tags := make([]*tagEntity.Tag, count)
	for i := 0; i < count; i++ {
		tags[i] = f.BuildEntity()
	}
	return tags
}
