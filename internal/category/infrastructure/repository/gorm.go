package persistence

import (
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	categoryQuery "github.com/jambo0624/blog/internal/category/domain/query"
	categoryRepository "github.com/jambo0624/blog/internal/category/domain/repository"
	"gorm.io/gorm"
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

func (r *GormCategoryRepository) FindAll(q *categoryQuery.CategoryQuery) ([]*categoryEntity.Category, int64, error) {
	var categories []*categoryEntity.Category
	var total int64

	// Build query with filters
	query := r.db.Model(&categoryEntity.Category{})
	if len(q.IDs) > 0 {
		query = query.Where("id IN ?", q.IDs)
	}
	if q.NameLike != "" {
		query = query.Where("name LIKE ?", "%"+q.NameLike+"%")
	}
	if q.SlugLike != "" {
		query = query.Where("slug LIKE ?", "%"+q.SlugLike+"%")
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	if q.Limit > 0 {
		query = query.Limit(q.Limit)
	}
	if q.Offset > 0 {
		query = query.Offset(q.Offset)
	}

	// Apply sorting
	if q.OrderBy != "" {
		query = query.Order(q.OrderBy)
	}

	// Get results
	if err := query.Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

func (r *GormCategoryRepository) Update(category *categoryEntity.Category) error {
	return r.db.Save(category).Error
}

func (r *GormCategoryRepository) Delete(id uint) error {
	return r.db.Delete(&categoryEntity.Category{}, id).Error
}
