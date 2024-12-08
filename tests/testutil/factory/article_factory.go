package factory

import (
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
)

type ArticleFactory struct {
	BaseFactory
	categoryFactory *CategoryFactory
	tagFactory      *TagFactory
}

func NewArticleFactory(categoryFactory *CategoryFactory, tagFactory *TagFactory) *ArticleFactory {
	return &ArticleFactory{
		BaseFactory:     NewBaseFactory(),
		categoryFactory: categoryFactory,
		tagFactory:      tagFactory,
	}
}

// buildDependencies creates the necessary category and tags.
func (f *ArticleFactory) buildDependencies() (*categoryEntity.Category, []*tagEntity.Tag) {
	category := f.categoryFactory.BuildEntity()
	defaultTagLength := 2
	tags := f.tagFactory.BuildList(defaultTagLength)
	return category, tags
}

// getTagIDs extracts IDs from tags.
func (f *ArticleFactory) getTagIDs(tags []*tagEntity.Tag) []uint {
	tagIDs := make([]uint, len(tags))
	for i, tag := range tags {
		tagIDs[i] = tag.ID
	}
	return tagIDs
}

func (f *ArticleFactory) BuildEntity(opts ...func(*articleEntity.Article)) (
	*articleEntity.Article,
	*categoryEntity.Category,
	*tagEntity.Tag,
) {
	seq := f.NextSequence()
	category, tagPtrs := f.buildDependencies()

	tags := make([]tagEntity.Tag, len(tagPtrs))
	for i, t := range tagPtrs {
		tags[i] = *t
	}

	article := &articleEntity.Article{
		ID:         seq,
		CategoryID: category.ID,
		Title:      f.FormatTestName("Article"),
		Content:    f.FormatTestName("Content"),
		Tags:       tags,
	}

	return ApplyOptions(article, opts), category, &tags[0]
}

func (f *ArticleFactory) buildRequest(isUpdate bool) interface{} {
	category, tags := f.buildDependencies()
	tagIDs := f.getTagIDs(tags)

	title := f.FormatTestName("Article")
	content := f.FormatTestName("Content")
	if isUpdate {
		title = f.FormatUpdatedName("Article")
		content = f.FormatUpdatedName("Content")
	}

	if isUpdate {
		return &dto.UpdateArticleRequest{
			Title:      title,
			Content:    content,
			CategoryID: category.ID,
			TagIDs:     tagIDs,
		}
	}
	return &dto.CreateArticleRequest{
		Title:      title,
		Content:    content,
		CategoryID: category.ID,
		TagIDs:     tagIDs,
	}
}

func (f *ArticleFactory) BuildCreateRequest(opts ...func(*dto.CreateArticleRequest)) (
	*dto.CreateArticleRequest,
	*categoryEntity.Category,
	*tagEntity.Tag,
) {
	category, tags := f.buildDependencies()
	req := BuildRequest[*dto.CreateArticleRequest](false, f.buildRequest)
	return ApplyOptions(req, opts), category, tags[0]
}

func (f *ArticleFactory) BuildUpdateRequest(opts ...func(*dto.UpdateArticleRequest)) (
	*dto.UpdateArticleRequest,
	*categoryEntity.Category,
	*tagEntity.Tag,
) {
	category, tags := f.buildDependencies()
	req := BuildRequest[*dto.UpdateArticleRequest](true, f.buildRequest)
	return ApplyOptions(req, opts), category, tags[0]
}

func (f *ArticleFactory) BuildList(count int) []*articleEntity.Article {
	articles := make([]*articleEntity.Article, count)
	for i := range articles {
		articles[i], _, _ = f.BuildEntity()
	}
	return articles
}

// Helper methods for customization.
func (f *ArticleFactory) WithCategoryID(categoryID uint) func(*articleEntity.Article) {
	return func(a *articleEntity.Article) {
		a.CategoryID = categoryID
	}
}

func (f *ArticleFactory) WithTags(tags []tagEntity.Tag) func(*articleEntity.Article) {
	return func(a *articleEntity.Article) {
		a.Tags = tags
	}
}

func (f *ArticleFactory) WithTitle(title string) func(*articleEntity.Article) {
	return func(a *articleEntity.Article) {
		a.Title = title
	}
}

func (f *ArticleFactory) WithContent(content string) func(*articleEntity.Article) {
	return func(a *articleEntity.Article) {
		a.Content = content
	}
}
