package tag

import (
	"github.com/stretchr/testify/mock"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
)

type MockTagRepository struct {
	mock.Mock
}

func (m *MockTagRepository) Save(tag *tagEntity.Tag) error {
	args := m.Called(tag)
	return args.Error(0)
}

func (m *MockTagRepository) FindByID(id uint) (*tagEntity.Tag, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*tagEntity.Tag), args.Error(1)
}

func (m *MockTagRepository) FindByName(name string) (*tagEntity.Tag, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*tagEntity.Tag), args.Error(1)
}

func (m *MockTagRepository) Update(tag *tagEntity.Tag) error {
	args := m.Called(tag)
	return args.Error(0)
}

func (m *MockTagRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
} 