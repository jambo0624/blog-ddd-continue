package service

import (
	"fmt"

	"github.com/getsentry/sentry-go"
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

func (s *BaseService[T, Q]) FindByID(id uint, preloadAssociations ...string) (*T, error) {
	if entity, err := s.Repo.FindByID(id, preloadAssociations...); err != nil {
		sentry.CaptureException(err)
		return nil, fmt.Errorf("failed to find entity by id: %w", err)
	} else {
		return entity, nil
	}
}

func (s *BaseService[T, Q]) FindAll(query Q) ([]*T, int64, error) {
	if err := query.Validate(); err != nil {
		sentry.CaptureException(err)
		return nil, 0, fmt.Errorf("failed to validate query: %w", err)
	}

	if entities, total, err := s.Repo.FindAll(query); err != nil {
		sentry.CaptureException(err)
		return nil, 0, fmt.Errorf("failed to find all entities: %w", err)
	} else {
		return entities, total, nil
	}
}

func (s *BaseService[T, Q]) Delete(id uint) error {
	if err := s.Repo.Delete(id); err != nil {
		sentry.CaptureException(err)
		return fmt.Errorf("failed to delete entity by id: %w", err)
	}
	return nil
}