package service

import (
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	categoryRepository "github.com/jambo0624/blog/internal/category/domain/repository"
	"github.com/jambo0624/blog/internal/category/interfaces/http/dto"
)

type CategoryService struct {
	categoryRepo categoryRepository.CategoryRepository
}

func NewCategoryService(cr categoryRepository.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: cr,
	}
}

func (s *CategoryService) Create(req *dto.CreateCategoryRequest) (*categoryEntity.Category, error) {
	category, err := categoryEntity.NewCategory(req.Name, req.Slug)
	if err != nil {
		return nil, err
	}

	if err := s.categoryRepo.Save(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) FindByID(id uint) (*categoryEntity.Category, error) {
	return s.categoryRepo.FindByID(id)
}

func (s *CategoryService) FindAll() ([]*categoryEntity.Category, error) {
	return s.categoryRepo.FindAll()
}

func (s *CategoryService) Update(id uint, req *dto.UpdateCategoryRequest) (*categoryEntity.Category, error) {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	category.Update(req)

	if err := s.categoryRepo.Update(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) Delete(id uint) error {
	return s.categoryRepo.Delete(id)
}
