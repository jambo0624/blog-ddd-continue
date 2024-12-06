package persistence

import (
    "gorm.io/gorm"
    categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
    categoryRepository "github.com/jambo0624/blog/internal/category/domain/repository"
    categoryQuery "github.com/jambo0624/blog/internal/category/domain/query"
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

func (r *GormCategoryRepository) FindAll(q *categoryQuery.CategoryQuery) ([]*categoryEntity.Category, error) {
    var categories []*categoryEntity.Category
    db := r.db

    // Apply filters
    if len(q.IDs) > 0 {
        db = db.Where("id IN ?", q.IDs)
    }
    if q.NameLike != "" {
        db = db.Where("name LIKE ?", "%"+q.NameLike+"%")
    }
    if q.SlugLike != "" {
        db = db.Where("slug LIKE ?", "%"+q.SlugLike+"%")
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

    err := db.Find(&categories).Error
    if err != nil {
        return nil, err
    }
    return categories, nil
}

func (r *GormCategoryRepository) Update(category *categoryEntity.Category) error {
    return r.db.Save(category).Error
}

func (r *GormCategoryRepository) Delete(id uint) error {
    return r.db.Delete(&categoryEntity.Category{}, id).Error
} 