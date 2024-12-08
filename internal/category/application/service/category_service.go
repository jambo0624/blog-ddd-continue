package service

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	"github.com/jambo0624/blog/internal/shared/application/service"
	"github.com/jambo0624/blog/internal/category/domain/entity"
	"github.com/jambo0624/blog/internal/category/domain/query"
	"github.com/jambo0624/blog/internal/category/interfaces/http/dto"
	"github.com/jambo0624/blog/internal/shared/domain/repository"
)

type CategoryService struct {
	*service.BaseService[entity.Category, *query.CategoryQuery]
}

func NewCategoryService(repo repository.BaseRepository[entity.Category, *query.CategoryQuery]) *CategoryService {
	baseService := service.NewBaseService(repo)
	return &CategoryService{
		BaseService: baseService,
	}
}

func (s *CategoryService) Create(req *dto.CreateCategoryRequest) (*entity.Category, error) {
	category, err := entity.NewCategory(req.Name, req.Slug)
	if err != nil {
		sentry.CaptureException(err)
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	if err := s.Repo.Save(category); err != nil {
		sentry.CaptureException(err)
		return nil, fmt.Errorf("failed to save category: %w", err)
	}

	return category, nil
}

func (s *CategoryService) Update(id uint, req *dto.UpdateCategoryRequest) (*entity.Category, error) {
	category, err := s.FindByID(id)
	if err != nil {
		sentry.CaptureException(err)
		return nil, fmt.Errorf("failed to find category by id: %w", err)
	}

	category.Update(req)

	if err := s.Repo.Update(category); err != nil {
		sentry.CaptureException(err)
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	return category, nil
}
