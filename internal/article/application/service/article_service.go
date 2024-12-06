package service

import (
	"fmt"
	"github.com/jambo0624/blog/internal/shared/application/service"
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	categoryRepository "github.com/jambo0624/blog/internal/category/domain/repository"
	tagRepository "github.com/jambo0624/blog/internal/tag/domain/repository"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	"github.com/jambo0624/blog/internal/article/domain/query"
	"github.com/jambo0624/blog/internal/article/interfaces/http/dto"
	"github.com/jambo0624/blog/internal/shared/domain/repository"
)

type ArticleService struct {
	*service.BaseService[articleEntity.Article, *query.ArticleQuery]
	categoryRepo categoryRepository.CategoryRepository
	tagRepo      tagRepository.TagRepository
}

func NewArticleService(
	repo repository.BaseRepository[articleEntity.Article, *query.ArticleQuery],
	cr categoryRepository.CategoryRepository,
	tr tagRepository.TagRepository,
) *ArticleService {
	baseService := service.NewBaseService(repo)
	return &ArticleService{
		BaseService:  baseService,
		categoryRepo: cr,
		tagRepo:     tr,
	}
}

func (s *ArticleService) Create(req *dto.CreateArticleRequest) (*articleEntity.Article, error) {
	category, err := s.categoryRepo.FindByID(req.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}

	var tags []tagEntity.Tag
	for _, tagID := range req.TagIDs {
		tag, err := s.tagRepo.FindByID(tagID)
		if err != nil {
			return nil, fmt.Errorf("tag not found: %w", err)
		}
		tags = append(tags, *tag)
	}

	article, err := articleEntity.NewArticle(category, req.Title, req.Content, tags)
	if err != nil {
		return nil, err
	}

	if err := s.Repo.Save(article); err != nil {
		return nil, err
	}

	return article, nil
}

func (s *ArticleService) Update(id uint, req *dto.UpdateArticleRequest) (*articleEntity.Article, error) {
	article, err := s.FindByID(id)
	if err != nil {
		return nil, err
	}

	category, err := s.categoryRepo.FindByID(req.CategoryID)
	if err != nil {
		return nil, fmt.Errorf("category not found: %w", err)
	}

	var tags []tagEntity.Tag
	for _, tagID := range req.TagIDs {
		tag, err := s.tagRepo.FindByID(tagID)
		if err != nil {
			return nil, fmt.Errorf("tag not found: %w", err)
		}
		tags = append(tags, *tag)
	}

	article.Update(req, category, tags)

	if err := s.Repo.Update(article); err != nil {
		return nil, err
	}

	return article, nil
}
