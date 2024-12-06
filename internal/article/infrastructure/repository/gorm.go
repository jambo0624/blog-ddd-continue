package persistence

import (
    "gorm.io/gorm"
    articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
    articleRepository "github.com/jambo0624/blog/internal/article/domain/repository"
    articleQuery "github.com/jambo0624/blog/internal/article/domain/query"
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

func (r *GormArticleRepository) FindAll(q *articleQuery.ArticleQuery) ([]*articleEntity.Article, error) {
    var articles []*articleEntity.Article
    db := r.db

    // Apply filters
    if len(q.IDs) > 0 {
        db = db.Where("id IN ?", q.IDs)
    }
    if q.CategoryID != nil {
        db = db.Where("category_id = ?", *q.CategoryID)
    }
    if len(q.TagIDs) > 0 {
        db = db.Joins("JOIN article_tags ON articles.id = article_tags.article_id").
            Where("article_tags.tag_id IN ?", q.TagIDs)
    }
    if q.TitleLike != "" {
        db = db.Where("title LIKE ?", "%"+q.TitleLike+"%")
    }
    if q.ContentLike != "" {
        db = db.Where("content LIKE ?", "%"+q.ContentLike+"%")
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

    // Eager loading
    db = db.Preload("Category").Preload("Tags")

    err := db.Find(&articles).Error
    if err != nil {
        return nil, err
    }
    return articles, nil
}

func (r *GormArticleRepository) Update(article *articleEntity.Article) error {
    return r.db.Save(article).Error
}

