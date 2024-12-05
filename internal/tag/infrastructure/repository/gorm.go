package persistence

import (
    "gorm.io/gorm"
    tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
    tagRepository "github.com/jambo0624/blog/internal/tag/domain/repository"
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

func (r *GormTagRepository) FindByName(name string) (*tagEntity.Tag, error) {
    var tag tagEntity.Tag
    err := r.db.Where("name = ?", name).First(&tag).Error
    if err != nil {
        return nil, err
    }
    return &tag, nil
}

func (r *GormTagRepository) Update(tag *tagEntity.Tag) error {
    return r.db.Save(tag).Error
}

func (r *GormTagRepository) Delete(id uint) error {
    return r.db.Delete(&tagEntity.Tag{}, id).Error
} 