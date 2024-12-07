package tag

import (
	"github.com/stretchr/testify/mock"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	tagQuery "github.com/jambo0624/blog/internal/tag/domain/query"
)

type MockTagRepository struct {
	mock.Mock
}

func (m *MockTagRepository) Save(tag *tagEntity.Tag) error {
	args := m.Called(tag)
	return args.Error(0)
}

func (m *MockTagRepository) FindByID(id uint, preloads ...string) (*tagEntity.Tag, error) {
	args := m.Called(id, preloads)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*tagEntity.Tag), args.Error(1)
}

func (m *MockTagRepository) FindAll(query *tagQuery.TagQuery) ([]*tagEntity.Tag, int64, error) {
	args := m.Called(query)
	return args.Get(0).([]*tagEntity.Tag), args.Get(1).(int64), args.Error(2)
}

func (m *MockTagRepository) Update(tag *tagEntity.Tag) error {
	args := m.Called(tag)
	return args.Error(0)
}

func (m *MockTagRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
} 