package bootstrap

import (
	articleService "github.com/jambo0624/blog/internal/article/application/service"
	categoryService "github.com/jambo0624/blog/internal/category/application/service"
	tagService "github.com/jambo0624/blog/internal/tag/application/service"
)

type Services struct {
	Article  *articleService.ArticleService
	Category *categoryService.CategoryService
	Tag      *tagService.TagService
}

func SetupServices(repos *Repositories) *Services {
	return &Services{
		Article:  articleService.NewArticleService(repos.Article, repos.Category, repos.Tag),
		Category: categoryService.NewCategoryService(repos.Category),
		Tag:      tagService.NewTagService(repos.Tag),
	}
}
