package persistence

import (
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
	tagQuery "github.com/jambo0624/blog/internal/tag/domain/query"
	tagRepository "github.com/jambo0624/blog/internal/tag/domain/repository"
	"gorm.io/gorm"
)

type GormTagRepository struct {
	db *gorm.DB
}

func NewGormTagRepository(db *gorm.DB) tagRepository.TagRepository {
	return &GormTagRepository{db: db}
}

func (r *GormTagRepository) Save(tag *tagEntity.Tag) error {
	return r.db.Create(tag).Error
}

func (r *GormTagRepository) FindByID(id uint) (*tagEntity.Tag, error) {
	var tag tagEntity.Tag
	err := r.db.First(&tag, id).Error
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *GormTagRepository) FindAll(q *tagQuery.TagQuery) ([]*tagEntity.Tag, int64, error) {
	var tags []*tagEntity.Tag
	var total int64

	// Build query with filters
	query := r.db.Model(&tagEntity.Tag{})
	if len(q.IDs) > 0 {
		query = query.Where("id IN ?", q.IDs)
	}
	if q.NameLike != "" {
		query = query.Where("name LIKE ?", "%"+q.NameLike+"%")
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination and sorting
	if q.Limit > 0 {
		query = query.Limit(q.Limit)
	}
	if q.Offset > 0 {
		query = query.Offset(q.Offset)
	}
	if q.OrderBy != "" {
		query = query.Order(q.OrderBy)
	}

	// Get results
	if err := query.Find(&tags).Error; err != nil {
		return nil, 0, err
	}

	return tags, total, nil
}

func (r *GormTagRepository) Update(tag *tagEntity.Tag) error {
	return r.db.Save(tag).Error
}

func (r *GormTagRepository) Delete(id uint) error {
	return r.db.Delete(&tagEntity.Tag{}, id).Error
}

func (r *GormTagRepository) Count(q *tagQuery.TagQuery) (int64, error) {
	var count int64
	db := r.db

	// Apply filters (same as FindAll)
	if len(q.IDs) > 0 {
		db = db.Where("id IN ?", q.IDs)
	}
	if q.NameLike != "" {
		db = db.Where("name LIKE ?", "%"+q.NameLike+"%")
	}
	// ... other filters

	err := db.Model(&tagEntity.Tag{}).Count(&count).Error
	return count, err
}
