package repository

import "github.com/jambo0624/blog/internal/shared/domain/query"

type Entity interface {
	GetID() uint
}

type Query interface {
	GetBaseQuery() query.BaseQuery
	Validate() error
}

type BaseRepository[T Entity, Q Query] interface {
	Save(entity *T) error
	FindByID(id uint) (*T, error)
	FindAll(query Q) ([]*T, int64, error)
	Update(entity *T) error
	Delete(id uint) error
}
