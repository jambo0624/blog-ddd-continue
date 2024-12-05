package category

import (
	"github.com/stretchr/testify/mock"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
)

type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Save(category *categoryEntity.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) FindByID(id uint) (*categoryEntity.Category, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*categoryEntity.Category), args.Error(1)
}

func (m *MockCategoryRepository) FindBySlug(slug string) (*categoryEntity.Category, error) {
	args := m.Called(slug)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*categoryEntity.Category), args.Error(1)
}

func (m *MockCategoryRepository) Update(category *categoryEntity.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}