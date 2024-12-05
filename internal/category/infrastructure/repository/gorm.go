package persistence

import (
    "gorm.io/gorm"
    categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
    categoryRepository "github.com/jambo0624/blog/internal/category/domain/repository"
)

type GormCategoryRepository struct {
    db *gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) categoryRepository.CategoryRepository {
    return &GormCategoryRepository{db: db}
}

func (r *GormCategoryRepository) Save(category *categoryEntity.Category) error {
    return r.db.Create(category).Error
}

func (r *GormCategoryRepository) FindByID(id uint) (*categoryEntity.Category, error) {
    var category categoryEntity.Category
    err := r.db.First(&category, id).Error
    if err != nil {
        return nil, err
    }
    return &category, nil
}

func (r *GormCategoryRepository) FindBySlug(slug string) (*categoryEntity.Category, error) {
    var category categoryEntity.Category
    err := r.db.Where("slug = ?", slug).First(&category).Error
    if err != nil {
        return nil, err
    }
    return &category, nil
}

func (r *GormCategoryRepository) Update(category *categoryEntity.Category) error {
    return r.db.Save(category).Error
}

func (r *GormCategoryRepository) Delete(id uint) error {
    return r.db.Delete(&categoryEntity.Category{}, id).Error
} 