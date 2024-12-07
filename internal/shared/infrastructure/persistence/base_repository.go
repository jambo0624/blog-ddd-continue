package persistence

import (
	"time"

	"github.com/jambo0624/blog/internal/shared/domain/repository"
	"gorm.io/gorm"
)

type BaseGormRepository[T repository.Entity, Q repository.Query] struct {
	db *gorm.DB
}

func NewBaseGormRepository[T repository.Entity, Q repository.Query](db *gorm.DB) *BaseGormRepository[T, Q] {
	return &BaseGormRepository[T, Q]{db: db}
}

func (r *BaseGormRepository[T, Q]) Save(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *BaseGormRepository[T, Q]) FindByID(id uint, preloadAssociations ...string) (*T, error) {
	var entity T
	query := r.db.Model(new(T))
	if len(preloadAssociations) > 0 {
		for _, preload := range preloadAssociations {
			query = query.Preload(preload)
		}
	}
	err := query.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *BaseGormRepository[T, Q]) FindAll(q Q) ([]*T, int64, error) {
	var entities []*T
	var total int64

	// Build base query
	query := r.db.Model(new(T))

	// Apply filters (to be implemented by child repositories)
	if filterer, ok := any(q).(interface{ ApplyFilters(*gorm.DB) *gorm.DB }); ok {
		query = filterer.ApplyFilters(query)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// Apply preloads
	for _, preload := range any(q).(interface{ GetPreloadAssociations() []string }).GetPreloadAssociations() {
		query = query.Preload(preload)
	}

	baseQuery := q.GetBaseQuery()	

	// Apply pagination and sorting
	if baseQuery.Limit > 0 {
		query = query.Limit(baseQuery.Limit)
	}
	if baseQuery.Offset > 0 {
		query = query.Offset(baseQuery.Offset)
	}
	if baseQuery.OrderBy != "" {
		query = query.Order(baseQuery.OrderBy)
	}

	// Get results
	if err := query.Find(&entities).Error; err != nil {
		return nil, 0, err
	}

	return entities, total, nil
}

func (r *BaseGormRepository[T, Q]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

// Delete implements soft delete
func (r *BaseGormRepository[T, Q]) Delete(id uint) error {
	return r.db.Model(new(T)).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}