package service

import (
	"github.com/jambo0624/blog/internal/shared/domain/repository"
)

type BaseService[T repository.Entity, Q repository.Query] struct {
	Repo repository.BaseRepository[T, Q]
}

func NewBaseService[T repository.Entity, Q repository.Query](
	repo repository.BaseRepository[T, Q],
) *BaseService[T, Q] {
	return &BaseService[T, Q]{
		Repo: repo,
	}
}

func (s *BaseService[T, Q]) FindByID(id uint) (*T, error) {
	return s.Repo.FindByID(id)
}

func (s *BaseService[T, Q]) FindAll(query Q) ([]*T, error) {
	if err := query.Validate(); err != nil {
		return nil, err
	}
	return s.Repo.FindAll(query)
}

func (s *BaseService[T, Q]) Delete(id uint) error {
	return s.Repo.Delete(id)
}