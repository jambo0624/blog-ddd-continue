package persistence

import (
	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	articleQuery "github.com/jambo0624/blog/internal/article/domain/query"
	articleRepository "github.com/jambo0624/blog/internal/article/domain/repository"
	"gorm.io/gorm"
)

type GormArticleRepository struct {
	db *gorm.DB
}

func NewGormArticleRepository(db *gorm.DB) articleRepository.ArticleRepository {
	return &GormArticleRepository{db: db}
}

func (r *GormArticleRepository) Save(article *articleEntity.Article) error {
	return r.db.Create(article).Error
}

func (r *GormArticleRepository) FindByID(id uint) (*articleEntity.Article, error) {
	var article articleEntity.Article
	err := r.db.Preload("Category").Preload("Tags").First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *GormArticleRepository) Delete(id uint) error {
	return r.db.Delete(&articleEntity.Article{}, id).Error
}

func (r *GormArticleRepository) FindAll(q *articleQuery.ArticleQuery) ([]*articleEntity.Article, int64, error) {
	var articles []*articleEntity.Article
	var total int64
	// Build query with filters
	query := r.db.Model(&articleEntity.Article{})

	// Apply filters
	if len(q.IDs) > 0 {
		query = query.Where("id IN ?", q.IDs)
	}
	if q.CategoryID != nil {
		query = query.Where("category_id = ?", *q.CategoryID)
	}
	if len(q.TagIDs) > 0 {
		query = query.Joins("JOIN article_tags ON articles.id = article_tags.article_id").
			Where("article_tags.tag_id IN ?", q.TagIDs)
	}
	if q.TitleLike != "" {
		query = query.Where("title LIKE ?", "%"+q.TitleLike+"%")
	}
	if q.ContentLike != "" {
		query = query.Where("content LIKE ?", "%"+q.ContentLike+"%")
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

	// Eager loading
	query = query.Preload("Category").Preload("Tags")

	// Get results
	if err := query.Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

func (r *GormArticleRepository) Update(article *articleEntity.Article) error {
	return r.db.Save(article).Error
}
