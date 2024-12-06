package service

import (
	"fmt"

	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	articleRepository "github.com/jambo0624/blog/internal/article/domain/repository"
	categoryRepository "github.com/jambo0624/blog/internal/category/domain/repository"
	tagRepository "github.com/jambo0624/blog/internal/tag/domain/repository"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	"github.com/jambo0624/blog/internal/article/domain/query"
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
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

func (s *ArticleService) Create(req *dto.CreateArticleRequest) (*articleEntity.Article, error) {
	// 1. Get category
	category, err := s.categoryRepo.FindByID(req.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}

	// 2. Get tags
	var tags []tagEntity.Tag
	for _, tagID := range req.TagIDs {
		tag, err := s.tagRepo.FindByID(tagID)
		if err != nil {
			return nil, fmt.Errorf("tag not found: %w", err)
		}
		tags = append(tags, *tag)
	}

	// 3. Create article using domain logic
	article, err := articleEntity.NewArticle(category, req.Title, req.Content, tags)
	if err != nil {
		return nil, err
	}

	// 4. Save to repository
	if err := s.articleRepo.Save(article); err != nil {
		return nil, err
	}

	return article, nil
}

func (s *ArticleService) FindByID(id uint) (*articleEntity.Article, error) {
	return s.articleRepo.FindByID(id)
}

func (s *ArticleService) FindAll(q *query.ArticleQuery) ([]*articleEntity.Article, error) {
	if q == nil {
		q = query.NewArticleQuery()
	}
	return s.articleRepo.FindAll(q)
}

func (s *ArticleService) Update(id uint, req *dto.UpdateArticleRequest) (*articleEntity.Article, error) {
	// 1. Get article
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// 2. Get category
	category, err := s.categoryRepo.FindByID(req.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}

	// 3. Get tags
	var tags []tagEntity.Tag
	for _, tagID := range req.TagIDs {
		tag, err := s.tagRepo.FindByID(tagID)
		if err != nil {
			return nil, fmt.Errorf("tag not found: %w", err)
		}
		tags = append(tags, *tag)
	}

	// 4. Update using domain logic
	article.Update(req, category, tags)

	// 5. Save to repository
	if err := s.articleRepo.Update(article); err != nil {
		return nil, err
	}

	return article, nil
}

func (s *ArticleService) Delete(id uint) error {
	return s.articleRepo.Delete(id)
} 