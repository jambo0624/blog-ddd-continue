package factory

import (
	"fmt"
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
)

type ArticleFactory struct {
	sequence        int
	categoryFactory *CategoryFactory
	tagFactory      *TagFactory
}

func NewArticleFactory(categoryFactory *CategoryFactory, tagFactory *TagFactory) *ArticleFactory {
	return &ArticleFactory{
		sequence:        2,
		categoryFactory: categoryFactory,
		tagFactory:      tagFactory,
	}
}

// BuildEntity creates an Article entity with default or custom values
func (f *ArticleFactory) BuildEntity(opts ...func(*articleEntity.Article)) *articleEntity.Article {
	f.sequence++
	category := f.categoryFactory.BuildEntity()
	tagPtrs := f.tagFactory.BuildList(2)
	
	// Convert []*Tag to []Tag
	tags := make([]tagEntity.Tag, len(tagPtrs))
	for i, t := range tagPtrs {
		tags[i] = *t
	}

	article := &articleEntity.Article{
		CategoryID: category.ID,
		Title:      fmt.Sprintf("Test Article %d", f.sequence),
		Content:    fmt.Sprintf("Test Content %d", f.sequence),
		Tags:       tags,
	}

	for _, opt := range opts {
		opt(article)
	}

	return article
}

// BuildCreateRequest creates a CreateArticleRequest
func (f *ArticleFactory) BuildCreateRequest(opts ...func(*dto.CreateArticleRequest)) *dto.CreateArticleRequest {
	f.sequence++
	category := f.categoryFactory.BuildEntity()
	tags := f.tagFactory.BuildList(2)

	tagIDs := make([]uint, len(tags))
	for i, tag := range tags {
		tagIDs[i] = tag.ID
	}

	req := &dto.CreateArticleRequest{
		CategoryID: category.ID,
		Title:      fmt.Sprintf("Test Article %d", f.sequence),
		Content:    fmt.Sprintf("Test Content %d", f.sequence),
		TagIDs:     tagIDs,
	}

	for _, opt := range opts {
		opt(req)
	}

	return req
}

// BuildUpdateRequest creates an UpdateArticleRequest
func (f *ArticleFactory) BuildUpdateRequest(opts ...func(*dto.UpdateArticleRequest)) *dto.UpdateArticleRequest {
	f.sequence++
	category := f.categoryFactory.BuildEntity()
	tags := f.tagFactory.BuildList(2)

	tagIDs := make([]uint, len(tags))
	for i, tag := range tags {
		tagIDs[i] = tag.ID
	}

	req := &dto.UpdateArticleRequest{
		Title:      fmt.Sprintf("Test Article %d", f.sequence),
		Content:    fmt.Sprintf("Test Content %d", f.sequence),
		CategoryID: category.ID,
		TagIDs:     tagIDs,
	}

	for _, opt := range opts {
		opt(req)
	}

	return req
}

// BuildList creates a list of Article entities
func (f *ArticleFactory) BuildList(count int) []*articleEntity.Article {
	articles := make([]*articleEntity.Article, count)
	for i := 0; i < count; i++ {
		articles[i] = f.BuildEntity()
	}
	return articles
}

// WithCategoryID sets custom category id
func (f *ArticleFactory) WithCategoryID(categoryID uint) func(*articleEntity.Article) {
	return func(a *articleEntity.Article) {
		a.CategoryID = categoryID
	}
}

// WithTags sets custom tags
func (f *ArticleFactory) WithTags(tags []tagEntity.Tag) func(*articleEntity.Article) {
	return func(a *articleEntity.Article) {
		a.Tags = tags
	}
}

// WithTitle sets custom title
func (f *ArticleFactory) WithTitle(title string) func(*articleEntity.Article) {
	return func(a *articleEntity.Article) {
		a.Title = title
	}
}

// WithContent sets custom content
func (f *ArticleFactory) WithContent(content string) func(*articleEntity.Article) {
	return func(a *articleEntity.Article) {
		a.Content = content
	}
}
