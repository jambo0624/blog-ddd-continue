package factory

import (
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	"github.com/jambo0624/blog/internal/tag/interfaces/http/dto"
)

type TagFactory struct {
	BaseFactory
}

func NewTagFactory() *TagFactory {
	return &TagFactory{
		BaseFactory: NewBaseFactory(),
	}
}

func (f *TagFactory) BuildEntity(opts ...func(*tagEntity.Tag)) *tagEntity.Tag {
	seq := f.NextSequence()
	entity := &tagEntity.Tag{
		ID:    seq,
		Name:  f.FormatTestName("Tag"),
		Color: f.FormatHexColor(),
	}
	return ApplyOptions(entity, opts)
}

func (f *TagFactory) buildRequest(isUpdate bool) interface{} {
	name := f.FormatTestName("Tag")
	if isUpdate {
		name = f.FormatUpdatedName("Tag")
	}
	color := f.FormatHexColor()

	if isUpdate {
		return &dto.UpdateTagRequest{Name: name, Color: color}
	}
	return &dto.CreateTagRequest{Name: name, Color: color}
}

func (f *TagFactory) BuildCreateRequest(opts ...func(*dto.CreateTagRequest)) *dto.CreateTagRequest {
	req := BuildRequest[*dto.CreateTagRequest](false, f.buildRequest)
	return ApplyOptions(req, opts)
}

func (f *TagFactory) BuildUpdateRequest(opts ...func(*dto.UpdateTagRequest)) *dto.UpdateTagRequest {
	req := BuildRequest[*dto.UpdateTagRequest](true, f.buildRequest)
	return ApplyOptions(req, opts)
}

// WithName sets custom name.
func (f *TagFactory) WithName(name string) func(*tagEntity.Tag) {
	return func(t *tagEntity.Tag) {
		t.Name = name
	}
}

// WithColor sets custom color.
func (f *TagFactory) WithColor(color string) func(*tagEntity.Tag) {
	return func(t *tagEntity.Tag) {
		t.Color = color
	}
}

// BuildList creates a list of Tag entities.
func (f *TagFactory) BuildList(count int) []*tagEntity.Tag {
	tags := make([]*tagEntity.Tag, count)
	for i := range tags {
		tags[i] = f.BuildEntity()
	}
	return tags
}
