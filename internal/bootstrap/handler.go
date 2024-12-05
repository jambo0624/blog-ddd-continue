package bootstrap

import (
	articleHttp "github.com/jambo0624/blog/internal/article/interfaces/http"
	categoryHttp "github.com/jambo0624/blog/internal/category/interfaces/http"
	tagHttp "github.com/jambo0624/blog/internal/tag/interfaces/http"
)

type Handlers struct {
	Article  *articleHttp.ArticleHandler
	Category *categoryHttp.CategoryHandler
	Tag      *tagHttp.TagHandler
}

func SetupHandlers(services *Services) *Handlers {
	return &Handlers{
		Article:  articleHttp.NewArticleHandler(services.Article),
		Category: categoryHttp.NewCategoryHandler(services.Category),
		Tag:      tagHttp.NewTagHandler(services.Tag),
	}
} 