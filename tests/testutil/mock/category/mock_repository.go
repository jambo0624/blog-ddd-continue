package category

import (
	"github.com/stretchr/testify/mock"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	categoryQuery "github.com/jambo0624/blog/internal/category/domain/query"
)

type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) Save(category *categoryEntity.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) FindByID(id uint, preloads ...string) (*categoryEntity.Category, error) {
	args := m.Called(id, preloads)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*categoryEntity.Category), args.Error(1)
}

func (m *MockCategoryRepository) FindAll(query *categoryQuery.CategoryQuery) ([]*categoryEntity.Category, int64, error) {
	args := m.Called(query)
	return args.Get(0).([]*categoryEntity.Category), args.Get(1).(int64), args.Error(2)
}

func (m *MockCategoryRepository) Update(category *categoryEntity.Category) error {
	args := m.Called(category)
	return args.Error(0)
}

func (m *MockCategoryRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
} 