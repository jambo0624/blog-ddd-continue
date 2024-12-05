package article

import (
	"github.com/stretchr/testify/mock"
	"github.com/jambo0624/blog/internal/article/domain/entity"
)

type MockArticleRepository struct {
	mock.Mock
}

func (m *MockArticleRepository) Save(article *entity.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func (m *MockArticleRepository) FindByID(id uint) (*entity.Article, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Article), args.Error(1)
}

func (m *MockArticleRepository) Update(article *entity.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func (m *MockArticleRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockArticleRepository) FindAll() ([]*entity.Article, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Article), args.Error(1)
}