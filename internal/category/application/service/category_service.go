package service

import (
	"strings"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	categoryRepository "github.com/jambo0624/blog/internal/category/domain/repository"
)

type CategoryService struct {
	categoryRepo categoryRepository.CategoryRepository
}

func NewCategoryService(cr categoryRepository.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: cr,
	}
}

func (s *CategoryService) CreateCategory(name string) (*categoryEntity.Category, error) {
	// 生成 slug
	slug := strings.ToLower(strings.ReplaceAll(name, " ", "-"))

	category := &categoryEntity.Category{
		Name: name,
		Slug: slug,
	}

	if err := s.categoryRepo.Save(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) GetCategoryByID(id uint) (*categoryEntity.Category, error) {
	return s.categoryRepo.FindByID(id)
}

func (s *CategoryService) UpdateCategory(id uint, name string) (*categoryEntity.Category, error) {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	category.Name = name
	category.Slug = strings.ToLower(strings.ReplaceAll(name, " ", "-"))

	if err := s.categoryRepo.Update(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) DeleteCategory(id uint) error {
	return s.categoryRepo.Delete(id)
}
