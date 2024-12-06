package persistence

import (
    "gorm.io/gorm"
    tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
    tagRepository "github.com/jambo0624/blog/internal/tag/domain/repository"
    tagQuery "github.com/jambo0624/blog/internal/tag/domain/query"
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

func (r *GormTagRepository) FindAll(q *tagQuery.TagQuery) ([]*tagEntity.Tag, error) {
    var tags []*tagEntity.Tag
    db := r.db

    // Apply filters
    if len(q.IDs) > 0 {
        db = db.Where("id IN ?", q.IDs)
    }
    if q.NameLike != "" {
        db = db.Where("name LIKE ?", "%"+q.NameLike+"%")
    }

    // Apply pagination
    if q.Limit > 0 {
        db = db.Limit(q.Limit)
    }
    if q.Offset > 0 {
        db = db.Offset(q.Offset)
    }

    // Apply sorting
    if q.OrderBy != "" {
        db = db.Order(q.OrderBy)
    }

    err := db.Find(&tags).Error
    if err != nil {
        return nil, err
    }
    return tags, nil
}

func (r *GormTagRepository) Update(tag *tagEntity.Tag) error {
    return r.db.Save(tag).Error
}

func (r *GormTagRepository) Delete(id uint) error {
    return r.db.Delete(&tagEntity.Tag{}, id).Error
} 