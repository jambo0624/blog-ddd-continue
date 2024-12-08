package article

import (
	"github.com/stretchr/testify/mock"

	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	articleQuery "github.com/jambo0624/blog/internal/article/domain/query"
)

type MockArticleRepository struct {
	mock.Mock
}

const (
	resultsIndex = 0
	countIndex   = 1
	errorIndex   = 2
)

func (m *MockArticleRepository) Save(article *articleEntity.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func (m *MockArticleRepository) FindByID(id uint, preloads ...string) (*articleEntity.Article, error) {
	args := m.Called(id, preloads)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*articleEntity.Article), args.Error(1)
}

func (m *MockArticleRepository) FindAll(query *articleQuery.ArticleQuery) ([]*articleEntity.Article, int64, error) {
	args := m.Called(query)
	return args.Get(resultsIndex).([]*articleEntity.Article),
		args.Get(countIndex).(int64),
		args.Error(errorIndex)
}

func (m *MockArticleRepository) Update(article *articleEntity.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func (m *MockArticleRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
