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
	errorIndex := 0
	return args.Error(errorIndex)
}

func (m *MockCategoryRepository) FindByID(id uint, preloads ...string) (*categoryEntity.Category, error) {
	args := m.Called(id, preloads)
	resultsIndex := 0
	errorIndex := 1
	if args.Get(resultsIndex) == nil {
		return nil, args.Error(errorIndex)
	}
	return args.Get(resultsIndex).(*categoryEntity.Category), args.Error(errorIndex)
}

func (m *MockCategoryRepository) FindAll(query *categoryQuery.CategoryQuery) (
	[]*categoryEntity.Category, int64, error,
) {
	args := m.Called(query)
	resultsIndex := 0
	countIndex := 1
	errorIndex := 2
	return args.Get(resultsIndex).([]*categoryEntity.Category), args.Get(countIndex).(int64), args.Error(errorIndex)
}

func (m *MockCategoryRepository) Update(category *categoryEntity.Category) error {
	args := m.Called(category)
	errorIndex := 0
	return args.Error(errorIndex)
}

func (m *MockCategoryRepository) Delete(id uint) error {
	args := m.Called(id)
	errorIndex := 0
	return args.Error(errorIndex)
}
