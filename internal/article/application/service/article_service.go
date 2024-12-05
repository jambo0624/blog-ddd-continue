package service

import (
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	articleRepository "github.com/jambo0624/blog/internal/article/domain/repository"
	categoryRepository "github.com/jambo0624/blog/internal/category/domain/repository"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	tagRepository "github.com/jambo0624/blog/internal/tag/domain/repository"
)

type ArticleService struct {
	articleRepo  articleRepository.ArticleRepository
	categoryRepo categoryRepository.CategoryRepository
	tagRepo      tagRepository.TagRepository
}

func NewArticleService(
	ar articleRepository.ArticleRepository,
	cr categoryRepository.CategoryRepository,
	tr tagRepository.TagRepository,
) *ArticleService {
	return &ArticleService{
		articleRepo:  ar,
		categoryRepo: cr,
		tagRepo:     tr,
	}
}

func (s *ArticleService) CreateArticle(categoryID uint, title, content string, tagIDs []uint) (*articleEntity.Article, error) {
	// 检查分类是否存在
	category, err := s.categoryRepo.FindByID(categoryID)
	if err != nil {
		return nil, err
	}

	// 创建文章
	article := articleEntity.NewArticle(category.ID, title, content)

	// 添加标签
	for _, tagID := range tagIDs {
		tag, err := s.tagRepo.FindByID(tagID)
		if err != nil {
			return nil, err
		}
		article.AddTag(*tag)
	}

	// 保存文章
	if err := s.articleRepo.Save(article); err != nil {
		return nil, err
	}

	return article, nil
}

func (s *ArticleService) GetArticleByID(id uint) (*articleEntity.Article, error) {
	return s.articleRepo.FindByID(id)
}

func (s *ArticleService) GetAllArticles() ([]*articleEntity.Article, error) {
	return s.articleRepo.FindAll()
}

func (s *ArticleService) UpdateArticle(id uint, categoryID uint, title, content string, tagIDs []uint) (*articleEntity.Article, error) {
	// 获取现有文章
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// 检查分类是否存在
	category, err := s.categoryRepo.FindByID(categoryID)
	if err != nil {
		return nil, err
	}

	// 更新基本信息
	article.CategoryID = category.ID
	article.Title = title
	article.Content = content

	// 清除现有标签并添加新标签
	article.Tags = []tagEntity.Tag{}
	for _, tagID := range tagIDs {
		tag, err := s.tagRepo.FindByID(tagID)
		if err != nil {
			return nil, err
		}
		article.AddTag(*tag)
	}

	// 保存更新
	if err := s.articleRepo.Update(article); err != nil {
		return nil, err
	}

	return article, nil
}

func (s *ArticleService) DeleteArticle(id uint) error {
	return s.articleRepo.Delete(id)
} 