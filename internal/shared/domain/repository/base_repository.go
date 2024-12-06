package repository

type Entity interface {
    GetID() uint
}

type Query interface {
    Validate() error
}

type BaseRepository[T Entity, Q Query] interface {
    Save(entity *T) error
    FindByID(id uint) (*T, error)
    FindAll(query Q) ([]*T, error)
    Update(entity *T) error
    Delete(id uint) error
}