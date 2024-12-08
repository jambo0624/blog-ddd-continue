package fixtures

import (
	"gorm.io/gorm"

	articleEntity "github.com/jambo0624/blog/internal/article/domain/entity"
	categoryEntity "github.com/jambo0624/blog/internal/category/domain/entity"
	tagEntity "github.com/jambo0624/blog/internal/tag/domain/entity"
)

const InitFixturesLength = uint(2)

// TestData includes all test data.
type TestData struct {
	Categories []*categoryEntity.Category
	Tags       []*tagEntity.Tag
	Articles   []*articleEntity.Article
}

// LoadFixtures loads test data.
func LoadFixtures(db *gorm.DB) (*TestData, error) {
	data := &TestData{}

	// Create categories
	categories := []*categoryEntity.Category{
		{Name: "Technology", Slug: "technology"},
		{Name: "Life", Slug: "life"},
	}
	if err := db.Create(&categories).Error; err != nil {
		return nil, err
	}
	data.Categories = categories

	// Create tags
	tags := []*tagEntity.Tag{
		{Name: "Go", Color: "#00ADD8"},
		{Name: "DDD", Color: "#FF0000"},
	}
	if err := db.Create(&tags).Error; err != nil {
		return nil, err
	}
	data.Tags = tags

	// Create articles with tags
	articles := []*articleEntity.Article{
		{
			CategoryID: categories[0].ID,
			Title:      "Test Article 1",
			Content:    "Content 1",
		},
		{
			CategoryID: categories[1].ID,
			Title:      "Test Article 2",
			Content:    "Content 2",
		},
	}
	if err := db.Create(&articles).Error; err != nil {
		return nil, err
	}

	// Associate tags with articles using transaction
	if err := db.Transaction(func(tx *gorm.DB) error {
		// First article with first tag
		if err := tx.Model(articles[0]).Association("Tags").Append(tags[0]); err != nil {
			return err
		}
		// Second article with second tag
		if err := tx.Model(articles[1]).Association("Tags").Append(tags[1]); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	data.Articles = articles
	return data, nil
}
