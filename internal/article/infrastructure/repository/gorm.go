package persistence

import (
    "gorm.io/gorm"
    articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
    articleRepository "github.com/jambo0624/blog/internal/article/domain/repository"
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

func (r *GormArticleRepository) FindAll() ([]*articleEntity.Article, error) {
    var articles []*articleEntity.Article
    err := r.db.Preload("Category").Preload("Tags").Find(&articles).Error
    if err != nil {
        return nil, err
    }
    return articles, nil
}

func (r *GormArticleRepository) Update(article *articleEntity.Article) error {
    return r.db.Save(article).Error
}

